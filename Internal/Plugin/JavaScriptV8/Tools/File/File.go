/*
 * @Author: McPlus
 * @Date: 2022-03-28 17:24:26
 * @LastEditTime: 2022-04-02 09:28:39
 * @LastEditors: McPlus
 * @Description: 文件类
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\File\File.go
 */
package File

import (
	"os"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Tools/File"
	"rogchap.com/v8go"
)

var DataPath = Controller.DataPath + "/"

func Register(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.Object {
	File, _ := v8go.NewObjectTemplate(Isolate)

	Stream := RegisterConstructor(Isolate, Context)
	File.Set("Stream", Stream)

	FlagEnum := RegisterFlagEnum(Isolate, Context)
	File.Set("Flag", FlagEnum)

	FileInstance, _ := File.NewInstance(Context)

	return FileInstance
}

func RegisterConstructor(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Constructor, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Path := Info.Args()[0].String()
		Flag := Info.Args()[1].Int32()

		PluginName, _ := Context.Global().Get("PLUGIN_NAME")

		ReadWrite, _ := File.NewFileReadWrite(DataPath+PluginName.String()+"/"+Path, int(Flag))
		FileObject, _ := v8go.NewObjectTemplate(Isolate)
		Instance, _ := FileObject.NewInstance(Info.Context())
		Instance.Set("path", Path)
		Instance.Set("flag", Flag)

		RegisterObjectFunction(Isolate, Context, Instance, Info, ReadWrite)

		return Instance.Value
	})

	return Constructor
}

func RegisterObjectFunction(Isolate *v8go.Isolate, Context *v8go.Context, Instance *v8go.Object, ObjectInfo *v8go.FunctionCallbackInfo, ReadWrite *File.ReadWrite) {
	Read, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Content, _ := ReadWrite.Read()

		Value, _ := v8go.NewValue(Isolate, Content)

		return Value
	})
	Instance.Set("Read", Read.GetFunction(Context))

	Write, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Content := Info.Args()[0]

		err := ReadWrite.WriteTo(Content.String())

		if err == nil {
			val, _ := v8go.NewValue(Isolate, true)
			return val
		}

		val, _ := v8go.NewValue(Isolate, false)
		return val
	})
	Instance.Set("Write", Write.GetFunction(Context))

	Append, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		Content := Info.Args()[0]

		err := ReadWrite.WriteAppend(Content.String())

		if err == nil {
			val, _ := v8go.NewValue(Isolate, true)
			return val
		}

		val, _ := v8go.NewValue(Isolate, false)
		return val
	})
	Instance.Set("Append", Append.GetFunction(Context))

	Close, _ := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		err := ReadWrite.Close()

		if err == nil {
			val, _ := v8go.NewValue(Isolate, true)
			return val
		}

		val, _ := v8go.NewValue(Isolate, false)
		return val
	})
	Instance.Set("Close", Close.GetFunction(Context))
}

func RegisterFlagEnum(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.ObjectTemplate {
	FlagEnum, _ := v8go.NewObjectTemplate(Isolate)

	FlagEnum.Set("Read", int32(os.O_RDONLY))
	FlagEnum.Set("WriteCover", int32(os.O_WRONLY|os.O_TRUNC|os.O_CREATE))
	FlagEnum.Set("WriteAppend", int32(os.O_WRONLY|os.O_APPEND|os.O_CREATE))

	return FlagEnum
}
