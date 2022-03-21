/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 13:58:24
 * @LastEditTime: 2022-03-21 14:17:26
 * @LastEditors: NyanCatda
 * @Description: Get请求函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\HttpRequest\Get.go
 */
package HttpRequest

import (
	"net/http"

	"github.com/MomiziTech/Momizi/Utils/Log"
	HttpRequestFunc "github.com/nyancatda/HttpRequest"
)

/**
 * @description: GET请求函数注册
 * @param {string} URL
 * @param {[]string} Header
 * @return {*}
 */
func (HttpRequest HttpRequest) GetRequest(URL string, Header []string) (string, *http.Response) {
	Body, HttpResponse, err := HttpRequestFunc.GetRequest(URL, Header)
	if err != nil {
		Log.ErrorWrite("Plugin", err)
		return "", nil
	}

	StringBody := string(Body)

	return StringBody, HttpResponse
}
