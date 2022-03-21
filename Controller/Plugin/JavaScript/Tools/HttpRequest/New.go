/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 14:52:53
 * @LastEditTime: 2022-03-21 14:59:38
 * @LastEditors: NyanCatda
 * @Description: 请求请求函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\HttpRequest\New.go
 */
package HttpRequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 * @description: 新建请求函数注册
 * @param {string} Method 请求方法 GET/POST/PUT/DELETE...
 * @param {string} URL 请求地址
 * @param {[]string} Header 请求头
 * @param {string} RequestBody 请求内容
 * @return {string} 返回内容
 * @return {*http.Response} 请求响应信息
 */
func (HttpRequest HttpRequest) New(Method string, URL string, Header []string, RequestBody string) (string, *http.Response) {
	RequestBodyStr := []byte(RequestBody)
	req, err := http.NewRequest(Method, URL, bytes.NewBuffer(RequestBodyStr))
	if err != nil {
		return "", nil
	}

	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		req.Header.Set(Headervalue[0], Headervalue[1])
	}

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()
	Body, _ := ioutil.ReadAll(resp.Body)

	return string(Body), resp
}
