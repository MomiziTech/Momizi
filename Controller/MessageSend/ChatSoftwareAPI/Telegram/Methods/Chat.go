/*
 * @Author: McPlus
 * @Date: 2022-03-14 18:26:05
 * @LastEditTime: 2022-03-14 18:32:53
 * @LastEdit: McPlus
 * @Description: Chat方法
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Methods\Chat.go
 */
package TelegramMethods

import (
	"encoding/json"
	"strconv"

	"github.com/MomiziTech/Momizi/Controller/MessageSend/ChatSoftwareAPI/Telegram"
	"github.com/MomiziTech/Momizi/Utils/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

func GetAdministrators(ChatID int) ([]Telegram.ChatMemberAdministrator, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAdress := ConfigTelegram.BotAPILink + "bot" + ConfigTelegram.APIToken + "/getChatAdministrators"

	DataMap := map[string]string{
		"chat_id": strconv.Itoa(ChatID),
	}

	Buffer, Response, Error := HttpRequest.PostRequestXWWWForm(APIAdress, []string{}, DataMap)
	var JsonData []Telegram.ChatMemberAdministrator
	if Response.StatusCode == 200 {
		json.Unmarshal(Buffer, &JsonData)
		return JsonData, Error
	} else {
		return []Telegram.ChatMemberAdministrator{}, Error
	}
}
