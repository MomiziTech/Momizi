/*
 * @Author: NyanCatda
 * @Date: 2022-04-02 22:18:50
 * @LastEditTime: 2022-04-02 22:22:39
 * @LastEditors: NyanCatda
 * @Description: 帮助命令
 * @FilePath: \Momizi\Tools\Terminal\Command\Help.go
 */
package Command

import (
	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 帮助命令
 * @param {[]string} CommandParameters 命令参数
 * @return {*}
 */
func Help(CommandParameters []string) {
	Log.Info("System", Log.Green("---------------- Help ----------------"))
	// 遍历帮助信息列表
	for _, Help := range HelpList {
		for Command, HelpInfo := range Help {
			Log.Info("System", Command+":"+HelpInfo)
		}
	}
	Log.Info("System", Log.Green("--------------------------------------"))
}
