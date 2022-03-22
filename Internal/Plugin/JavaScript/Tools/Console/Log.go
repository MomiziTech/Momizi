/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 01:28:05
 * @LastEditTime: 2022-03-22 17:59:24
 * @LastEditors: NyanCatda
 * @Description: 日志输出函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScript\Tools\Console\Log.go
 */
package Console

import (
	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 打印Log
 * @param {string} Text 日志内容
 * @return {*}
 */
func (Console Console) Log(Text string) error {
	return logDefault(Console, Log.INFO, Text)
}

/**
 * @description: 打印警告Log
 * @param {string} Text 日志内容
 * @return {*}
 */
func (Console Console) Warning(Text string) error {
	return logDefault(Console, Log.WARNING, Text)
}

/**
 * @description: 打印错误Log
 * @param {string} Text 日志内容
 * @return {*}
 */
func (Console Console) Error(Text string) error {
	return logDefault(Console, Log.ERROR, Text)
}

/**
 * @description: 打印Debug Log
 * @param {string} Text 日志内容F
 * @return {*}
 */
func (Console Console) Debug(Text string) error {
	return logDefault(Console, Log.DEBUG, Text)
}

/**
 * @description: 打印Log
 * @param {Console} Console 控制台对象
 * @param {int} Level 日志等级
 * @param {string} Text 日志内容
 * @return {*}
 */
func logDefault(Console Console, Level int, Text string) error {
	return Log.Print(Console.VM.Get("PLUGIN_NAME").String(), Level, Text)
}

/**
 * @description: 设置输出文字颜色
 * @param {string} Color 颜色 Black/Red/Green/Yellow/Blue/Magenta/Cyan/White
 * @param {string} Text 文字
 * @return {*}
 */
func (Console Console) SetColor(Color string, Text string) string {
	switch Color {
	case "Black":
		return Log.Black(Text)
	case "Red":
		return Log.Red(Text)
	case "Green":
		return Log.Green(Text)
	case "Yellow":
		return Log.Yellow(Text)
	case "Blue":
		return Log.Blue(Text)
	case "Magenta":
		return Log.Magenta(Text)
	case "Cyan":
		return Log.Cyan(Text)
	case "White":
		return Log.White(Text)
	default:
		return Text
	}
}
