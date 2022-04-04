/*
 * @Author: McPlus
 * @Date: 2022-03-24 21:47:22
 * @LastEditTime: 2022-04-04 13:17:50
 * @LastEditors: NyanCatda
 * @Description: 工具函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Tools.go
 */
package Tools

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Config"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Console"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/File"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/HttpRequest"
	"rogchap.com/v8go"
)

/**
 * @description: 工具函数注册
 * @param {*v8go.Isolate} Isolate
 * @param {*v8go.Context} Context
 * @return {*}
 */
func Register(Isolate *v8go.Isolate, Context *v8go.Context) error {
	Global := Context.Global()
	// 控制台函数注册
	err := Global.Set("Console", Console.Register(Isolate, Context))
	if err != nil {
		return err
	}
	// HttpRequest函数注册
	HttpRequestObject, err := HttpRequest.Register(Isolate, Context)
	if err == nil {
		err = Global.Set("HttpRequest", HttpRequestObject)
		if err != nil {
			return err
		}
	}
	// File函数注册
	err = Global.Set("File", File.Register(Isolate, Context))
	if err != nil {
		return err
	}
	// 配置文件操作函数注册
	err = Global.Set("Config", Config.Register(Isolate, Context))
	if err != nil {
		return err
	}

	return nil
}
