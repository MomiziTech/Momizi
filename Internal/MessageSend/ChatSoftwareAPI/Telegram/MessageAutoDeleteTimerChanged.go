/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:28:01
 * @LastEditTime: 2022-03-14 17:28:02
 * @LastEdit: McPlus
 * @Description: MessageAutoDeleteTimerChanged结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\MessageAutoDeleteTimerChanged.go
 */
package Telegram

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"` // New auto-delete time for messages in the chat; in seconds
}
