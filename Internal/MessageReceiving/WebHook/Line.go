/*
 * @Author: NyanCatda
 * @Date: 2022-03-28 13:57:28
 * @LastEditTime: 2022-03-28 14:58:11
 * @LastEditors: NyanCatda
 * @Description: Line消息处理
 * @FilePath: \Momizi\Internal\MessageReceiving\WebHook\Line.go
 */
package WebHook

import (
	"path/filepath"
	"strconv"

	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/WebHook/Struct"
	LineMethods "github.com/MomiziTech/Momizi/Internal/MessageSend/ChatSoftwareAPI/Line"
	"github.com/MomiziTech/Momizi/Tools/File"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
)

func Line(WebHookJson Struct.WebHook) (MessageStruct.MessageStruct, error) {
	ChatSoftwareName := "Line"

	var MessageType string // 消息类型
	var MessageID string   // 消息ID
	var Timestamp int64    // 消息接收时间
	var Sender MessageStruct.MessageSender
	// 根据消息类型组成消息链
	var MessageChain []MessageStruct.MessageChain

	// 读取配置文件
	ConfigLine := ReadConfig.GetConfig.ChatSoftware.Line

	// 遍历消息
	for _, Event := range WebHookJson.Events {
		// 获取消息类型
		switch Event.Source.Type {
		case "user":
			MessageType = "User"
		case "group":
			MessageType = "Group"
		case "room":
			MessageType = "Group"
		default:
			MessageType = "Other"
		}
		// 获取消息ID
		MessageID = Event.ReplyToken
		// 获取消息接收时间
		Timestamp = Event.Timestamp

		var GroupInfo MessageStruct.MessageSenderGroup
		// 如果是群聊消息则获取群聊信息
		if MessageType == "Group" {
			GroupInfo = MessageStruct.MessageSenderGroup{
				ID: Event.Source.GroupID,
				// ToDo: 获取群聊名称
				Title:   Event.Source.GroupID,
				IsAdmin: false,
			}
		}

		// 获取发送者信息
		Sender = MessageStruct.MessageSender{
			ID: Event.Source.UserID,
			// ToDo: 获取用户名
			Username: Event.Source.UserID,
			Group:    GroupInfo,
		}

		// 如果为文字消息
		if Event.Message.Type == "text" {
			Text := Event.Message.Text
			TextMessage := MessageStruct.MessageChain{
				Type: "Text",
				Text: Text,
			}
			MessageChain = append(MessageChain, TextMessage)
		}

		// 如果为图片消息
		if Event.Message.Type == "image" {
			// 获取文件ID
			FileID := Event.Message.ID
			// 下载文件
			FilePath, FileSize, err := LineMethods.GetContent(FileID)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 获取文件名
			FileName := filepath.Base(FilePath)

			// 创建图片消息
			ImageMessage := MessageStruct.MessageChain{
				Type: "Image",
				File: MessageStruct.MessageChainFile{
					MimeType: "image/jpeg",
					Path:     FilePath,
					URL:      ConfigLine.DataAPILink + "v2/bot/message/" + FileID + "/content",
					Name:     FileName,
					Size:     FileSize,
				},
			}
			MessageChain = append(MessageChain, ImageMessage)
		}

		// 如果为文件消息
		if Event.Message.Type == "file" {
			// 获取文件ID
			FileID := Event.Message.ID
			// 下载文件
			FilePath, _, err := LineMethods.GetContent(FileID)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 根据文件名字重命名文件
			FileDirPath, _ := filepath.Split(FilePath)
			NewFilePath := filepath.Join(FileDirPath, Event.Message.FileName)
			err = File.Move(FilePath, NewFilePath)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 创建文件消息
			FileMessage := MessageStruct.MessageChain{
				Type: "File",
				File: MessageStruct.MessageChainFile{
					MimeType: "application/octet-stream",
					Path:     NewFilePath,
					URL:      ConfigLine.DataAPILink + "v2/bot/message/" + FileID + "/content",
					Name:     Event.Message.FileName,
					Size:     Event.Message.FileSize,
				},
			}
			MessageChain = append(MessageChain, FileMessage)
		}

		// 如果为语音消息
		if Event.Message.Type == "audio" {
			// 获取文件ID
			FileID := Event.Message.ID
			// 下载文件
			FilePath, FileSize, err := LineMethods.GetContent(FileID)
			if err != nil {
				return MessageStruct.MessageStruct{}, err
			}

			// 获取文件名
			FileName := filepath.Base(FilePath)

			// 创建语音消息
			AudioMessage := MessageStruct.MessageChain{
				Type: "Audio",
				File: MessageStruct.MessageChainFile{
					MimeType: "audio/*",
					Path:     FilePath,
					URL:      ConfigLine.DataAPILink + "v2/bot/message/" + FileID + "/content",
					Name:     FileName,
					Size:     FileSize,
				},
			}
			MessageChain = append(MessageChain, AudioMessage)
		}
	}

	// 时间戳转换为时间到秒
	Timestamp = Timestamp / 1000
	strInt64 := strconv.FormatInt(Timestamp, 10)
	Time, err := strconv.Atoi(strInt64)
	if err != nil {
		return MessageStruct.MessageStruct{}, err
	}

	// 创建消息结构体
	Message := MessageStruct.MessageStruct{
		ID:           MessageID,
		Type:         MessageType,
		ChatSoftware: ChatSoftwareName,
		Time:         Time,
		MessageChain: MessageChain,
		Sender:       Sender,
	}
	return Message, nil
}
