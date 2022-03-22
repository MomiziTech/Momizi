/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:26:02
 * @LastEditTime: 2022-03-22 22:07:53
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

	"github.com/MomiziTech/Momizi/Utils/File"
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
	File.MKDir(LogPath)

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
 * @description: 打印错误
 * @param {error} Error 错误信息
 * @return {*}
 */
func Error(Source string, Error error) {
	// 打印错误
	Print(Source, ERROR, Error)
}

/**
 * @description: 打印警告
 * @param {*}
 * @return {*}
 */
func Warning(Source string, Text ...any) {
	// 打印警告
	Print(Source, WARNING, Text...)
}

const (
	INFO = iota + 0
	WARNING
	ERROR
	DEBUG
)

/**
 * @description:  标准日志打印
 * @param {string} Source 日志来源
 * @param {string} Level 日志等级 INFO/WARNING/ERROR/DEBUG
 * @param {...any} Text 日志内容
 * @return {*}
 */
func Print(Source string, Level int, Text ...any) error {
	NowTime := time.Now().Format("2006-01-02 15:04:05")

	// Source拼接
	Source = "[" + Source + "]"

	// 判断level颜色
	var LevelStr string
	switch Level {
	case 0:
		LevelStr = Blue("INFO")
	case 1:
		LevelStr = Yellow("WARNING")
	case 2:
		LevelStr = Red("ERROR")
	case 3:
		LevelStr = Green("DEBUG")
	default:
		LevelStr = Magenta("Other")
	}

	// 打印日志
	Text = append([]any{Cyan(NowTime), LevelStr, Source}, Text...)
	_, err := fmt.Println(Text...)
	if err != nil {
		return err
	}

	// 写入日志
	logFile, err := LogFile()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer logFile.Close()
	write := bufio.NewWriter(logFile)

	// 遍历消息内容去除颜色
	var LogText string
	for _, v := range Text {
		DelColorText := DelColor(fmt.Sprint(v))
		LogText += DelColorText
		LogText += " "
	}

	write.WriteString(LogText + "\n")
	write.Flush()

	return nil
}
