/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:49:56
 * @LastEditTime: 2022-03-14 17:49:58
 * @LastEdit: McPlus
 * @Description: ProximityAlertTriggered结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\ProximityAlertTriggered.go
 */
package Telegram

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"` // User that triggered the alert
	Watcher  User `json:"watcher"`  // User that set the alert
	Distance int  `json:"distance"` // The distance between the users
}
