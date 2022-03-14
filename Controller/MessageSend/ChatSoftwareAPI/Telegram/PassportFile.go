/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:35:57
 * @LastEditTime: 2022-03-14 17:35:58
 * @LastEdit: McPlus
 * @Description: PassportFile结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Passport\PassportFile.go
 */
package Telegram

type PassportFile struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size"`      // File size in bytes
	FileDate     int    `json:"file_date"`      // Unix time when the file was uploaded
}
