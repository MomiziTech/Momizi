/*
 * @Author: NyanCatda
 * @Date: 2022-04-04 12:14:09
 * @LastEditTime: 2022-04-04 12:42:36
 * @LastEditors: NyanCatda
 * @Description: 插件操作命令
 * @FilePath: \Momizi\Tools\Terminal\Command\Plugin.go
 */
package Command

import (
	Plugins "github.com/MomiziTech/Momizi/Internal/Plugin"
	"github.com/MomiziTech/Momizi/Internal/Plugin/Tools/PluginList"
	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 插件操作命令
 * @param {[]string} CommandParameters 命令参数
 * @return {*}
 */
func Plugin(CommandParameters []string) {
	if len(CommandParameters) <= 0 {
		Log.Info("System", "plugin <list/reload> 列出插件/重新加载插件")
		return
	}

	switch CommandParameters[0] {
	case "list":
		// 列出插件
		Log.Info("System", Log.Green("---------------- Plugin List ----------------"))
		for _, PluginInfo := range PluginList.PluginList {
			Log.Info("System", PluginInfo.Name, PluginInfo.Version, "Author:"+PluginInfo.Author)
		}
		Log.Info("System", Log.Green("---------------------------------------------"))
		return
	case "reload":
		// 重载插件
		PluginList.ClearPluginList() // 清空插件列表
		if err := Plugins.InitPlugin(); err != nil {
			Log.Error("System", err)
		}
		return
	}
}
