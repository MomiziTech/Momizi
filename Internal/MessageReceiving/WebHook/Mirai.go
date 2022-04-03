/*
 * @Author: NyanCatda
 * @Date: 2022-03-12 22:42:49
 * @LastEditTime: 2022-04-03 22:28:54
 * @LastEditors: NyanCatda
 * @Description: Mirai消息处理
 * @FilePath: \Momizi\Internal\MessageReceiving\WebHook\Mirai.go
 */
package WebHook

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/MomiziTech/Momizi/Internal/Controller"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/WebHook/Struct"
	MiraiAPI "github.com/MomiziTech/Momizi/Internal/MessageSend/ChatSoftwareAPI/Mirai"
	"github.com/MomiziTech/Momizi/Tools"
	"github.com/MomiziTech/Momizi/Tools/File"
)

func Mirai(WebHookJson Struct.WebHook) (MessageStruct.MessageStruct, error) {
	ChatSoftwareName := "QQ"
	SaveFilePath := Controller.BotFilePath + ChatSoftwareName + "/"

	// 判断聊天类型
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
		return MessageStruct.MessageStruct{}, nil
	}

	// 获取聊天ID
	var ChatID int
	if Type == "Group" {
		ChatID = WebHookJson.Mirai.Sender.Group.ID
	} else {
		ChatID = WebHookJson.Mirai.Sender.ID
	}

	var Time int
	var MessageID any

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
			FilePath, FileSize, err := Tools.DownloadFile(Message.URL, []string{}, SaveFilePath+strconv.FormatInt(timeUnix, 10)+"/", true, 120)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 纠正文件类型
			FilePath, err = File.CorrectFileType(FilePath)
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
			FilePath, FileSize, err := Tools.DownloadFile(Message.URL, []string{}, SaveFilePath+strconv.FormatInt(timeUnix, 10)+"/", true, 120)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 纠正文件类型
			FilePath, err = File.CorrectFileType(FilePath)
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

		// 解析文件消息
		if Message.Type == "File" {
			// 获取文件URL
			FileInfo, err := MiraiAPI.GetFileInfo(Message.ID.(string), strconv.Itoa(ChatID))
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}
			URL := FileInfo.Data.DownloadInfo.URL
			// 将文件下载至本地
			Path, FileSize, err := Tools.DownloadFile(URL, []string{}, SaveFilePath+strconv.FormatInt(time.Now().Unix(), 10)+"/", true, 120)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 修正文件名字
			CorrectionPath := filepath.Dir(Path) + "/" + FileInfo.Data.Name
			err = File.Move(Path, CorrectionPath)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 组成消息链
			FileMessage := MessageStruct.MessageChainFile{
				MimeType: "application/octet-stream",
				Path:     CorrectionPath,
				URL:      URL,
				Name:     FileInfo.Data.Name,
				Size:     FileSize,
			}
			DocumentMessage := MessageStruct.MessageChain{
				Type: "File",
				File: FileMessage,
			}
			MessageChain = append(MessageChain, DocumentMessage)
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

	// 获取消息发送者用户名
	var Username string
	if Type == "Group" {
		Username = WebHookJson.Mirai.Sender.MemberName
	} else {
		Username = WebHookJson.Mirai.Sender.Nickname
	}

	Message := MessageStruct.MessageStruct{
		ID:           fmt.Sprintln(MessageID),
		ChatID:       strconv.Itoa(ChatID),
		Type:         Type,
		ChatSoftware: ChatSoftwareName,
		Time:         Time,
		MessageChain: MessageChain,
		Sender: MessageStruct.MessageSender{
			ID:       strconv.Itoa(WebHookJson.Mirai.Sender.ID),
			Username: Username,
			Group:    GroupInfo,
		},
	}

	return Message, nil
}
