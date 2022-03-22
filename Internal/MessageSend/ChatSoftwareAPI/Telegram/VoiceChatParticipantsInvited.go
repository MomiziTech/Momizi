/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:53:27
 * @LastEditTime: 2022-03-14 17:53:29
 * @LastEdit: McPlus
 * @Description: VoiceChatParticipantsInvited
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\VoiceChatParticipantsInvited.go
 */
package Telegram

type VoiceChatParticipantsInvited struct {
	Users []User `json:"users"` // Optional. New members that were invited to the voice chat
}
