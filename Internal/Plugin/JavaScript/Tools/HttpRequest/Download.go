/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 17:36:42
 * @LastEditTime: 2022-03-22 21:54:56
 * @LastEditors: NyanCatda
 * @Description: 下载文件函数
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\HttpRequest\Download.go
 */
package HttpRequest

import (
	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Tools"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/dop251/goja"
)

/**
 * @description: 下载文件函数
 * @param {string} URL 下载地址
 * @param {string} SavePath 保存路径
 * @param {goja.Callable} Func 回调函数
 * @return {*}
 */
func (HttpRequest HttpRequest) Download(URL string, SavePath string, Func goja.Callable) {
	go func() {
		PluginName := HttpRequest.VM.Get("PLUGIN_NAME").String()
		FilePath, FileSize, err := Tools.DownloadFile(URL, Controller.DataPath+"/"+PluginName+"/"+SavePath, false, 120)
		if err != nil {
			Log.Error("Plugin", err)
			Func(nil, HttpRequest.VM.ToValue(""), HttpRequest.VM.ToValue(0))
			return
		}
		Func(nil, HttpRequest.VM.ToValue(FilePath), HttpRequest.VM.ToValue(FileSize))
	}()
}
