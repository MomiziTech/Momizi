/*
 * @Author: NyanCatda
 * @Date: 2022-03-25 20:19:59
 * @LastEditTime: 2022-03-25 21:06:41
 * @LastEditors: NyanCatda
 * @Description: 文本消息发送模块
 * @FilePath: \Momizi\Internal\MessageSend\TextMessage.go
 */
package MessageSend

import (
	"strconv"

	"github.com/MomiziTech/Momizi/Internal/MessageSend/ChatSoftwareAPI/Telegram"
	"github.com/MomiziTech/Momizi/Tools/Log"
)

/**
 * @description: 发送文本消息
 * @param {string} Content 消息内容
 * @return {*}
 */
func (MessageSend *MessageSend) Text(Content string) error {
	switch MessageSend.ChatSoftware {
	case "Telegram":
		// 组成消息内容
		ChatID, err := strconv.Atoi(MessageSend.ChatID)
		if err != nil {
			return err
		}
		Chat := Telegram.NewChat(ChatID)
		ReplyID, err := strconv.Atoi(MessageSend.ReplyID)
		if err != nil {
			return err
		}

		// 发送消息
		MessageCallback, err := Chat.SendMessage(Content, "", nil, false, false, false, ReplyID, true)
		if err != nil {
			return err
		}

		// 如果消息发送成功，则记录日志
		if MessageCallback.MessageID != 0 {
			Log.SendMessage(MessageSend.ChatSoftware, MessageSend.ChatType, MessageSend.ChatID, Content)
		}
	}

	return nil
}
