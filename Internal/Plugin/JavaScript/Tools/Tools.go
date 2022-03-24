/*
 * @Author: NyanCatda
 * @Date: 2022-03-20 22:20:17
 * @LastEditTime: 2022-03-24 14:40:03
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\Tools.go
 */
package Tools

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScript/Tools/Config"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScript/Tools/Console"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScript/Tools/File"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScript/Tools/HttpRequest"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScript/Tools/LevelDB"
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

	// 注册配置文件类
	if err := Config.RegistrationFunction(VM); err != nil {
		return err
	}

	// 注册LevelDB类
	if err := LevelDB.RegistrationFunction(VM); err != nil {
		return err
	}

	return nil
}
