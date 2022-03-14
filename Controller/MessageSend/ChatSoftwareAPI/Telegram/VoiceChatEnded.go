/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:52:14
 * @LastEditTime: 2022-03-14 17:52:15
 * @LastEdit: McPlus
 * @Description: VoiceChatEnded
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\VoiceChatEnded.go
 */
package Telegram

type VoiceChatEnded struct {
	Duration int `json:"duration"` // Voice chat duration in seconds
}