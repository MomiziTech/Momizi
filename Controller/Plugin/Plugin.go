/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:58:27
 * @LastEditTime: 2022-03-21 09:48:47
 * @LastEditors: NyanCatda
 * @Description: 插件加载模块
 * @FilePath: \Momizi\Controller\Plugin\Plugin.go
 */
package Plugin

import (
	"fmt"

	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript"
	"github.com/MomiziTech/Momizi/Utils/Log"
)

/**
 * @description: 运行插件消息监听器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {error} 错误信息
 */
func RunPluginMessageListener(Message MessageStruct.MessageStruct) error {
	// 运行JavaScript插件
	err := JavaScript.ExecutionMessageListener(Message)
	if err != nil {
		Log.ErrorWrite(err)
		return err
	}

	return nil
}

/**
 * @description: 初始化插件
 * @param {*}
 * @return {error} 错误信息
 */
func InitPlugin() error {
	// 初始化JavaScript插件
	err := JavaScript.InitJavaScriptPlugin()
	if err != nil {
		Log.ErrorWrite(err)
		return err
	}

	fmt.Println("插件加载完毕！")

	return nil
}
