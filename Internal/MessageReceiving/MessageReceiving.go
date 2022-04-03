/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:57:36
 * @LastEditTime: 2022-03-28 14:59:23
 * @LastEditors: NyanCatda
 * @Description: 消息接收模块
 * @FilePath: \Momizi\Internal\MessageReceiving\MessageReceiving.go
 */
package MessageReceiving

import (
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/WebHook"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/WebHook/Struct"
	"github.com/MomiziTech/Momizi/Internal/Plugin"
	"github.com/MomiziTech/Momizi/Tools/Log"
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
		Message, err = WebHook.Line(JsonBody)
		if err != nil {
			return err
		}
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
	var ChatID string
	if Message.Type == "Group" {
		ChatID = Message.Sender.Group.ID
	} else {
		ChatID = Message.Sender.ID
	}

	var Content string
	for Num, MessageChain := range Message.MessageChain {
		switch MessageChain.Type {
		case "Text":
			Content = Message.MessageChain[Num].Text
		case "Image":
			Content = MessageChain.File.Name
		case "Audio":
			Content = MessageChain.File.Name
		case "File":
			Content = MessageChain.File.Name
		}
		Log.ReceivedMessage(Message.ChatSoftware, Message.Type, ChatID, Message.Sender.Username, Message.Sender.ID, Content)
	}
}
