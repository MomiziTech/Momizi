/*
 * @Author: McPlus
 * @Date: 2022-03-14 18:23:22
 * @LastEditTime: 2022-03-14 18:25:43
 * @LastEdit: McPlus
 * @Description: File方法
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Methods\File.go
 */
package TelegramMethods

import (
	"encoding/json"

	"github.com/MomiziTech/Momizi/Controller/MessageSend/ChatSoftwareAPI/Telegram"
	"github.com/MomiziTech/Momizi/Utils/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

/**
 * @description: 通过文件ID获取文件结构体
 * @param {string} FileID
 * @return {FileStruct, error}
 */
func GetFile(FileID string) (Telegram.File, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAdress := ConfigTelegram.BotAPILink + "bot" + ConfigTelegram.APIToken + "/getFile"

	DataMap := make(map[string]string)
	DataMap["file_id"] = FileID

	Buffer, Response, Error := HttpRequest.PostRequestXWWWForm(APIAdress, []string{}, DataMap)
	var JsonData Telegram.File
	if Response.StatusCode == 200 {
		json.Unmarshal(Buffer, &JsonData)
		return JsonData, Error
	} else {
		return Telegram.File {
			FileID: FileID,
		}, Error
	}
}
