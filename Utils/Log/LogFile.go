/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:26:02
 * @LastEditTime: 2022-03-21 10:21:57
 * @LastEditors: NyanCatda
 * @Description: 日志模块
 * @FilePath: \Momizi\Utils\Log\LogFile.go
 */

package Log

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/MomiziTech/Momizi/Utils"
)

var (
	LogPath = "./logs/"
)

/**
 * @description: 读写Log文件，按天分割日志
 * @param {*}
 * @return {*os.File}
 * @return {error}
 */
func LogFile() (*os.File, error) {
	// 判断文件夹是否存在
	Utils.PathExists(LogPath)

	logFileName := time.Now().Format("2006-01-02") + ".log"

	logfile, err := os.OpenFile(LogPath+logFileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果文件不存在则创建
		logfile, err := os.Create(LogPath + logFileName)
		if err != nil {
			return logfile, err
		}
		return logfile, nil
	}

	return logfile, nil
}

/**
 * @description: 将错误写入日志
 * @param {error} Error 错误信息
 * @return {*}
 */
func ErrorWrite(Source string, Error error) {
	// 打印错误
	Print(Source, "ERROR", Error)
	// 将错误写入日志
	logFile, err := LogFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close()
	write := bufio.NewWriter(logFile)

	write.WriteString("Error: " + Error.Error() + "\n")
	write.Flush()
}

/**
 * @description:  标准日志打印
 * @param {string} Source 日志来源
 * @param {string} level 日志等级 INFO/WARNING/ERROR
 * @param {...interface{}} Text 日志内容
 * @return {*}
 */
func Print(Source string, level string, Text ...interface{}) error {
	NowTime := time.Now().Format("2006-01-02 15:04:05")

	// Source拼接
	Source = "[" + Source + "]"

	// 打印日志
	Text = append([]interface{}{NowTime, level, Source}, Text...)
	_, err := fmt.Println(Text...)
	if err != nil {
		return err
	}
	return nil
}
