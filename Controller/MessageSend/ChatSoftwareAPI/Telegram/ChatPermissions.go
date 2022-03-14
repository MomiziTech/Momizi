/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:08:58
 * @LastEditTime: 2022-03-14 17:08:59
 * @LastEdit: McPlus
 * @Description: ChatPermissions结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\ChatPermissions.go
 */

package Telegram

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`         // Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMessages  bool `json:"can_send_media_messages"`   // Optional. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendPolls          bool `json:"can_send_polls"`            // Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendOtherMessages  bool `json:"can_send_other_messages"`   // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // Optional. True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages
	CanChangeInfo         bool `json:"can_change_info"`           // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers        bool `json:"can_invite_users"`          // Optional. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages"`          // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
}
