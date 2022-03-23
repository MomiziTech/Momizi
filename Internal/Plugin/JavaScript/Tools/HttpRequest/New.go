/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 14:52:53
 * @LastEditTime: 2022-03-23 21:45:04
 * @LastEditors: NyanCatda
 * @Description: 请求请求函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\HttpRequest\New.go
 */
package HttpRequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/dop251/goja"
)

/**
 * @description: 新建请求函数注册
 * @param {string} Method 请求方法 GET/POST/PUT/DELETE...
 * @param {string} URL 请求地址
 * @param {[]string} Header 请求头
 * @param {string} RequestBody 请求内容
 * @param {goja.Callable} Callback 回调函数
 * @return {[]byte} Body 请求返回内容
 * @return {*http.Response} HttpResponse Http响应
 */
func (HttpRequest HttpRequest) New(Method string, URL string, Header []string, RequestBody string, Callback goja.Callable) {
	go func() {
		RequestBodyStr := []byte(RequestBody)
		req, err := http.NewRequest(Method, URL, bytes.NewBuffer(RequestBodyStr))
		if err != nil {
			PluginName := HttpRequest.VM.Get("PLUGIN_NAME").String()
			Log.Error(PluginName, err)
			Callback(nil, HttpRequest.VM.ToValue(""), HttpRequest.VM.ToValue(nil))
			return
		}

		for _, value := range Header {
			Headervalue := strings.Split(value, ":")
			req.Header.Set(Headervalue[0], Headervalue[1])
		}

		// 发起请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			PluginName := HttpRequest.VM.Get("PLUGIN_NAME").String()
			Log.Error(PluginName, err)
			Callback(nil, HttpRequest.VM.ToValue(""), HttpRequest.VM.ToValue(nil))
			return
		}
		defer resp.Body.Close()
		Body, _ := ioutil.ReadAll(resp.Body)

		Callback(nil, HttpRequest.VM.ToValue(string(Body)), HttpRequest.VM.ToValue(resp))
	}()
}
