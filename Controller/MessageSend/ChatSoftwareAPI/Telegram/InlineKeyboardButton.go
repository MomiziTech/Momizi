/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:54:54
 * @LastEditTime: 2022-03-14 17:56:28
 * @LastEdit: McPlus
 * @Description: InlineKeyboardButton
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\InlineKeyboardButton.go
 */
package Telegram

type InlineKeyboardButton struct {
	Text                         string       `json:"text"`                             // Label text on the button
	Url                          string       `json:"url"`                              // Optional. HTTP or tg:// url to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their ID without using a username, if this is allowed by their privacy settings.
	LoginUrl                     LoginUrl     `json:"login_url"`                        // Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	CallbackData                 string       `json:"callback_data"`                    // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery            string       `json:"switch_inline_query"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. Can be empty, in which case just the bot's username will be inserted.
	SwitchInlineQueryCurrentChat string       `json:"switch_inline_query_current_chat"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot's username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
	CallbackGame                 CallbackGame `json:"callback_game"`                    // Optional. Description of the game that will be launched when the user presses the button.
	Pay                          bool         `json:"pay"`                              // Optional. Specify True, to send a Pay button.
}
