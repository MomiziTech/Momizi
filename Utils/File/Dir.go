/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:26:55
 * @LastEditTime: 2022-03-22 23:21:11
 * @LastEditors: NyanCatda
 * @Description: 路径检查
 * @FilePath: \Momizi\Utils\File\Dir.go
 */
package File

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

/**
 * @description: 创建文件夹，如果不存在则创建
 * @param {string} path 文件夹路径
 * @return {*}
 */
func MKDir(Path string) (bool, error) {
	Path = filepath.Clean(Path)
	_, err := os.Stat(Path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(Path, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

/**
 * @description: 判断所给路径是否为文件夹
 * @param {string} Path 文件夹路径
 * @return {bool} 是否为文件夹
 */
func IsDir(Path string) bool {
	Path = filepath.Clean(Path)
	s, err := os.Stat(Path)
	if err != nil {
		return false
	}

	return s.IsDir()
}

/**
 * @description: 遍历目录下的所有文件(包含子目录)
 * @param {string} DirPth
 * @return {[]string} 文件路径列表
 * @return {error} 错误信息
 */
func GetFilesList(DirPth string) ([]string, error) {
	DirPth = filepath.Clean(DirPth)
	var dirs []string
	dir, err := ioutil.ReadDir(DirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	var Files []string
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, filepath.Clean(DirPth+PthSep+fi.Name()))
			GetFilesList(DirPth + PthSep + fi.Name())
		} else {
			Files = append(Files, filepath.Clean(DirPth+PthSep+fi.Name()))
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetFilesList(table)
		for _, temp1 := range temp {
			Files = append(Files, filepath.Clean(temp1))
		}
	}

	return Files, nil
}
