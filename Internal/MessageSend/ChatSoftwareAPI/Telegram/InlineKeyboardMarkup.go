/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:54:19
 * @LastEditTime: 2022-03-14 17:54:20
 * @LastEdit: McPlus
 * @Description:InlineKeyboardMarkup
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\InlineKeyboardMarkup.go
 */
package Telegram

type InlineKeyboardMarkup struct {
	InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}
