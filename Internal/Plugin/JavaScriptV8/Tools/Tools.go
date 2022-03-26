/*
 * @Author: McPlus
 * @Date: 2022-03-24 21:47:22
 * @LastEditTime: 2022-03-26 19:05:36
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Tools.go
 */
package Tools

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Console"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/HttpRequest"
	"rogchap.com/v8go"
)

func Register(Isolate *v8go.Isolate, Context *v8go.Context) error {
	Global := Context.Global()
	// 控制台函数注册
	err := Global.Set("Console", Console.Register(Isolate, Context))
	if err != nil {
		return err
	}
	// HttpRequest函数注册
	err = Global.Set("HttpRequest", HttpRequest.Register(Isolate, Context))
	if err != nil {
		return err
	}

	return nil
}
