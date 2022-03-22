/*
 * @Author: McPlus
 * @Date: 2022-03-19 17:22:42
 * @LastEditTime: 2022-03-19 18:10:07
 * @LastEdit: McPlus
 * @Description: 基础返回结构
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\BasicReturn.go
 */
package Telegram

type BasicReturn struct {
	Success     bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}
