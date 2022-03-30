/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:26:02
 * @LastEditTime: 2022-03-30 20:40:41
 * @LastEditors: NyanCatda
 * @Description: 日志模块
 * @FilePath: \Momizi\Tools\Log\LogFile.go
 */

package Log

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Tools/File"
)

var (
	LogPath    = Controller.LogPath + "/"
	ColorPrint bool
)

const (
	INFO = iota + 0
	WARNING
	ERROR
	DEBUG
)

/**
 * @description: 打印错误
 * @param {string} Source 日志来源
 * @param {error} Error 错误信息
 * @return {*}
 */
func Error(Source string, Error error) {
	// 打印错误
	Print(Source, ERROR, Error)
}

/**
 * @description: 打印警告
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func Warning(Source string, Text ...any) {
	// 打印警告
	Print(Source, WARNING, Text...)
}

/**
 * @description: 打印信息
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func Info(Source string, Text ...any) {
	Print(Source, INFO, Text...)
}

/**
 * @description: 打印DeBug错误
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func DeBug(Source string, Text ...any) {
	Print(Source, DEBUG, Text...)
}

/**
 * @description: 打印发送的消息
 * @param {string} ChatSoftware 聊天软件, QQ/Telegram/Line
 * @param {string} ChatType 聊天类型, User/Group/Other
 * @param {string} ChatID 聊天ID
 * @param {string} Content 消息内容
 * @return {*}
 */
func SendMessage(ChatSoftware string, ChatType string, ChatID string, Content string) {
	LogText := fmt.Sprintf("%s: %s[%s] <- %s", ChatSoftware, ChatType, ChatID, Content)

	Print("Message", INFO, LogText)
}

/**
 * @description: 打印接收到的消息
 * @param {string} ChatSoftware 聊天软件, QQ/Telegram/Line
 * @param {string} ChatType 聊天类型, User/Group/Other
 * @param {string} ChatID 聊天ID
 * @param {string} UserName 用户名
 * @param {string} SenderID 发送者ID
 * @param {string} Content 消息内容
 * @return {*}
 */
func ReceivedMessage(ChatSoftware string, ChatType string, ChatID string, UserName string, SenderID string, Content string) {
	LogText := fmt.Sprintf("%s: %s[%s] %s(%s) -> %s", ChatSoftware, ChatType, ChatID, UserName, SenderID, Content)

	Print("Message", INFO, LogText)
}

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

	Text = append([]any{"\r" + Cyan(NowTime), LevelStr, Source}, Text...)

	// 如果彩色输出被关闭
	var LogText []any
	if !ColorPrint {
		// 遍历消息内容去除颜色
		for _, v := range Text {
			DelColorText := DelColor(fmt.Sprint(v))
			LogText = append(LogText, DelColorText)
		}
	} else {
		LogText = Text
	}

	// 打印日志
	_, err := fmt.Println(LogText...)
	if err != nil {
		return err
	}
	fmt.Print(Controller.TerminalPrompt)

	// 写入日志
	logFile, err := LogFile()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer logFile.Close()
	write := bufio.NewWriter(logFile)

	// 遍历消息内容去除颜色
	var LogFileText string
	for _, v := range Text {
		DelColorText := DelColor(fmt.Sprint(v))
		LogFileText += DelColorText
		LogFileText += " "
	}

	write.WriteString(LogFileText + "\n")
	write.Flush()

	return nil
}

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
