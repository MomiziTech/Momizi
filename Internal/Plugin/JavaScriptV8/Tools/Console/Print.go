/*
 * @Author: McPlus
 * @Date: 2022-03-25 05:22:47
 * @LastEditTime: 2022-03-28 15:44:31
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\Console\Print.go
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

		Log.Info(PLUGIN_NAME.String(), Text)

		return nil
	})

	return Log
}

func RegisterWarningPrint(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Warning, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Text := Info.Args()[0]

		PLUGIN_NAME, _ := Context.RunScript("PLUGIN_NAME", "")

		Log.Warning(PLUGIN_NAME.String(), Text)

		return nil
	})

	return Warning
}

func RegisterErrorPrint(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Error, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Text := Info.Args()[0]

		PLUGIN_NAME, _ := Context.RunScript("PLUGIN_NAME", "")

		Log.Print(PLUGIN_NAME.String(), Log.ERROR, Text)

		return nil
	})

	return Error
}

func RegisterDebugPrint(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Debug, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Text := Info.Args()[0]

		PLUGIN_NAME, _ := Context.RunScript("PLUGIN_NAME", "")

		Log.DeBug(PLUGIN_NAME.String(), Text)

		return nil
	})

	return Debug
}

func ResisterSetColorFunction(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	SetColor, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Color := Info.Args()[0].String()
		Text := Info.Args()[1].String()

		Isolate, _ := v8go.NewIsolate()

		switch Color {
		case "Black":
			Value, _ := v8go.NewValue(Isolate, Log.Black(Text))
			return Value
		case "Red":
			Value, _ := v8go.NewValue(Isolate, Log.Red(Text))
			return Value
		case "Green":
			Value, _ := v8go.NewValue(Isolate, Log.Green(Text))
			return Value
		case "Yellow":
			Value, _ := v8go.NewValue(Isolate, Log.Yellow(Text))
			return Value
		case "Blue":
			Value, _ := v8go.NewValue(Isolate, Log.Blue(Text))
			return Value
		case "Magenta":
			Value, _ := v8go.NewValue(Isolate, Log.Magenta(Text))
			return Value
		case "Cyan":
			Value, _ := v8go.NewValue(Isolate, Log.Cyan(Text))
			return Value
		case "White":
			Value, _ := v8go.NewValue(Isolate, Log.White(Text))
			return Value
		}

		return nil
	})

	return SetColor
}
