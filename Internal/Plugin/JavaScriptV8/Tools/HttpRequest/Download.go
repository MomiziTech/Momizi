/*
 * @Author: NyanCatda
 * @Date: 2022-03-27 02:31:39
 * @LastEditTime: 2022-04-04 12:57:23
 * @LastEditors: NyanCatda
 * @Description: 下载文件函数封装
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\HttpRequest\Download.go
 */
package HttpRequest

import (
	"path/filepath"
	"strings"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8/Tools/Loader"
	"github.com/MomiziTech/Momizi/Tools"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"rogchap.com/v8go"
)

func Download(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.FunctionTemplate {
	Download, err := v8go.NewFunctionTemplate(Isolate, func(Info *v8go.FunctionCallbackInfo) *v8go.Value {
		URL := Info.Args()[0]      // {string}请求地址
		Header := Info.Args()[1]   // {[]string}请求头
		SavePath := Info.Args()[2] // {string}文件保存路径
		CallBack := Info.Args()[3] // 回调函数

		if CallBack.IsFunction() {
			go func() {
				// 获取请求头
				Headers, err := Loader.V8StringArrayToGoStringArray(Header)
				if err != nil {
					PluginName, _ := Context.RunScript("PLUGIN_NAME", "")
					Log.Error(PluginName.String(), err)
					return
				}

				// 发起请求
				PluginName, err := Context.RunScript("PLUGIN_NAME", "")
				if err != nil {
					Log.Error("Plugin", err)
					return
				}
				FilePath, FileSize, err := Tools.DownloadFile(URL.String(), Headers, Controller.DataPath+"/"+PluginName.String()+"/"+SavePath.String(), false, 120)
				if err != nil {
					PluginName, _ := Context.RunScript("PLUGIN_NAME", "")
					Log.Error(PluginName.String(), err)
					return
				}

				// 去除文件路径前缀
				FilePath = strings.TrimPrefix(FilePath, filepath.Clean(Controller.DataPath+"/"+PluginName.String()+"/"))

				// 转换返回值类型
				FilePathValue, err := v8go.NewValue(Isolate, FilePath)
				if err != nil {
					PluginName, _ := Context.RunScript("PLUGIN_NAME", "")
					Log.Error(PluginName.String(), err)
					return
				}

				FileSizeValue, err := v8go.NewValue(Isolate, FileSize)
				if err != nil {
					PluginName, _ := Context.RunScript("PLUGIN_NAME", "")
					Log.Error(PluginName.String(), err)
					return
				}

				// 返回结果
				CallBack, _ := CallBack.AsFunction()
				CallBack.Call(FilePathValue, FileSizeValue)
			}()
		}
		return nil
	})
	if err != nil {
		PluginName, _ := Context.RunScript("PLUGIN_NAME", "")
		Log.Error(PluginName.String(), err)
		return nil
	}

	return Download
}
