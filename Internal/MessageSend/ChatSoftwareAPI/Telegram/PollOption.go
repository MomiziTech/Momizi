/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:26:48
 * @LastEditTime: 2022-03-14 17:26:49
 * @LastEdit: McPlus
 * @Description: PollOption结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\PollOption.go
 */
package Telegram

type PollOption struct {
	Text       string `json:"text"`        // Option text, 1-100 characters
	VoterCount int    `json:"voter_count"` // Number of users that voted for this option
}
