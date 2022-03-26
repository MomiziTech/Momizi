/*
 * @Author: NyanCatda
 * @Date: 2022-03-12 23:31:49
 * @LastEditTime: 2022-03-27 02:43:45
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Momizi\Tools\DownloadFile.go
 */
package Tools

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/MomiziTech/Momizi/Tools/File"
)

/**
 * @description: 下载文件
 * @param {string} URL 文件地址
 * @param {string} Header 请求头
 * @param {string} SavePath 保存路径
 * @param {int} timeOut 超时时间(秒)
 * @return {string} 文件保存路径
 * @param {bool} RandomFileName 是否生成随机文件名
 * @param {int} timeOut 超时时间(秒)
 * @return {string} 文件保存路径
 * @return {int64} 文件大小
 * @return {error} 错误信息
 */
func DownloadFile(URL string, Header []string, SavePath string, RandomFileName bool, timeOut int) (string, int64, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", 0, err
	}

	// 设置请求头
	for _, value := range Header {
		Headervalue := strings.Split(value, ":")
		// 如果解析失败则不设置请求头
		if len(Headervalue) <= 0 {
			break
		}
		req.Header.Set(Headervalue[0], Headervalue[1])
	}

	client := http.Client{
		Timeout: time.Duration(timeOut) * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	// 规整路径
	SavePath = filepath.Clean(SavePath)

	// 文件夹不存在则创建
	_, err = File.MKDir(SavePath)
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
	defer file.Close()
	Size, err := io.Copy(writer, reader)
	if err != nil {
		return "", 0, err
	}

	return filepath.Clean(FilePath), Size, nil
}
