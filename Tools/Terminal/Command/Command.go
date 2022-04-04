/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:57:58
 * @LastEditTime: 2022-04-04 12:28:51
 * @LastEditors: NyanCatda
 * @Description: 命令处理
 * @FilePath: \Momizi\Tools\Terminal\Command\Command.go
 */
package Command

import (
	"strings"

	"github.com/MomiziTech/Momizi/Tools/Log"
)

type CommandFunc func([]string)

type CommandInfo struct {
	Command  string
	Help     string
	Callback CommandFunc
}

var CommandList []CommandInfo

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

	// 匹配命令
	for _, Command := range CommandList {
		if Command.Command == arrCommandStr[0] {
			if Command.Callback != nil {
				// 获取命令参数
				CommandParameters := arrCommandStr[1:]

				// 执行命令回调
				Command.Callback(CommandParameters)
				return nil
			}
		}
	}

	// 如果没有匹配到任何命令则返回命令不存在
	Log.Info("System", arrCommandStr[0]+"命令不存在，请输入help查看帮助")

	return nil
}

/**
 * @description: 添加命令
 * @param {string} Command 命令
 * @param {string} Help 帮助信息
 * @param {CommandFunc} Callback 回调函数
 * @return {*}
 */
func AddCommand(Command string, Help string, Callback CommandFunc) {
	CommandList = append(CommandList, CommandInfo{Command: Command, Help: Help, Callback: Callback})
}

/**
 * @description: 初始化默认命令列表
 * @param {*}
 * @return {*}
 */
func InitCommandList() {
	AddCommand("help", "查看帮助", Help)
	AddCommand("exit", "退出程序", Exit)
	AddCommand("plugin", "<list/reload> 插件命令", Plugin)
}
