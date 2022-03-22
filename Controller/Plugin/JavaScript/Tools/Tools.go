/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 22:20:17
 * @LastEditTime: 2022-03-22 23:05:31
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\Tools.go
 */
package Tools

import (
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript/Tools/Console"
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript/Tools/File"
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript/Tools/HttpRequest"
	"github.com/dop251/goja"
)

/**
 * @description: 工具函数注册
 * @param {*goja.Runtime} VM 虚拟机
 * @return {*}
 */
func Tools(VM *goja.Runtime) error {
	// 注册控制台类
	if err := Console.RegistrationFunction(VM); err != nil {
		return err
	}

	// 注册Http请求类
	if err := HttpRequest.HttpRequests(VM); err != nil {
		return err
	}

	// 注册文件操作类
	if err := File.RegistrationFunction(VM); err != nil {
		return err
	}

	return nil
}
