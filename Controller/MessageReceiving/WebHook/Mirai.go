/*
 * @Author: NyanCatda
 * @Date: 2022-03-12 22:42:49
 * @LastEditTime: 2022-03-13 12:33:09
 * @LastEditors: NyanCatda
 * @Description: Mirai消息处理
 * @FilePath: \Momizi\Controller\MessageReceiving\WebHook\Mirai.go
 */
package WebHook

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Controller/MessageReceiving/WebHook/Struct"
	"github.com/MomiziTech/Momizi/Utils"
)

var (
	SaveFilePath = "/data/file/Mirai/"
)

func Mirai(WebHookJson Struct.WebHook) (MessageStruct.MessageStruct, error) {
	ChatSoftwareName := "QQ"

	var Type string
	switch WebHookJson.Mirai.Type {
	case "FriendMessage":
		Type = "User"
	case "GroupMessage":
		Type = "Group"
	case "TempMessage":
		Type = "User"
	default:
		Type = "Other"
	}

	var Time int
	var MessageID int

	var Text string

	var MessageChain []MessageStruct.MessageChain

	// 遍历消息链
	for _, Message := range WebHookJson.Mirai.MessageChain {
		// 解析Source消息信息
		if Message.Type == "Source" {
			Time = Message.Time
			MessageID = Message.ID
		}

		// 解析At消息
		if Message.Type == "At" || Message.Type == "AtAll" {
			Text += "@" + Message.Display
		}

		// 解析文字消息信息
		if Message.Type == "Plain" {
			Text += Message.Text
		}

		// 解析图片消息
		if Message.Type == "Image" {
			// 将文件下载至本地
			timeUnix := time.Now().Unix()
			FilePath, FileSize, err := Utils.DownloadFile(Message.URL, SaveFilePath+strconv.FormatInt(timeUnix, 10)+"/", 120)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			_, fileName := filepath.Split(FilePath)

			FileInfo := MessageStruct.MessageChainFile{
				MimeType: "image/jpeg",
				Path:     FilePath,
				URL:      Message.URL,
				Name:     fileName,
				Size:     FileSize,
			}
			ImageMessage := MessageStruct.MessageChain{
				Type: "Image",
				File: FileInfo,
			}
			MessageChain = append(MessageChain, ImageMessage)
		}

		// 解析语音消息
		if Message.Type == "Voice" {
			timeUnix := time.Now().Unix()
			FilePath, FileSize, err := Utils.DownloadFile(Message.URL, SaveFilePath+strconv.FormatInt(timeUnix, 10)+"/", 120)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			_, fileName := filepath.Split(FilePath)

			FileInfo := MessageStruct.MessageChainFile{
				MimeType: "audio/*",
				Path:     FilePath,
				URL:      Message.URL,
				Name:     fileName,
				Size:     FileSize,
			}
			VoiceMessage := MessageStruct.MessageChain{
				Type: "Audio",
				File: FileInfo,
			}
			MessageChain = append(MessageChain, VoiceMessage)
		}
	}

	// 如果文本消息不为空则添加进入消息链
	if Text != "" {
		TextMessage := MessageStruct.MessageChain{
			Type: "Text",
			Text: Text,
		}
		MessageChain = append(MessageChain, TextMessage)
	}

	// 判断是否需要初始化群聊信息
	var GroupInfo MessageStruct.MessageSenderGroup
	if Type == "Group" {
		GroupInfo = MessageStruct.MessageSenderGroup{
			ID:      strconv.Itoa(WebHookJson.Mirai.Sender.Group.ID),
			Title:   WebHookJson.Mirai.Sender.Group.Name,
			IsAdmin: WebHookJson.Mirai.Sender.Permission == "ADMINISTRATOR",
		}
	}

	Message := MessageStruct.MessageStruct{
		ID:           strconv.Itoa(MessageID),
		Type:         Type,
		ChatSoftware: ChatSoftwareName,
		Time:         Time,
		MessageChain: MessageChain,
		Sender: MessageStruct.MessageSender{
			ID:       strconv.Itoa(WebHookJson.Mirai.Sender.ID),
			Username: WebHookJson.Mirai.Sender.Nickname,
			Group:    GroupInfo,
		},
	}

	return Message, nil
}
