/*
 * @Author: NyanCatda
 * @Date: 2022-04-02 22:17:05
 * @LastEditTime: 2022-04-02 23:24:13
 * @LastEditors: NyanCatda
 * @Description: 退出命令
 * @FilePath: \Momizi\Tools\Terminal\Command\Exit.go
 */
package Command

import "os"

/**
 * @description: 退出命令
 * @param {[]string} CommandParameters 命令参数
 * @return {*}
 */
func Exit(CommandParameters []string) {
	os.Exit(0)
}
