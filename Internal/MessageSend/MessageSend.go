/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:57:58
 * @LastEditTime: 2022-03-25 20:28:28
 * @LastEditors: NyanCatda
 * @Description: 消息发送模块
 * @FilePath: \Momizi\Internal\MessageSend\MessageSend.go
 */
package MessageSend

type MessageSend struct {
	ChatSoftware string
	ChatType     string
	ChatID       string
	ReplyID      string
}

/**
 * @description: 创建一个消息发送类
 * @param {string} ChatSoftware 聊天软件名称, QQ/Telegram/Line
 * @param {string} ChatType 聊天类型, User/Group
 * @param {string} ChatID 聊天ID
 * @param {string} ReplyID 回复ID(可选)
 * @return {MessageSend} 消息发送类
 */
func New(ChatSoftware string, ChatType string, ChatID string, ReplyID string) *MessageSend {
	return &MessageSend{ChatSoftware: ChatSoftware, ChatType: ChatType, ChatID: ChatID, ReplyID: ReplyID}
}
