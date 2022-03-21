/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 22:20:17
 * @LastEditTime: 2022-03-22 01:29:08
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\Tools.go
 */
package Tools

import (
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript/Tools/HttpRequest"
	"github.com/dop251/goja"
)

func Tools(VM *goja.Runtime) error {
	// 注册Http请求类
	if err := HttpRequest.HttpRequests(VM); err != nil {
		return err
	}

	return nil
}
