/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 22:10:18
 * @LastEditTime: 2022-03-22 23:21:06
 * @LastEditors: NyanCatda
 * @Description: 文件操作
 * @FilePath: \Momizi\Tools\File\File.go
 */
package File

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/**
 * @description: 删除文件/文件夹
 * @param {string} Path 文件/文件夹路径
 * @return {bool} 是否删除成功
 * @return {error} 错误信息
 */
func Delete(Path string) (bool, error) {
	Path = filepath.Clean(Path)
	if Exists(Path) {
		if IsDir(Path) {
			err := os.RemoveAll(Path)
			if err != nil {
				return false, err
			}
		}
		if IsFile(Path) {
			err := os.Remove(Path)
			if err != nil {
				return false, err
			}
		}
		return true, nil
	}
	return false, errors.New("文件不存在")
}

/**
 * @description: 判断所给路径是否为文件
 * @param {string} Path
 * @return {bool} 是否为文件
 */
func IsFile(Path string) bool {
	Path = filepath.Clean(Path)
	return !IsDir(Path)
}

/**
 * @description: 判断所给路径文件/文件夹是否存在
 * @param {string} Path 文件/文件夹路径
 * @return {bool} 是否存在
 */
func Exists(Path string) bool {
	Path = filepath.Clean(Path)
	_, err := os.Stat(Path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
 * @description: 复制文件
 * @param {string} Path 源文件路径
 * @param {string} NewPath 目标文件路径
 * @return {int64} 文件大小(Byte)
 * @return {error} 错误信息
 */
func Copy(Path string, NewPath string) (int64, error) {
	Path = filepath.Clean(Path)
	NewPath = filepath.Clean(NewPath)
	sourceFileStat, err := os.Stat(Path)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s不是常规文件", Path)
	}

	source, err := os.Open(Path)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(NewPath)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

/**
 * @description: 移动文件
 * @param {string} Path 源文件路径
 * @param {string} NewPath 目标文件路径
 * @return {*}
 */
func Move(Path string, NewPath string) error {
	Path = filepath.Clean(Path)
	NewPath = filepath.Clean(NewPath)
	if Exists(Path) {
		err := os.Rename(Path, NewPath)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("文件不存在")
}

/**
 * @description: 获取文件大小
 * @param {string} FilePath 文件路径
 * @return {int64} 文件大小(Byte)
 * @return {error} 错误信息
 */
func GetFileSize(FilePath string) (int64, error) {
	FilePath = filepath.Clean(FilePath)
	if Exists(FilePath) {
		File, err := os.Stat(FilePath)
		if err != nil {
			return 0, err
		}
		return File.Size(), nil
	}
	return 0, errors.New("文件不存在")
}
