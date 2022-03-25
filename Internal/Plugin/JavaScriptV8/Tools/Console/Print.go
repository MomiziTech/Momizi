/*
 * @Author: McPlus
 * @Date: 2022-03-25 05:22:47
 * @LastEditTime: 2022-03-25 05:37:31
 * @LastEditors: McPlus
 * @Description:
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Console\Log.go
 */

package Console

import (
	"github.com/MomiziTech/Momizi/Tools/Log"
	"rogchap.com/v8go"
)

func RegisterInfoPrint(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Log, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Text := Info.Args()[0]

		PLUGIN_NAME, _ := Context.RunScript("PLUGIN_NAME", "")

		Log.Print(PLUGIN_NAME.String(), Log.INFO, Text)

		return nil
	})

	return Log
}

func RegisterWarningPrint(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Warning, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Text := Info.Args()[0]

		PLUGIN_NAME, _ := Context.RunScript("PLUGIN_NAME", "")

		Log.Print(PLUGIN_NAME.String(), Log.INFO, Text)

		return nil
	})

	return Warning
}

func RegisterErrorPrint(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Error, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Text := Info.Args()[0]

		PLUGIN_NAME, _ := Context.RunScript("PLUGIN_NAME", "")

		Log.Print(PLUGIN_NAME.String(), Log.INFO, Text)

		return nil
	})

	return Error
}

func RegisterDebugPrint(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Debug, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Text := Info.Args()[0]

		PLUGIN_NAME, _ := Context.RunScript("PLUGIN_NAME", "")

		Log.Print(PLUGIN_NAME.String(), Log.INFO, Text)

		return nil
	})

	return Debug
}
