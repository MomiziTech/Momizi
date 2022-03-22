/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 13:55:04
 * @LastEditTime: 2022-03-21 13:57:29
 * @LastEditors: NyanCatda
 * @Description: Http请求函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\HttpRequest\HttpRequest.go
 */
package HttpRequest

import "github.com/dop251/goja"

/**
 * @description: Http请求类
 */
type HttpRequest struct {
	VM *goja.Runtime
}

/**
 * @description: Http请求类
 * @param {*goja.Runtime} VM 加载器
 * @return {*}
 */
func HttpRequests(VM *goja.Runtime) error {
	// 注册Http请求类
	err := VM.Set("HttpRequest", HttpRequest{VM})
	if err != nil {
		return err
	}

	return nil
}
