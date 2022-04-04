/*
 * @Author: NyanCatda
 * @Date: 2022-03-27 02:20:17
 * @LastEditTime: 2022-03-27 02:25:58
 * @LastEditors: NyanCatda
 * @Description: New请求函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\HttpRequest\New.go
 */
package HttpRequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Loader"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"rogchap.com/v8go"
)

func New(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	PluginName, err := Context.RunScript("PLUGIN_NAME", "")
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}
	New, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Method := Info.Args()[0]      // {string}Method 请求方法 GET/POST/PUT/DELETE...
		URL := Info.Args()[1]         // {string}请求地址
		Header := Info.Args()[2]      // {[]string}请求头
		RequestBody := Info.Args()[3] // {string}请求体
		CallBack := Info.Args()[4]    // 回调函数

		if CallBack.IsFunction() {
			go func() {
				// 获取请求头
				Headers, err := Loader.V8StringArrayToGoStringArray(Header)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 准备请求
				RequestBodyStr := []byte(RequestBody.String())
				req, err := http.NewRequest(Method.String(), URL.String(), bytes.NewBuffer(RequestBodyStr))
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 设置请求头
				for _, value := range Headers {
					Headervalue := strings.Split(value, ":")
					// 如果解析失败则不设置请求头
					if len(Headervalue) <= 0 {
						break
					}
					req.Header.Set(Headervalue[0], Headervalue[1])
				}

				// 发起请求
				client := &http.Client{}
				HttpResponseValue, err := client.Do(req)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}
				defer HttpResponseValue.Body.Close()
				Body, _ := ioutil.ReadAll(HttpResponseValue.Body)

				// 解析返回值
				BodyValue, HttpResponseObject, err := CallbackParameter(Isolate, Context, Body, HttpResponseValue)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				CallBack, _ := CallBack.AsFunction()
				CallBack.Call(BodyValue, HttpResponseObject)
			}()
		}
		return nil
	})
	if err != nil {
		Log.Error(PluginName.String(), err)
		return nil
	}

	return New
}
