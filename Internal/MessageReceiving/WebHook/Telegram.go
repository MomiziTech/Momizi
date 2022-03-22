/*
 * @Author: NyanCatda
 * @Date: 2022-03-19 15:57:39
 * @LastEditTime: 2022-03-21 00:38:13
 * @LastEditors: NyanCatda
 * @Description: Telegram消息处理
 * @FilePath: \Momizi\Internal\MessageReceiving\WebHook\Telegram.go
 */
package WebHook

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/MessageStruct"
	"github.com/MomiziTech/Momizi/Internal/MessageReceiving/WebHook/Struct"
	TelegramMethods "github.com/MomiziTech/Momizi/Internal/MessageSend/ChatSoftwareAPI/Telegram/Methods"
	"github.com/MomiziTech/Momizi/Tools"
	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
)

func Telegram(WebHookJson Struct.WebHook) (MessageStruct.MessageStruct, error) {
	ChatSoftwareName := "Telegram"
	SaveFilePath := "data/file/Telegram/"

	// 获取消息类型
	var MessageType string
	switch WebHookJson.Message.Chat.Type {
	case "private":
		MessageType = "User"
	case "group":
		MessageType = "Group"
	case "supergroup":
		MessageType = "Group"
	default:
		MessageType = "Other"
		return MessageStruct.MessageStruct{}, nil
	}

	// 根据消息类型组成消息链
	var MessageChain []MessageStruct.MessageChain
	// 如果为文字消息
	if WebHookJson.Message.Text != "" {
		Text := WebHookJson.Message.Text
		TextMessage := MessageStruct.MessageChain{
			Type: "Text",
			Text: Text,
		}
		MessageChain = append(MessageChain, TextMessage)
	}

	// 如果为图片消息
	if WebHookJson.Message.Photo != nil {
		// 获取最清晰的图片
		PhotoInfo := WebHookJson.Message.Photo[len(WebHookJson.Message.Photo)-1]
		// 获取图片链接
		PhotoID := PhotoInfo.FileID
		PhotoFileInfo, err := TelegramMethods.GetFile(PhotoID)
		if err != nil {
			return MessageStruct.MessageStruct{}, err
		}
		PhotoURL := GetTelegramFileURL(PhotoFileInfo.Path)
		// 下载图片
		timeUnix := time.Now().Unix()
		FilePath, _, err := Tools.DownloadFile(PhotoURL, SaveFilePath+strconv.FormatInt(timeUnix, 10)+"/", false, 120)
		if err != nil {
			return MessageStruct.MessageStruct{}, err
		}

		_, FileName := filepath.Split(FilePath)

		// 组成图片消息链
		FileMessage := MessageStruct.MessageChainFile{
			MimeType: "image/jpeg",
			Path:     FilePath,
			URL:      PhotoURL,
			Name:     FileName,
			Size:     PhotoFileInfo.Size,
		}
		PhotoMessage := MessageStruct.MessageChain{
			Type: "Image",
			File: FileMessage,
		}
		MessageChain = append(MessageChain, PhotoMessage)
	}

	// 如果为文件消息
	if WebHookJson.Message.Document.FileID != "" {
		// 获取文件链接
		DocumentID := WebHookJson.Message.Document.FileID
		DocumentFileInfo, err := TelegramMethods.GetFile(DocumentID)
		if err != nil {
			return MessageStruct.MessageStruct{}, err
		}
		DocumentURL := GetTelegramFileURL(DocumentFileInfo.Path)
		// 下载文件
		timeUnix := time.Now().Unix()
		FilePath, _, err := Tools.DownloadFile(DocumentURL, SaveFilePath+strconv.FormatInt(timeUnix, 10)+"/", false, 120)
		if err != nil {
			return MessageStruct.MessageStruct{}, err
		}

		// 将文件名字重命名
		OriginalFilePath, _ := filepath.Split(FilePath)
		if err := os.Rename(FilePath, OriginalFilePath+WebHookJson.Message.Document.FileName); err != nil {
			return MessageStruct.MessageStruct{}, err
		}
		FilePath = OriginalFilePath + WebHookJson.Message.Document.FileName

		// 组成文件消息链
		FileMessage := MessageStruct.MessageChainFile{
			MimeType: WebHookJson.Message.Document.MimeType,
			Path:     FilePath,
			URL:      DocumentURL,
			Name:     WebHookJson.Message.Document.FileName,
			Size:     DocumentFileInfo.Size,
		}
		DocumentMessage := MessageStruct.MessageChain{
			Type: "File",
			File: FileMessage,
		}
		MessageChain = append(MessageChain, DocumentMessage)
	}

	// 如果为语音消息
	if WebHookJson.Message.Voice.FileID != "" {
		// 获取语音链接
		VoiceID := WebHookJson.Message.Voice.FileID
		VoiceFileInfo, err := TelegramMethods.GetFile(VoiceID)
		if err != nil {
			return MessageStruct.MessageStruct{}, err
		}
		VoiceURL := GetTelegramFileURL(VoiceFileInfo.Path)
		// 下载语音
		timeUnix := time.Now().Unix()
		FilePath, _, err := Tools.DownloadFile(VoiceURL, SaveFilePath+strconv.FormatInt(timeUnix, 10)+"/", false, 120)
		if err != nil {
			return MessageStruct.MessageStruct{}, err
		}

		_, FileName := filepath.Split(FilePath)

		// 组成语音消息链
		FileMessage := MessageStruct.MessageChainFile{
			MimeType: WebHookJson.Message.Voice.MimeType,
			Path:     FilePath,
			URL:      VoiceURL,
			Name:     FileName,
			Size:     VoiceFileInfo.Size,
		}
		VoiceMessage := MessageStruct.MessageChain{
			Type: "Audio",
			File: FileMessage,
		}
		MessageChain = append(MessageChain, VoiceMessage)
	}

	// 判断是否是群组消息
	var GroupInfo MessageStruct.MessageSenderGroup
	if MessageType == "Group" {
		// 判断发送者是否为管理员
		AdminList, err := TelegramMethods.GetAdministrators(WebHookJson.Message.Chat.ID)
		if err != nil {
			return MessageStruct.MessageStruct{}, err
		}
		var IsAdmin bool
		for _, AdminInfo := range AdminList {
			if AdminInfo.User.ID == WebHookJson.Message.From.ID {
				IsAdmin = true
				break
			}
		}

		GroupInfo = MessageStruct.MessageSenderGroup{
			ID:      strconv.Itoa(WebHookJson.Message.Chat.ID),
			Title:   WebHookJson.Message.Chat.Title,
			IsAdmin: IsAdmin,
		}
	}

	// 获取发送者信息
	SenderInfo := MessageStruct.MessageSender{
		ID:       strconv.Itoa(WebHookJson.Message.From.ID),
		Username: WebHookJson.Message.From.Username,
		Group:    GroupInfo,
	}

	Message := MessageStruct.MessageStruct{
		ID:           strconv.Itoa(WebHookJson.Telegram.Message.MessageID),
		Type:         MessageType,
		ChatSoftware: ChatSoftwareName,
		Time:         WebHookJson.Telegram.Message.Date,
		MessageChain: MessageChain,
		Sender:       SenderInfo,
	}

	return Message, nil
}

func GetTelegramFileURL(FilePath string) string {
	APIToken := ReadConfig.GetConfig.ChatSoftware.Telegram.APIToken
	APIURL := ReadConfig.GetConfig.ChatSoftware.Telegram.BotAPILink

	FileURL := APIURL + "file/bot" + APIToken + "/" + FilePath
	return FileURL
}
