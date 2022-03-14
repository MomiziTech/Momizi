/*
 * @Author: McPlus
 * @Date: 2022-03-14 18:26:05
 * @LastEditTime: 2022-03-14 18:46:03
 * @LastEdit: McPlus
 * @Description: Chat方法
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Methods\Chat.go
 */
package TelegramMethods

import (
	"github.com/MomiziTech/Momizi/Controller/MessageSend/ChatSoftwareAPI/Telegram"
)

func GetAdministrators(ChatID int) ([]Telegram.ChatMemberAdministrator, error) {
	return Telegram.NewChat(ChatID).GetAdministrators()
}
