/*
 * @Author: NyanCatda
 * @Date: 2022-03-22 01:28:05
 * @LastEditTime: 2022-03-22 17:33:19
 * @LastEditors: NyanCatda
 * @Description: 日志输出函数注册
 * @FilePath: \Momizi\Controller\Plugin\JavaScript\Tools\Console\Log.go
 */
package Console

import (
	"github.com/MomiziTech/Momizi/Utils/Log"
)

func (Console Console) Log(Text string) error {
	return logDefault(Console, Log.INFO, Text)
}

func (Console Console) Warning(Text string) error {
	return logDefault(Console, Log.WARNING, Text)
}

func (Console Console) Error(Text string) error {
	return logDefault(Console, Log.ERROR, Text)
}

func (Console Console) Debug(Text string) error {
	return logDefault(Console, Log.DEBUG, Text)
}

func logDefault(Console Console, Level int, Text string) error {
	return Log.Print(Console.VM.Get("PLUGIN_NAME").String(), Level, Text)
}

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
