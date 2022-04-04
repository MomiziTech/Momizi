/*
 * @Author: NyanCatda
 * @Date: 2022-03-26 10:27:08
 * @LastEditTime: 2022-04-04 12:57:47
 * @LastEditors: NyanCatda
 * @Description: Get请求函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\HttpRequest\Get.go
 */
package HttpRequest

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Loader"
	"github.com/MomiziTech/Momizi/Tools/Log"

	HttpRequestFunc "github.com/nyancatda/HttpRequest"
	"rogchap.com/v8go"
)

/**
 * @description: Get请求函数注册
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @return {*v8go.FunctionTemplate} Get请求函数
 */
func Get(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	PluginName, err := Context.RunScript("PLUGIN_NAME", "")
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}
	Get, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		URL := Info.Args()[0]      // {string}请求地址
		Header := Info.Args()[1]   // {[]string}请求头
		CallBack := Info.Args()[2] // 回调函数

		if CallBack.IsFunction() {
			go func() {
				// 获取请求头
				Headers, err := Loader.V8StringArrayToGoStringArray(Header)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 发起请求
				Body, HttpResponseValue, err := HttpRequestFunc.GetRequest(URL.String(), Headers)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

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

	return Get
}
