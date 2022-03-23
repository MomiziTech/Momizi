/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 17:36:42
 * @LastEditTime: 2022-03-23 21:44:44
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
 * @param {[]string} Header 请求头
 * @param {string} SavePath 保存路径
 * @param {goja.Callable} Callback 回调函数
 * @return {string} FilePath 文件保存路径
 * @return {int64} FileSize 文件大小(Byte)
 */
func (HttpRequest HttpRequest) Download(URL string, Header []string, SavePath string, Callback goja.Callable) {
	go func() {
		PluginName := HttpRequest.VM.Get("PLUGIN_NAME").String()
		FilePath, FileSize, err := Tools.DownloadFile(URL, Header, Controller.DataPath+"/"+PluginName+"/"+SavePath, false, 120)
		if err != nil {
			PluginName := HttpRequest.VM.Get("PLUGIN_NAME").String()
			Log.Error(PluginName, err)
			Callback(nil, HttpRequest.VM.ToValue(""), HttpRequest.VM.ToValue(0))
			return
		}
		Callback(nil, HttpRequest.VM.ToValue(FilePath), HttpRequest.VM.ToValue(FileSize))
	}()
}
