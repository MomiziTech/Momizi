/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:56:48
 * @LastEditTime: 2022-03-14 17:56:50
 * @LastEdit: McPlus
 * @Description: LoginUrl结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\LoginUrl.go
 */
package Telegram

type LoginUrl struct {
	Url                string `json:"url"`                  // An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.
	ForwardText        string `json:"forward_text"`         // Optional. New text of the button in forwarded messages.
	BotUsername        string `json:"bot_username"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess bool   `json:"request_write_access"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}
