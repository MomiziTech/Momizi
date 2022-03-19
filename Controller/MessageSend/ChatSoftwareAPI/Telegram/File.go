/*
 * @Author: McPlus
 * @Date: 2022-03-09 13:11:32
 * @LastEditTime: 2022-03-19 17:38:50
 * @LastEdit: McPlus
 * @Description:
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\File.go
 */
package Telegram

import (
	"encoding/json"
	"net/http"

	"github.com/MomiziTech/Momizi/Utils/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type File struct {
	ID       string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	Size     int    `json:"file_size"`
	Path     string `json:"file_path"`
}

type GetFileReturn struct {
	BasicReturn

	Result *File `json:"result"`
}

func NewFile(ID string) (GetFileReturn, *http.Response, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAdress := ConfigTelegram.BotAPILink + "bot" + ConfigTelegram.APIToken + "/getFile"

	DataMap := map[string]string{
		"file_id": ID,
	}

	Buffer, Response, Error := HttpRequest.PostRequestXWWWForm(APIAdress, []string{}, DataMap)
	var JsonData GetFileReturn
	json.Unmarshal(Buffer, &JsonData)
	return JsonData, Response, Error
}
