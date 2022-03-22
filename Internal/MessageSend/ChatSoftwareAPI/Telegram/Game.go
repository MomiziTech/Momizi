/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:24:58
 * @LastEditTime: 2022-03-14 17:49:39
 * @LastEdit: McPlus
 * @Description: Game结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Game.go
 */
package Telegram

type Game struct {
	Title        string          `json:"title"`         // Title of the game
	Description  string          `json:"description"`   // Description of the game
	Photo        []PhotoSize     `json:"photo"`         // Photo that will be displayed in the game message in chats.
	Text         string          `json:"text"`          // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []MessageEntity `json:"text_entities"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    Animation       `json:"animation"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}
