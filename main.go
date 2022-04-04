/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:19:51
 * @LastEditTime: 2022-04-04 13:02:56
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Momizi\main.go
 */
package main

import (
	"os"
	"sync"

	"github.com/MomiziTech/Momizi/Internal/Controller/Initialization"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/ReceivingStart"
	"github.com/MomiziTech/Momizi/Tools/Log"
	"github.com/MomiziTech/Momizi/Tools/Terminal"
)

/**
 * @description: 主函数错误处理
 * @param {error} Error 错误信息
 * @return {*}
 */
func Error(Error error) {
	Log.Error("System", Error)
	key := make([]byte, 1)
	os.Stdin.Read(key)
	os.Exit(1)
}

func main() {
	var wg sync.WaitGroup
	// 初始化程序
	FlagConfig, err := Initialization.Initialization()
	if err != nil {
		Error(err)
	}

	// 启动WebHook接收
	wg.Add(1)
	go ReceivingStart.Run(FlagConfig.RunMode)

	// 启动控制台
	Terminal.Start()

	// 等待所有goroutine结束
	wg.Wait()
}
