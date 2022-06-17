/*
 * @Author: NyanCatda
 * @Date: 2022-03-21 19:51:14
 * @LastEditTime: 2022-06-17 23:41:24
 * @LastEditors: NyanCatda
 * @Description: 终端输出颜色
 * @FilePath: \Momizi\Tools\Log\Color.go
 */
package Log

import (
	"fmt"
	"regexp"
)

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func Black(msg string) string {
	return SetColor(msg, 0, 0, TextBlack)
}

func Red(msg string) string {
	return SetColor(msg, 0, 0, TextRed)
}

func Green(msg string) string {
	return SetColor(msg, 0, 0, TextGreen)
}

func Yellow(msg string) string {
	return SetColor(msg, 0, 0, TextYellow)
}

func Blue(msg string) string {
	return SetColor(msg, 0, 0, TextBlue)
}

func Magenta(msg string) string {
	return SetColor(msg, 0, 0, TextMagenta)
}

func Cyan(msg string) string {
	return SetColor(msg, 0, 0, TextCyan)
}

func White(msg string) string {
	return SetColor(msg, 0, 0, TextWhite)
}

func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}

/**
 * @description: 去除文字颜色
 * @param {string} msg 需要去除颜色的文字
 * @return {string} 去除颜色后的文字
 */
func DelColor(msg string) string {
	reg := regexp.MustCompile(`\x1b(\[.*?[@-~]|\].*?(\x07|\x1b\\))`)
	return reg.ReplaceAllString(msg, "")
}
