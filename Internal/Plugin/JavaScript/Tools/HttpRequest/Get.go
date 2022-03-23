/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 13:58:24
 * @LastEditTime: 2022-03-23 21:44:55
 * @LastEditors: NyanCatda
 * @Description: Get请求函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\HttpRequest\Get.go
 */
package HttpRequest

import (
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/dop251/goja"
	HttpRequestFunc "github.com/nyancatda/HttpRequest"
)

/**
 * @description: GET请求函数注册
 * @param {string} URL 请求地址
 * @param {[]string} Header 头部信息
 * @param {goja.Callable} Callback 回调函数
 * @return {[]byte} Body 请求返回内容
 * @return {*http.Response} HttpResponse Http响应
 */
func (HttpRequest HttpRequest) Get(URL string, Header []string, Callback goja.Callable) {
	go func() {
		Body, HttpResponse, err := HttpRequestFunc.GetRequest(URL, Header)
		if err != nil {
			PluginName := HttpRequest.VM.Get("PLUGIN_NAME").String()
			Log.Error(PluginName, err)
			Callback(nil, HttpRequest.VM.ToValue(""), HttpRequest.VM.ToValue(nil))
			return
		}

		Callback(nil, HttpRequest.VM.ToValue(string(Body)), HttpRequest.VM.ToValue(HttpResponse))
	}()
}
