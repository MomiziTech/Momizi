/*
 * @Author: NyanCatda
 * @Date: 2022-04-02 22:18:50
 * @LastEditTime: 2022-04-03 13:12:09
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
	for _, CommandInfo := range CommandList {
		Log.Info("System", CommandInfo.Command+" "+CommandInfo.Help)
	}
	Log.Info("System", Log.Green("--------------------------------------"))
}
