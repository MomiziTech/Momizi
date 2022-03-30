/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:57:58
 * @LastEditTime: 2022-03-30 20:46:30
 * @LastEditors: NyanCatda
 * @Description: 命令处理
 * @FilePath: \Momizi\Tools\Terminal\Command.go
 */
package Terminal

import (
	"os"
	"strings"

	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 命令处理
 * @param {string} commandStr
 * @return {*}
 */
func Command(CommandStr string) error {
	// 分割命令
	CommandStr = strings.TrimSuffix(CommandStr, "\n")
	arrCommandStr := strings.Fields(CommandStr)
	// 判断命令是否为空
	if len(arrCommandStr) == 0 {
		return nil
	}
	// 执行命令
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "help":
		Log.Info("System", Log.Green("---------------- Help ----------------"))
		Log.Info("System", "exit: 退出程序")
		Log.Info("System", Log.Green("--------------------------------------"))
	default:
		Log.Info("System", arrCommandStr[0]+"命令不存在，请输入help查看帮助")
	}
	return nil
}
