/*
 * @Author: McPlus
 * @Date: 2022-03-09 13:11:32
 * @LastEditTime: 2022-03-14 18:31:53
 * @LastEdit: McPlus
 * @Description:
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\File.go
 */
package Telegram

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     string `json:"file_path"`
}
