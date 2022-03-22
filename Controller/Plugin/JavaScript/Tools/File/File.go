/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 21:57:52
 * @LastEditTime: 2022-03-22 22:43:33
 * @LastEditors: NyanCatda
 * @Description: 文件操作函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\File\File.go
 */
package File

import (
	Files "github.com/MomiziTech/Momizi/Utils/File"
	"github.com/MomiziTech/Momizi/Utils/Log"
	"github.com/dop251/goja"
)

var DataPath = "./data/"

type File struct {
	VM *goja.Runtime
}

/**
 * @description: 注册文件操作函数
 * @param {*goja.Runtime} VM
 * @return {*}
 */
func RegistrationFunction(VM *goja.Runtime) error {
	return VM.Set("File", File{VM: VM})
}

/**
 * @description: 删除文件/文件夹
 * @param {string} Path 路径
 * @return {bool} 是否成功
 */
func (File File) Delete(Path string) bool {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	OK, err := Files.Delete(DataPath + PluginName + "/" + Path)
	if err != nil {
		Log.Error("Plugin", err)
		return false
	}
	return OK
}

/**
 * @description: 判断文件/文件夹是否存在
 * @param {string} Path 路径
 * @return {*}
 */
func (File File) Exists(Path string) bool {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	return Files.Exists(DataPath + PluginName + "/" + Path)
}

/**
 * @description: 复制文件
 * @param {string} Path 路径
 * @param {string} NewPath 新路径
 * @return {int64} 文件大小(Byte)
 * @return {bool} 是否成功
 */
func (File File) Copy(Path string, NewPath string) (int64, bool) {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	FileSize, err := Files.Copy(DataPath+PluginName+"/"+Path, DataPath+PluginName+"/"+NewPath)
	if err != nil {
		Log.Error("Plugin", err)
		return 0, false
	}
	return FileSize, true
}

/**
 * @description: 移动文件
 * @param {string} Path 路径
 * @param {string} NewPath 新路径
 * @return {bool} 是否成功
 */
func (File File) Move(Path string, NewPath string) bool {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	err := Files.Move(DataPath+PluginName+"/"+Path, DataPath+PluginName+"/"+NewPath)
	if err != nil {
		Log.Error("Plugin", err)
		return false
	}
	return true
}

/**
 * @description: 获取文件大小
 * @param {string} FilePath 文件路径
 * @return {int64} 文件大小(Byte)
 */
func (File File) GetFileSize(FilePath string) int64 {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	FileSize, err := Files.GetFileSize(DataPath + PluginName + "/" + FilePath)
	if err != nil {
		Log.Error("Plugin", err)
		return 0
	}
	return FileSize
}
