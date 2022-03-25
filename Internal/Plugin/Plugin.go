/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:58:27
 * @LastEditTime: 2022-03-25 22:50:03
 * @LastEditors: NyanCatda
 * @Description: 插件加载模块
 * @FilePath: \Momizi\Internal\Plugin\Plugin.go
 */
package Plugin

import (
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	JavascriptV8 "github.com/MomiziTech/Momizi/Internal/Plugin/JavaScriptV8"
	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 运行插件消息监听器
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {error} 错误信息
 */
func RunPluginMessageListener(Message MessageStruct.MessageStruct) error {
	// 运行JavaScript插件
	err := JavascriptV8.ExecutionMessageListener(Message)
	if err != nil {
		Log.Error("Plugin", err)
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
	err := JavascriptV8.InitJavaScriptPlugin()
	if err != nil {
		Log.Error("Plugin", err)
		return err
	}

	Log.Print("Plugin", Log.INFO, "插件加载完成")

	return nil
}
