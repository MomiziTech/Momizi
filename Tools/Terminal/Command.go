/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:57:58
 * @LastEditTime: 2022-04-02 22:09:27
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

type CommandFunc func()

var CommandList []map[string]CommandFunc

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
	}

	// 如果没有匹配到默认命令，则尝试匹配自定义命令
	for _, Command := range CommandList {
		if Command[arrCommandStr[0]] != nil {
			Command[arrCommandStr[0]]()
			return nil
		}
	}

	// 如果没有匹配到任何命令则返回命令不存在
	Log.Info("System", arrCommandStr[0]+"命令不存在，请输入help查看帮助")

	return nil
}

/**
 * @description: 添加命令(优先级低于默认命令)
 * @param {string} Command 命令
 * @param {CommandFunc} Callback 回调函数
 * @return {*}
 */
func AddCommand(Command string, Callback CommandFunc) {
	CommandList = append(CommandList, map[string]CommandFunc{Command: Callback})
}
