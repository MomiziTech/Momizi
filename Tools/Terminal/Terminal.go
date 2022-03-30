/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:57:08
 * @LastEditTime: 2022-03-30 20:40:28
 * @LastEditors: NyanCatda
 * @Description: 终端控制台
 * @FilePath: \Momizi\Tools\Terminal\Terminal.go
 */
package Terminal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MomiziTech/Momizi/Internal/Controller"
)

/**
 * @description: 启动控制台
 * @param {*}
 * @return {*}
 */
func Start() {
	Reader := bufio.NewReader(os.Stdin)
	// 循环处理输入
	for {
		fmt.Printf(Controller.TerminalPrompt)
		// 获取输入
		cmdString, err := Reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// 执行命令
		err = Command(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
