/*
 * @Author: NyanCatda
 * @Date: 2022-04-02 09:32:33
 * @LastEditTime: 2022-04-02 09:36:17
 * @LastEditors: NyanCatda
 * @Description: 音频消息发送封装
 * @FilePath: \Momizi\Internal\MessageSend\AudioMessage.go
 */
package MessageSend

import (
	"strconv"

	"github.com/MomiziTech/Momizi/Internal/MessageSend/ChatSoftwareAPI/Telegram"
	"github.com/MomiziTech/Momizi/Tools/Log"
)

func (MessageSend *MessageSend) Audio(Path string) error {
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
		MessageCallback, err := Chat.SendAudio(Path, "", "", nil, 0, "", "", false, false, ReplyID, true)

		// 如果消息发送成功，则记录日志
		if MessageCallback.MessageID != 0 {
			Log.SendMessage(MessageSend.ChatSoftware, MessageSend.ChatType, MessageSend.ChatID, Path)
		}
	}
	return nil
}
