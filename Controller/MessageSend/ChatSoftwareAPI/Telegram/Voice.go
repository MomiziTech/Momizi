/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:23:06
 * @LastEditTime: 2022-03-14 17:23:07
 * @LastEdit: McPlus
 * @Description: Voice结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Voice.go
 */
package Telegram

type Voice struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int    `json:"duration"`       // Duration of the audio in seconds as defined by sender
	MimeType     string `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int    `json:"file_size"`      // Optional. File size in bytes
}
