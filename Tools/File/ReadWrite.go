/*
 * @Author: NyanCatda
 * @Date: 2022-03-23 13:59:28
 * @LastEditTime: 2022-03-23 14:33:18
 * @LastEditors: NyanCatda
 * @Description: 文件读写操作
 * @FilePath: \Momizi\Tools\File\ReadWrite.go
 */
package File

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

/**
 * @description: 读取文件
 * @param {string} Path 文件路径
 * @return {string} 文件内容
 * @return {error} 错误
 */
func Read(Path string) (string, error) {
	Path = filepath.Clean(Path)
	f, err := os.Open(Path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	Content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(Content), nil
}

/**
 * @description: 覆盖写入文件
 * @param {string} Path 文件路径
 * @param {string} Content 文件内容
 * @return {error} 错误
 */
func WriteTo(Path string, Content string) error {
	Path = filepath.Clean(Path)
	f, err := os.OpenFile(Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt([]byte(Content), n)
	if err != nil {
		return err
	}

	return err
}

/**
 * @description: 追加写入文件
 * @param {string} Path 文件路径
 * @param {string} Content 文件内容
 * @return {error} 错误
 */
func WriteAppend(Path string, Content string) error {
	Path = filepath.Clean(Path)
	file, err := os.OpenFile(Path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	_, err = write.WriteString(Content)
	if err != nil {
		return err
	}

	if err := write.Flush(); err != nil {
		return err
	}
	return nil
}
