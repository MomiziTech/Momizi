package Telegram

import (
	"encoding/json"

	"github.com/MomiziTech/Momizi/Utils/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type FileStruct struct {
	FileID string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize int `json:"file_size"`
	FilePath string `json:"file_path"`
}

func GetFile(FileID string) (FileStruct, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAdress := ConfigTelegram.BotAPILink + "bot" + ConfigTelegram.APIToken + "/getFile"

	DataMap := make(map[string]string)
	DataMap["file_id"] = FileID

	Buffer, Response, Error := HttpRequest.PostRequestXWWWForm(APIAdress, []string{}, DataMap)
	var JsonData FileStruct
	if Response.StatusCode == 200 {
		json.Unmarshal(Buffer, &JsonData)
		return JsonData, Error
	} else {
		return FileStruct{
			FileID: FileID,
		}, Error
	}
}