/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:58:27
 * @LastEditTime: 2022-03-21 00:25:33
 * @LastEditors: NyanCatda
 * @Description: 插件加载模块
 * @FilePath: \Momizi\Controller\Plugin\Plugin.go
 */
package Plugin

import (
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Controller/Plugin/JavaScript"
	"github.com/MomiziTech/Momizi/Utils/Log"
)

/**
 * @description: 执行插件
 * @param {MessageStruct.MessageStruct} Message
 * @return {error} 错误信息
 */
func RunPlugin(Message MessageStruct.MessageStruct) error {
	// 运行JavaScript插件
	err := JavaScript.RunJavaScriptPlugin(Message)
	if err != nil {
		Log.ErrorWrite(err)
		return err
	}

	return err
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

	return nil
}
