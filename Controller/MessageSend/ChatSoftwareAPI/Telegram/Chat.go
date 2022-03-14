/*
 * @Author: McPlus
 * @Date: 2022-03-14 16:11:54
 * @LastEditTime: 2022-03-14 18:45:59
 * @LastEdit: McPlus
 * @Description: Chat功能
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Chat.go
 */
package Telegram

import (
	"encoding/json"
	"strconv"

	"github.com/MomiziTech/Momizi/Utils/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type Chat struct {
	ID                    int             `json:"id"`                       // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                  string          `json:"type"`                     // Type of chat, can be either "private", "group", "supergroup" or "channel"
	Title                 string          `json:"title"`                    // 	Optional. Title, for supergroups, channels and group chats
	UserName              string          `json:"username"`                 // Optional. Username, for private chats, supergroups and channels if available
	FirstName             string          `json:"first_name"`               // Optional. First name of the other party in a private chat
	LastName              string          `json:"last_name"`                // 	Optional. Last name of the other party in a private chat
	Photo                 ChatPhoto       `json:"photo"`                    // Optional. Chat photo. Returned only in getChat.
	Bio                   string          `json:"bio"`                      // Optional. Bio of the other party in a private chat. Returned only in getChat.
	HasPrivateForwards    bool            `json:"has_private_forwards"`     // Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	Description           string          `json:"description"`              // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink            string          `json:"invite_link"`              // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	Permissions           ChatPermissions `json:"permissions"`              // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay         int             `json:"slow_mode_delay"`          // 	Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
	MessageAutoDeleteTime int             `json:"message_auto_delete_time"` // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	HasProtectedContent   bool            `json:"has_protected_content"`    // Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
	StickerSetName        string          `json:"sticker_set_name"`         // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet      bool            `json:"can_set_sticker_set"`      // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	LinkedChatID          int             `json:"linked_chat_id"`           // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
	Location              ChatLocation    `json:"location"`                 // 	Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
	// PinnedMessage         Message   `json:"pinned_message"`           // Optional. The most recent pinned message (by sending date). Returned only in getChat.
}

func NewChat(ID int) *Chat {
	return &Chat{ID: ID}
}

func (Chat Chat) GetAdministrators() ([]ChatMemberAdministrator, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAdress := ConfigTelegram.BotAPILink + "bot" + ConfigTelegram.APIToken + "/getChatAdministrators"

	DataMap := map[string]string{
		"chat_id": strconv.Itoa(Chat.ID),
	}

	Buffer, Response, Error := HttpRequest.PostRequestXWWWForm(APIAdress, []string{}, DataMap)
	var JsonData []ChatMemberAdministrator
	if Response.StatusCode == 200 {
		json.Unmarshal(Buffer, &JsonData)
		return JsonData, Error
	} else {
		return []ChatMemberAdministrator{}, Error
	}
}
