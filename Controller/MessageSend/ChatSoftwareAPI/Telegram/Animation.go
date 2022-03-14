/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:16:14
 * @LastEditTime: 2022-03-14 17:16:16
 * @LastEdit: McPlus
 * @Description: Animation结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Animation.go
 */
package Telegram

type Animation struct {
	FileID       string    `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string    `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int       `json:"width"`          // Video width as defined by sender
	Height       int       `json:"height"`         // Video height as defined by sender
	Duration     int       `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        PhotoSize `json:"thumb"`          // Optional. Animation thumbnail as defined by sender
	FileName     string    `json:"file_name"`      // Optional. Original animation filename as defined by sender
	MimeType     string    `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int       `json:"file_size"`      // Optional. File size in bytes
}
