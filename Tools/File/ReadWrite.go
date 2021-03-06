/*
 * @Author: NyanCatda
 * @Date: 2022-03-23 13:59:28
 * @LastEditTime: 2022-04-02 02:57:34
 * @LastEditors: McPlus
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

type ReadWrite struct {
	Path string
	File *os.File
}

/**
 * @description: 新建文件读写类
 * @param {string} Path 文件路径
 * @return {string} 文件内容
 */
func NewFileReadWrite(Path string, Flag int) (*ReadWrite, error) {
	Path = filepath.Clean(Path)
	File, err := os.OpenFile(Path, Flag, 0666)
	if err != nil {
		return nil, err
	}
	return &ReadWrite{Path, File}, nil
}

/**
 * @description: 读取文件(需要权限os.O_RDONLY)
 * @return {string} 文件内容
 * @return {error} 错误
 */
func (ReadWrite *ReadWrite) Read() (string, error) {
	Content, err := ioutil.ReadAll(ReadWrite.File)
	if err != nil {
		return "", err
	}

	return string(Content), nil
}

/**
 * @description: 覆盖写入文件(需要权限os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
 * @param {string} Content 文件内容
 * @return {error} 错误
 */
func (ReadWrite *ReadWrite) WriteTo(Content string) error {
	n, _ := ReadWrite.File.Seek(0, os.SEEK_END)
	_, err := ReadWrite.File.WriteAt([]byte(Content), n)
	if err != nil {
		return err
	}

	return err
}

/**
 * @description: 追加写入文件(需要权限os.O_WRONLY|os.O_APPEND|os.O_CREATE)
 * @param {string} Content 文件内容
 * @return {error} 错误
 */
func (ReadWrite *ReadWrite) WriteAppend(Content string) error {
	write := bufio.NewWriter(ReadWrite.File)
	_, err := write.WriteString(Content)
	if err != nil {
		return err
	}

	if err := write.Flush(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: 关闭文件
 * @param {*}
 * @return {error} 错误
 */
func (ReadWrite *ReadWrite) Close() error {
	return ReadWrite.File.Close()
}
