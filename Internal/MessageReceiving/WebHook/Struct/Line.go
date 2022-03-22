/*
 * @Author: NyanCatda
 * @Date: 2022-03-09 14:21:18
 * @LastEditTime: 2022-03-09 15:41:50
 * @LastEditors: NyanCatda
 * @Description: Line WebHook结构体
 * @FilePath: \Momizi\Internal\MessageReceiving\WebHookStruct\Line.go
 */
package Struct

type Line struct {
	Destination string `json:"destination"` // 消息接收者
	Events      []struct {
		Type    string `json:"type"` // 消息类型，可能类型："message"，
		Message struct {
			Type string `json:"type"` // 消息内容类型，可能类型："text"，"image"，"video"，"audio"，"file"，"sticker"
			ID   string `json:"id"`   // 消息ID(文件需要通过此id获得)

			// 文本消息
			Text string `json:"text"` // 消息内容

			// 文件消息
			FileName string `json:"fileName"` // 文件名
			FileSize int    `json:"fileSize"` // 文件大小(Byte)

			// 语音消息
			Duration int `json:"duration"` // 语音时长(秒)

			ContentProvider struct {
				Type string `json:"type"` // 消息内容提供者类型，可能类型："line"，"external"
			} `json:"contentProvider"`
		} `json:"message"`
		Timestamp int64 `json:"timestamp"` // 消息时间戳
		Source    struct {
			Type    string `json:"type"`    // 消息发送者类型，可能类型："user"，"group"，"room"
			GroupID string `json:"groupId"` // 消息发送群ID
			UserID  string `json:"userId"`  // 消息发送者ID
		} `json:"source"`
		ReplyToken string `json:"replyToken"` // 消息回复Token
		Mode       string `json:"mode"`       // 消息模式，可能模式："active"，"standby"
	} `json:"events"`
}
