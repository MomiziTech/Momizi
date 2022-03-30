/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:19:51
 * @LastEditTime: 2022-03-30 19:48:12
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Momizi\main.go
 */
package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/MomiziTech/Momizi/Internal/Controller/Initialization"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/ReceivingStart"
)

/**
 * @description: 主函数错误处理
 * @param {error} Error 错误信息
 * @return {*}
 */
func Error(Error error) {
	fmt.Println(Error.Error())
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

	// 等待所有goroutine结束
	wg.Wait()
}
