/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 01:28:05
 * @LastEditTime: 2022-03-22 17:16:47
 * @LastEditors: McPlus
 * @Description: 日志输出函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\Console\Log.go
 */
package Console

import (
	"github.com/MomiziTech/Momizi/Utils/Log"
	"github.com/dop251/goja"
)

func (Console Console) Log(Text string) error {
	return logDefault(Console, Log.INFO, Text)
}

func (Console Console) Warning(Text string) error {
	return logDefault(Console, Log.WARNING, Text)
}

func (Console Console) Error(Text string) error {
	return logDefault(Console, Log.ERROR, Text)
}

func (Console Console) Debug(Text string) error {
	return logDefault(Console, Log.DEBUG, Text)
}

func logDefault(Console Console, Level int, Text string) error {
	return Log.Print(Console.VM.Get("PLUGIN_NAME").String(), Level, Text)
}

func RegistrationFunction(VM *goja.Runtime) error {
	return VM.Set("Console", Console{VM: VM})
}
