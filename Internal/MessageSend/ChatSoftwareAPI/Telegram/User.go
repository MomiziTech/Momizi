/*
 * @Author: McPlus
 * @Date: 2022-03-14 16:25:27
 * @LastEditTime: 2022-03-14 16:34:11
 * @LastEdit: McPlus
 * @Description:
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\User.go
 */
package Telegram

type User struct {
	ID                      int    `json:"id"`                          // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	IsBot                   bool   `json:"is_bot"`                      // True, if this user is a bot
	FirstName               string `json:"first_name"`                  // User's or bot's first name
	LastName                string `json:"last_name"`                   // User's or bot's first name
	UserName                string `json:"user_name"`                   // Optional. User's or bot's username
	LanguageCOde            string `json:"language_code"`               // Optional. IETF language tag of the user's language
	CanJoinGroups           bool   `json:"can_join_groups"`             // Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     // Optional. True, if the bot supports inline queries. Returned only in getMe.
}
