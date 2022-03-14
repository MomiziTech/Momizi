/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:16:57
 * @LastEditTime: 2022-03-14 17:16:58
 * @LastEdit: McPlus
 * @Description: PhotoSize结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\PhotoSize.go
 */
package Telegram

type PhotoSize struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int    `json:"width"`          // Photo width
	Height       int    `json:"height"`         // Photo height
	FileSize     int    `json:"file_size"`      // Optional. File size in bytes
}
