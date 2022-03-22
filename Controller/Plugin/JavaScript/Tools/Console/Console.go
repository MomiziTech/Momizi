/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 01:27:22
 * @LastEditTime: 2022-03-22 17:27:41
 * @LastEditors: NyanCatda
 * @Description: 控制台函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\Console\Console.go
 */
package Console

import "github.com/dop251/goja"

type Console struct {
	VM *goja.Runtime
}

func RegistrationFunction(VM *goja.Runtime) error {
	return VM.Set("Console", Console{VM: VM})
}
