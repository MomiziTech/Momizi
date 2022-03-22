/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:24:22
 * @LastEditTime: 2022-03-14 17:24:24
 * @LastEdit: McPlus
 * @Description: Dice结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Dice.go
 */
package Telegram

type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int    `json:"value"` // Value of the dice, 1-6 for "🎲", "🎯" and "🎳" base emoji, 1-5 for "🏀" and "⚽" base emoji, 1-64 for "🎰" base emoji
}
