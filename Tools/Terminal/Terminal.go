/*
 * @Author: NyanCatda
 * @Date: 2022-03-30 19:57:08
 * @LastEditTime: 2022-04-02 22:15:24
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
	"github.com/MomiziTech/Momizi/Tools/Terminal/Command"
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
		err = Command.Command(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
