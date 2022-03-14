/*
 * @Author: McPlus
 * @Date: 2022-03-14 16:43:36
 * @LastEditTime: 2022-03-14 16:43:37
 * @LastEdit: McPlus
 * @Description: ChatMembet结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\ChatMember.go
 */
package Telegram

type ChatMemberAdministrator struct {
	Status              string `json:"status"`                 // The member's status in the chat, always "administrator"
	User                User   `json:"user"`                   // Information about the user
	CanBeEdited         bool   `json:"can_be_edited"`          // True, if the bot is allowed to edit administrator privileges of that user
	IsAnonymous         bool   `json:"is_anonymous"`           // True, if the user's presence in the chat is hidden
	CanManageChat       bool   `json:"can_manage_chat"`        // True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
	CanDeleteMessages   bool   `json:"can_delete_messages"`    // True, if the administrator can delete messages of other users
	CanManageVoiceChats bool   `json:"can_manage_voice_chats"` //True, if the administrator can manage voice chats
	CanRestrictMembers  bool   `json:"can_restrict_members"`   // True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers   bool   `json:"can_promote_members"`    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       bool   `json:"can_change_info"`        // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      bool   `json:"can_invite_users"`       // True, if the user is allowed to invite new users to the chat
	CanPostMessage      bool   `json:"can_post_messages"`      // Optional. True, if the administrator can post in the channel; channels only
	CanEditMessage      bool   `json:"can_edit_messages"`      // Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanPinMessages      bool   `json:"can_pin_messages"`       // Optional. True, if the user is allowed to pin messages; groups and supergroups only
	CustomTitle         string `json:"custom_title"`           // Optional. Custom title for this user
}