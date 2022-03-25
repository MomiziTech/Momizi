/*
 * @Author: McPlus
 * @Date: 2022-03-24 21:47:22
 * @LastEditTime: 2022-03-25 22:41:29
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Tools.go
 */
package Tools

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Console"
	"rogchap.com/v8go"
)

func Register(Isolate *v8go.Isolate, Context *v8go.Context) error {
	Global := Context.Global()
	// 控制台函数注册
	err := Global.Set("Console", Console.Register(Isolate, Context))
	if err != nil {
		return err
	}

	return nil
}
