/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:50:40
 * @LastEditTime: 2022-03-14 17:50:42
 * @LastEdit: McPlus
 * @Description: VoiceChatScheduled结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\VoiceChatScheduled.go
 */
package Telegram

type VoiceChatScheduled struct {
	StartDate int `json:"start_date"` // Point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator
}
