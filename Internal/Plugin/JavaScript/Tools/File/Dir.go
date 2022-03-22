/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 21:59:31
 * @LastEditTime: 2022-03-22 23:16:55
 * @LastEditors: NyanCatda
 * @Description: 目录操作
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\File\Dir.go
 */
package File

import (
	"path/filepath"
	"strings"

	Files "github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 创建文件夹
 * @param {string} Path
 * @return {*}
 */
func (File File) MKDir(Path string) bool {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	_, err := Files.MKDir(DataPath + PluginName + "/" + Path)
	if err != nil {
		return false
	}
	return true
}

/**
 * @description: 判断是否是文件夹
 * @param {string} Path 路径
 * @return {bool} 是否是文件夹
 */
func (File File) IsDir(Path string) bool {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	return Files.IsDir(DataPath + PluginName + "/" + Path)
}

// 列出文件夹下的全部文件(包含子目录)
func (File File) GetFilesList(Path string) []string {
	PluginName := File.VM.Get("PLUGIN_NAME").String()
	Files, err := Files.GetFilesList(DataPath + PluginName + "/" + Path)
	if err != nil {
		Log.Error("Plugin", err)
		return nil
	}

	// 去除前缀路径
	for i := 0; i < len(Files); i++ {
		Files[i] = strings.TrimPrefix(Files[i], filepath.Clean(DataPath+PluginName+"/"))
	}

	return Files
}
