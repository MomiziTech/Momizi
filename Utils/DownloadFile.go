/*
 * @Author: NyanCatda
 * @Date: 2022-03-12 23:31:49
 * @LastEditTime: 2022-03-19 16:18:17
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Momizi\Utils\DownloadFile.go
 */
package Utils

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

/**
 * @description: 下载文件
 * @param {string} URL 文件地址
 * @param {string} SavePath 保存路径
 * @param {int} timeOut 超时时间(秒)
 * @return {string} 文件保存路径
 * @param {bool} RandomFileName 是否生成随机文件名
 * @param {int} timeOut 超时时间(秒)
 * @return {string} 文件保存路径
 * @return {int64} 文件大小
 * @return {error} 错误信息
 */
func DownloadFile(URL string, SavePath string, RandomFileName bool, timeOut int) (string, int64, error) {
	//设置超时
	client := http.Client{
		Timeout: time.Duration(timeOut) * time.Second,
	}
	res, err := client.Get(URL)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	// 文件夹不存在则创建
	_, err = PathExists(SavePath)
	if err != nil {
		return "", 0, err
	}

	// 按照时间戳生成文件
	var FileName string
	if RandomFileName {
		Time := strconv.FormatInt(time.Now().Unix(), 10)
		FileSuffix := path.Ext(URL)
		FileName = Time + FileSuffix
	} else {
		FileName = path.Base(URL)
	}

	FilePath := path.Join(SavePath, FileName)
	file, err := os.Create(FilePath)
	if err != nil {
		return "", 0, err
	}
	writer := bufio.NewWriter(file)
	Size, err := io.Copy(writer, reader)
	if err != nil {
		return "", 0, err
	}

	return FilePath, Size, nil
}
