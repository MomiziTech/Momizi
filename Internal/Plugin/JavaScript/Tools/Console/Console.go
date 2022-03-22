/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 01:27:22
 * @LastEditTime: 2022-03-22 18:00:46
 * @LastEditors: NyanCatda
 * @Description: 控制台函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\Console\Console.go
 */
package Console

import "github.com/dop251/goja"

type Console struct {
	VM *goja.Runtime
}

/**
 * @description: 注册控制台函数
 * @param {*goja.Runtime} VM
 * @return {*}
 */
func RegistrationFunction(VM *goja.Runtime) error {
	return VM.Set("Console", Console{VM: VM})
}
