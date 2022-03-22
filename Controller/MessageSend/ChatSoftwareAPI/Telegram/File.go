/*
 * @Author: McPlus
 * @Date: 2022-03-09 13:11:32
 * @LastEditTime: 2022-03-19 18:08:15
 * @LastEdit: McPlus
 * @Description:
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\File.go
 */
package Telegram

import (
	"encoding/json"
	"errors"

	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type File struct {
	ID       string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	Size     int64  `json:"file_size"`
	Path     string `json:"file_path"`
}

type GetFileReturn struct {
	BasicReturn

	File *File `json:"result"`
}

func NewFile(ID string) (*File, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAdress := ConfigTelegram.BotAPILink + "bot" + ConfigTelegram.APIToken + "/getFile"

	DataMap := map[string]string{
		"file_id": ID,
	}

	Buffer, _, Error := HttpRequest.PostRequestXWWWForm(APIAdress, []string{}, DataMap)
	var JsonData GetFileReturn
	json.Unmarshal(Buffer, &JsonData)

	if Error != nil {
		return &File{}, Error
	}

	if JsonData.Success {
		return JsonData.File, nil
	}

	return &File{}, errors.New(string(Buffer))
}
