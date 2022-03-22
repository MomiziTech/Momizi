/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:15:24
 * @LastEditTime: 2022-03-14 17:15:25
 * @LastEdit: McPlus
 * @Description: MessageEntity结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\MessageEntity.go
 */
package Telegram

type MessageEntity struct {
	Type     string `json:"type"`     // Type of the entity. Currently, can be "mention" (@username), "hashtag" (#hashtag), "cashtag" ($USD), "bot_command" (/start@jobs_bot), "url" (https://telegram.org), "email" (do-not-reply@telegram.org), "phone_number" (+1-212-555-0123), "bold" (bold text), "italic" (italic text), "underline" (underlined text), "strikethrough" (strikethrough text), "spoiler" (spoiler message), "code" (monowidth string), "pre" (monowidth block), "text_link" (for clickable text URLs), "text_mention" (for users without usernames)
	Offset   int    `json:"offset"`   // Offset in UTF-16 code units to the start of the entity
	Length   int    `json:"length"`   // Length of the entity in UTF-16 code units
	Url      string `json:"url"`      // Optional. For "text_link" only, url that will be opened after user taps on the text
	User     User   `json:"user"`     // Optional. For "text_mention" only, the mentioned user
	Language string `json:"language"` // Optional. For "pre" only, the programming language of the entity text
}
