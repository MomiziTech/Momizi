/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:10:03
 * @LastEditTime: 2022-03-14 17:10:04
 * @LastEdit: McPlus
 * @Description:
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\ChatLocation.go
 */
package Telegram

type ChatLocation struct {
	Location Location `json:"location"` // The location to which the supergroup is connected. Can't be a live location.
	Address  string   `json:"address"`  // Location address; 1-64 characters, as defined by the chat owner
}
