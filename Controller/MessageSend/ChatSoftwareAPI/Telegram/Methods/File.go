/*
 * @Author: McPlus
 * @Date: 2022-03-14 18:23:22
 * @LastEditTime: 2022-03-14 18:50:05
 * @LastEdit: McPlus
 * @Description: File方法
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Methods\File.go
 */
package TelegramMethods

import (
	"github.com/MomiziTech/Momizi/Controller/MessageSend/ChatSoftwareAPI/Telegram"
)

/**
 * @description: 通过文件ID获取文件结构体
 * @param {string} FileID
 * @return {FileStruct, error}
 */
func GetFile(FileID string) (Telegram.File, error) {
	return Telegram.NewFile(FileID).GetFile()
}
