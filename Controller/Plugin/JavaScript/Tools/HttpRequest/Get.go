/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 13:58:24
 * @LastEditTime: 2022-03-21 19:17:23
 * @LastEditors: NyanCatda
 * @Description: Get请求函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\HttpRequest\Get.go
 */
package HttpRequest

import (
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/dop251/goja"
	HttpRequestFunc "github.com/nyancatda/HttpRequest"
)

/**
 * @description: GET请求函数注册
 * @param {string} URL
 * @param {[]string} Header
 * @return {*}
 */
func (HttpRequest HttpRequest) Get(URL string, Header []string, Func goja.Callable) {
	go func() {
		Body, HttpResponse, err := HttpRequestFunc.GetRequest(URL, Header)
		if err != nil {
			Log.Error("Plugin", err)
			Func(nil, HttpRequest.VM.ToValue(""), HttpRequest.VM.ToValue(nil))
			return
		}

		Func(nil, HttpRequest.VM.ToValue(string(Body)), HttpRequest.VM.ToValue(HttpResponse))
	}()
}
