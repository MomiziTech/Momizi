/*
 * @Author: NyanCatda
 * @Date: 2022-03-23 13:57:19
 * @LastEditTime: 2022-03-27 21:22:48
 * @LastEditors: NyanCatda
 * @Description:文件读写操作
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\File\ReadWrite.go
 */
package File

import (
	"os"

	Files "github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/dop251/goja"
)

/**
 * @description: 读取文件
 * @param {string} Path 文件路径
 * @return {string} 文件内容(读取失败返回nil)
 */
func (File File) Read(Path string) any {
	PluginName := File.VM.Get("PLUGIN_NAME").String()

	Files, err := Files.NewFileReadWrite(DataPath+PluginName+"/"+Path, os.O_RDONLY)
	if err != nil {
		return nil
	}
	defer Files.Close()

	Str, err := Files.Read()
	if err != nil {
		Log.Error(PluginName, err)
		return nil
	}
	return Str
}

/**
 * @description: 覆盖写入文件
 * @param {string} Path 文件路径
 * @param {string} Content 文件内容
 * @return {bool} 是否成功
 */
func (File File) WriteTo(Path string, Content string) bool {
	PluginName := File.VM.Get("PLUGIN_NAME").String()

	Files, err := Files.NewFileReadWrite(DataPath+PluginName+"/"+Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return false
	}
	defer Files.Close()

	err = Files.WriteTo(Content)
	if err != nil {
		Log.Error(PluginName, err)
		return false
	}
	return true
}

/**
 * @description: 追加写入文件
 * @param {string} Path 文件路径
 * @param {string} Content 文件内容
 * @return {bool} 是否成功
 */
func (File File) WriteAppend(Path string, Content string) bool {
	PluginName := File.VM.Get("PLUGIN_NAME").String()

	Files, err := Files.NewFileReadWrite(DataPath+PluginName+"/"+Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
	if err != nil {
		return false
	}
	defer Files.Close()

	err = Files.WriteAppend(Content)
	if err != nil {
		Log.Error(PluginName, err)
		return false
	}
	return true
}

/**
 * @description: 异步读取文件
 * @param {string} Path 文件路径
 * @param {goja.Callable} Callback 回调函数
 * @return {string} 文件内容(读取失败返回nil)
 */
func (File File) ReadAsync(Path string, Callback goja.Callable) {
	go func() {
		Callback(nil, File.VM.ToValue(File.Read(Path)))
	}()
}

/**
 * @description: 异步覆盖写入文件
 * @param {string} Path 文件路径
 * @param {string} Content 文件内容
 * @param {goja.Callable} Callback 回调函数
 * @return {bool} 是否成功
 */
func (File File) WriteToAsync(Path string, Content string, Callback goja.Callable) {
	go func() {
		Callback(nil, File.VM.ToValue(File.WriteTo(Path, Content)))
	}()
}

/**
 * @description: 异步追加写入文件
 * @param {string} Path 文件路径
 * @param {string} Content 文件内容
 * @param {goja.Callable} Callback 回调函数
 * @return {bool} 是否成功
 */
func (File File) WriteAppendAsync(Path string, Content string, Callback goja.Callable) {
	go func() {
		Callback(nil, File.VM.ToValue(File.WriteAppend(Path, Content)))
	}()
}
