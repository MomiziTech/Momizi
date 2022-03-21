/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:57:36
 * @LastEditTime: 2022-03-21 21:24:33
 * @LastEditors: NyanCatda
 * @Description: 消息接收模块
 * @FilePath: \Momizi\Controller\MessageReceiving\MessageReceiving.go
 */
package MessageReceiving

import (
	"fmt"

	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/WebHook"
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/WebHook/Struct"
	"github.com/MomiziTech/Momizi/Controller/Plugin"
	"github.com/MomiziTech/Momizi/Utils/Log"
	"github.com/gin-gonic/gin"
)

var (
	// 消息
	Message MessageStruct.MessageStruct
)

/**
 * @description: 消息接收处理
 * @param {*gin.Context} c
 * @return {*}
 */
func MessageReceiving(c *gin.Context) error {
	// 解析得到的Json
	var JsonBody Struct.WebHook
	if err := c.ShouldBindJSON(&JsonBody); err != nil {
		return err
	}

	var err error

	// 判断消息来源
	if JsonBody.Telegram.UpdateID != 0 {
		// 消息为Telegram消息
		Message, err = WebHook.Telegram(JsonBody)
		if err != nil {
			return err
		}
	}
	if JsonBody.Mirai.Type != "" {
		// 消息为Mirai消息
		Message, err = WebHook.Mirai(JsonBody)
		if err != nil {
			return err
		}
	}
	if JsonBody.Line.Destination != "" {
		// 消息为Line消息
		fmt.Println(JsonBody)
	}

	if Message.ID != "" {
		// 打印消息内容
		PrintMessage(Message)

		// 将消息传递给插件
		Plugin.RunPluginMessageListener(Message)
	}

	return nil
}

/**
 * @description: 打印消息内容
 * @param {MessageStruct.MessageStruct} Message 消息结构体
 * @return {*}
 */
func PrintMessage(Message MessageStruct.MessageStruct) {
	LogHeader := Message.ChatSoftware + ": " + Message.Sender.Username + "(" + Message.Sender.ID + ") -> "
	for Num, MessageChain := range Message.MessageChain {
		switch MessageChain.Type {
		case "Text":
			LogText := LogHeader + Message.MessageChain[Num].Text
			Log.Print("Message", Log.INFO, LogText)
		case "Image":
			LogText := LogHeader + MessageChain.File.Name
			Log.Print("Message", Log.INFO, LogText)
		case "Audio":
			LogText := LogHeader + MessageChain.File.Name
			Log.Print("Message", Log.INFO, LogText)
		case "Video":
			LogText := LogHeader + MessageChain.File.Name
			Log.Print("Message", Log.INFO, LogText)
		}
	}
}
