/*
 * @Author: NyanCatda
 * @Date: 2022-03-27 00:09:06
 * @LastEditTime: 2022-04-04 13:00:59
 * @LastEditors: NyanCatda
 * @Description: Post请求函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\HttpRequest\Post.go
 */
package HttpRequest

import (
	"path/filepath"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Loader"
	"github.com/MomiziTech/Momizi/Tools/Log"
	HttpRequestFunc "github.com/nyancatda/HttpRequest"
	"rogchap.com/v8go"
)

/**
 * @description: Post请求函数注册，传递Json
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @return {*v8go.FunctionTemplate} Post请求函数
 */
func PostJson(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	PluginName, err := Context.RunScript("PLUGIN_NAME", "")
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}
	PostJson, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		URL := Info.Args()[0]         // {string} 请求地址
		Header := Info.Args()[1]      // {[]string} 请求头
		RequestBody := Info.Args()[2] // {string} 请求内容(Json)
		CallBack := Info.Args()[3]    // 回调函数

		if CallBack.IsFunction() {
			go func() {
				// 获取请求头
				Headers, err := Loader.V8StringArrayToGoStringArray(Header)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 发起请求
				Body, HttpResponseValue, err := HttpRequestFunc.PostRequestJson(URL.String(), Headers, RequestBody.String())
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析返回值
				BodyValue, HttpResponseObject, err := CallbackParameter(Isolate, Context, Body, HttpResponseValue)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				CallBack, _ := CallBack.AsFunction()
				// 参数写入回调函数
				CallBack.Call(BodyValue, HttpResponseObject)
			}()
		}

		return nil
	})
	if err != nil {
		Log.Error(PluginName.String(), err)
		return nil
	}

	return PostJson
}

/**
 * @description Post请求函数注册，传递x-www-from-urlencoded
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @return {*v8go.FunctionTemplate} Post请求函数
 */
func PostXWWWForm(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	PluginName, err := Context.RunScript("PLUGIN_NAME", "")
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}
	PostXWWWForm, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		URL := Info.Args()[0]      // {string} 请求地址
		Header := Info.Args()[1]   // {[]string} 请求头
		Data := Info.Args()[2]     // {map[string]string} 请求内容(x-www-from-urlencoded)
		CallBack := Info.Args()[3] // 回调函数

		if CallBack.IsFunction() {
			go func() {
				// 获取请求头
				Headers, err := Loader.V8StringArrayToGoStringArray(Header)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析Data
				DataMap, err := Loader.V8ObjectToGoStringMap(Context, Data)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 发起请求
				Body, HttpResponseValue, err := HttpRequestFunc.PostRequestXWWWForm(URL.String(), Headers, DataMap)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析返回值
				BodyValue, HttpResponseObject, err := CallbackParameter(Isolate, Context, Body, HttpResponseValue)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				CallBack, _ := CallBack.AsFunction()
				// 参数写入回调函数
				CallBack.Call(BodyValue, HttpResponseObject)
			}()
		}

		return nil
	})
	if err != nil {
		Log.Error(PluginName.String(), err)
		return nil
	}

	return PostXWWWForm
}

/**
 * @description: Post请求函数注册，传递multipart/form-data
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @return {*v8go.FunctionTemplate} Post请求函数
 */
func PostFormData(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	PluginName, err := Context.RunScript("PLUGIN_NAME", "")
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}
	PostFormData, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		URL := Info.Args()[0]      // {string} 请求地址
		Header := Info.Args()[1]   // {[]string} 请求头
		Data := Info.Args()[2]     // {map[string]string} 请求内容(multipart/form-data)
		CallBack := Info.Args()[3] // 回调函数

		if CallBack.IsFunction() {
			go func() {
				// 获取请求头
				Headers, err := Loader.V8StringArrayToGoStringArray(Header)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析Data
				DataMap, err := Loader.V8ObjectToGoStringMap(Context, Data)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 发起请求
				Body, HttpResponseValue, err := HttpRequestFunc.PostRequestFormData(URL.String(), Headers, DataMap)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析返回值
				BodyValue, HttpResponseObject, err := CallbackParameter(Isolate, Context, Body, HttpResponseValue)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				CallBack, _ := CallBack.AsFunction()
				// 参数写入回调函数
				CallBack.Call(BodyValue, HttpResponseObject)
			}()
		}

		return nil
	})
	if err != nil {
		Log.Error(PluginName.String(), err)
		return nil
	}

	return PostFormData
}

/**
 * @description: Post请求函数注册，带文件传递multipart/form-data
 * @param {*v8go.Isolate} Isolate v8实例
 * @param {*v8go.Context} Context v8上下文
 * @return {*v8go.FunctionTemplate} Post请求函数
 */
func PostFormDataFile(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	PluginName, err := Context.RunScript("PLUGIN_NAME", "")
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}
	PostFormDataFile, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		URL := Info.Args()[0]      // {string} 请求地址
		Header := Info.Args()[1]   // {[]string} 请求头
		Data := Info.Args()[2]     // {map[string]string} 请求内容(multipart/form-data)
		FileKey := Info.Args()[3]  // {string} 文件请求参数(key)
		FilePath := Info.Args()[4] // {[]string} 文件路径
		CallBack := Info.Args()[5] // 回调函数

		if CallBack.IsFunction() {
			go func() {
				// 获取请求头
				Headers, err := Loader.V8StringArrayToGoStringArray(Header)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析Data
				DataMap, err := Loader.V8ObjectToGoStringMap(Context, Data)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析FilePath
				FilePaths, err := Loader.V8StringArrayToGoStringArray(FilePath)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				var FilePathArray []string
				// 限制FilePath可以使用的位置
				for _, FilePath := range FilePaths {
					FilePath = filepath.Clean(Controller.DataPath + "/" + PluginName.String() + "/" + FilePath)
					FilePathArray = append(FilePathArray, FilePath)
				}

				// 发起请求
				Body, HttpResponseValue, err := HttpRequestFunc.PostRequestFormDataFile(URL.String(), Headers, DataMap, FileKey.String(), FilePathArray)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				// 解析返回值
				BodyValue, HttpResponseObject, err := CallbackParameter(Isolate, Context, Body, HttpResponseValue)
				if err != nil {
					Log.Error(PluginName.String(), err)
					return
				}

				CallBack, _ := CallBack.AsFunction()
				// 参数写入回调函数
				CallBack.Call(BodyValue, HttpResponseObject)
			}()
		}

		return nil
	})
	if err != nil {
		Log.Error(PluginName.String(), err)
		return nil
	}

	return PostFormDataFile
}
