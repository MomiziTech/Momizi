/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 14:14:03
 * @LastEditTime: 2022-03-21 14:22:31
 * @LastEditors: NyanCatda
 * @Description: Post请求函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\HttpRequest\Post.go
 */
package HttpRequest

import (
	"net/http"

	"github.com/MomiziTech/Momizi/Utils/Log"
	HttpRequestFunc "github.com/nyancatda/HttpRequest"
)

/**
 * @description: POST请求函数注册，传递Json
 * @param {string} URL 请求地址
 * @param {[]string} Header 请求头
 * @param {string} requestBody 请求内容(Json)
 * @return {*}
 */
func (HttpRequest HttpRequest) PostRequestJson(URL string, Header []string, requestBody string) (string, *http.Response) {
	Body, HttpResponse, err := HttpRequestFunc.PostRequestJson(URL, Header, requestBody)
	if err != nil {
		Log.ErrorWrite("Plugin", err)
		return "", nil
	}

	StringBody := string(Body)

	return StringBody, HttpResponse
}

/**
 * @description: POST请求函数注册，传递x-www-from-urlencoded
 * @param {string} URL 请求地址
 * @param {[]string} Header 请求头
 * @param {map[string]string} Data 请求内容(x-www-from-urlencoded)
 * @return {*}
 */
func (HttpRequest HttpRequest) PostRequestXWWWForm(URL string, Header []string, Data map[string]string) (string, *http.Response) {
	Body, HttpResponse, err := HttpRequestFunc.PostRequestXWWWForm(URL, Header, Data)
	if err != nil {
		Log.ErrorWrite("Plugin", err)
		return "", nil
	}

	StringBody := string(Body)

	return StringBody, HttpResponse
}

/**
 * @description: POST请求函数注册，传递multipart/form-data
 * @param {string} URL 请求地址
 * @param {[]string} Header 请求头
 * @param {map[string]string} Data 请求内容(multipart/form-data)
 * @return {*}
 */
func (HttpRequest HttpRequest) PostRequestFormData(URL string, Header []string, Data map[string]string) (string, *http.Response) {
	Body, HttpResponse, err := HttpRequestFunc.PostRequestFormData(URL, Header, Data)
	if err != nil {
		Log.ErrorWrite("Plugin", err)
		return "", nil
	}

	StringBody := string(Body)

	return StringBody, HttpResponse
}

/**
 * @description: POST请求函数注册，带文件传递multipart/form-data
 * @param {string} URL 请求地址
 * @param {[]string} Header 请求头
 * @param {map[string]string} Data 请求内容(multipart/form-data)
 * @param {string} FileKey 文件key
 * @param {[]string} FilePath 文件路径
 * @return {*}
 */
func (HttpRequest HttpRequest) PostRequestFormDataFile(URL string, Header []string, Data map[string]string, FileKey string, FilePath []string) (string, *http.Response) {
	Body, HttpResponse, err := HttpRequestFunc.PostRequestFormDataFile(URL, Header, Data, FileKey, FilePath)
	if err != nil {
		Log.ErrorWrite("Plugin", err)
		return "", nil
	}

	StringBody := string(Body)

	return StringBody, HttpResponse
}
