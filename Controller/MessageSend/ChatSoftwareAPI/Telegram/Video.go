/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:22:03
 * @LastEditTime: 2022-03-14 17:22:04
 * @LastEdit: McPlus
 * @Description:Video结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\Video.go
 */
package Telegram

type Video struct {
	FileID       string    `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string    `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int       `json:"width"`          // Video width as defined by sender
	Height       int       `json:"height"`         // Video height as defined by sender
	Duration     int       `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        PhotoSize `json:"thumb"`          // Optional. Video thumbnail
	FileName     string    `json:"file_name"`      // Optional. Original filename as defined by sender
	MimeType     string    `json:"mime_type"`      // Optional. Mime type of a file as defined by sender
	FileSize     int       `json:"file_size"`      // Optional. File size in bytes
}
