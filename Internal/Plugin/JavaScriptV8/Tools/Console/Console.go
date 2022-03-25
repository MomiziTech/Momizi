/*
 * @Author: McPlus
 * @Date: 2022-03-25 05:22:32
 * @LastEditTime: 2022-03-25 20:12:22
 * @LastEditors: McPlus
 * @Description: 控制台方法
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Console\Console.go
 */
package Console

import "rogchap.com/v8go"

func Register(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.Object {
	Console, _ := v8go.NewObjectTemplate(Isolate)

	Info := RegisterInfoPrint(Isolate, Context)
	Console.Set("Log", Info)

	Warning := RegisterWarningPrint(Isolate, Context)
	Console.Set("Warning", Warning)

	Error := RegisterErrorPrint(Isolate, Context)
	Console.Set("Error", Error)

	Debug := RegisterDebugPrint(Isolate, Context)
	Console.Set("Debug", Debug)

	ConsoleObject, _ := Console.NewInstance(Context)

	return ConsoleObject

}
