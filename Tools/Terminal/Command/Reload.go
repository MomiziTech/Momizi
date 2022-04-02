/*
 * @Author: NyanCatda
 * @Date: 2022-04-02 22:30:40
 * @LastEditTime: 2022-04-02 23:24:34
 * @LastEditors: NyanCatda
 * @Description: 重载命令
 * @FilePath: \Momizi\Tools\Terminal\Command\Reload.go
 */
package Command

import (
	"github.com/MomiziTech/Momizi/Internal/Plugin"
	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 重载插件
 * @param {[]string} CommandParameters 命令参数
 * @return {*}
 */
func Reload(CommandParameters []string) {
	// 加载插件
	if err := Plugin.InitPlugin(); err != nil {
		Log.Error("System", err)
	}
}
