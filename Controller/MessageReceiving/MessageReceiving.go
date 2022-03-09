/*
 * @Author: NyanCatda
 * @Date: 2022-03-08 21:57:36
 * @LastEditTime: 2022-03-09 16:50:25
 * @LastEditors: NyanCatda
 * @Description: 消息接收模块
 * @FilePath: \Momizi\Controller\MessageReceiving\MessageReceiving.go
 */
package MessageReceiving

import (
	"fmt"

	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/WebHookStruct"
	"github.com/gin-gonic/gin"
)

var (
	// 消息
	Message MessageStruct
)

type MessageStruct struct {
	ID           string `json:"id"`            // 消息ID
	Type         string `json:"type"`          // 消息来源类型，User, Group
	ChatSoftware string `json:"chat_software"` // 消息来源软件
	Time         int    `json:"time"`          // 消息接收时间戳
	// 消息链
	MessageChain []struct {
		Type string `json:"type"` // 消息类型，Text, Image, Audio, File
		Text string `json:"text"` // 文本消息内容
		// 文件消息链
		File []struct {
			MimeType string `json:"mime_type"` // 文件类型(MIME type, 如：image/jpeg)
			Path     string `json:"path"`      // 文件本地路径
			URL      string `json:"url"`       // 文件网络路径（并不是所有文件都有）
			Name     string `json:"name"`      // 文件名
			Size     int    `json:"size"`      // 文件大小
		}
	} `json:"message_chain"`
	// 消息发送者信息
	Sender struct {
		ID       string `json:"id"`       // 消息发送者ID
		Username string `json:"username"` // 消息发送者用户名
		// 群聊信息
		Group struct {
			ID      string `json:"id"`       // 群聊ID
			Title   string `json:"title"`    // 群聊名称
			IsAdmin string `json:"is_admin"` // 是否为管理员
		}
	}
}

func MessageReceiving(c *gin.Context) error {
	// 解析得到的Json
	var JsonBody WebHookStruct.WebHook
	if err := c.ShouldBindJSON(&JsonBody); err != nil {
		return err
	}

	// 判断消息来源
	if JsonBody.Telegram.UpdateID != 0 {
		// 消息为Telegram消息
		fmt.Println(JsonBody)
	}
	if JsonBody.Mirai.Type != "" {
		// 消息为Mirai消息
		fmt.Println(JsonBody)
	}
	if JsonBody.Line.Destination != "" {
		// 消息为Line消息
		fmt.Println(JsonBody)
	}

	return nil
}
