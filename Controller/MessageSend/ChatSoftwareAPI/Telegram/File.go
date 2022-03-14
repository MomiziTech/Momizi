/*
 * @Author: McPlus
 * @Date: 2022-03-09 13:11:32
 * @LastEditTime: 2022-03-14 19:02:13
 * @LastEdit: McPlus
 * @Description:
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\File.go
 */
package Telegram

import (
	"encoding/json"

	"github.com/MomiziTech/Momizi/Utils/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type File struct {
	ID       string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	Size     int    `json:"file_size"`
	Path     string `json:"file_path"`
}

func NewFile(ID string) *File {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAdress := ConfigTelegram.BotAPILink + "bot" + ConfigTelegram.APIToken + "/getFile"

	DataMap := map[string]string{
		"file_id": ID,
	}

	Buffer, Response, _ := HttpRequest.PostRequestXWWWForm(APIAdress, []string{}, DataMap)
	var JsonData *File
	if Response.StatusCode == 200 {
		json.Unmarshal(Buffer, &JsonData)
		return JsonData
	} else {
		return &File{}
	}
}
