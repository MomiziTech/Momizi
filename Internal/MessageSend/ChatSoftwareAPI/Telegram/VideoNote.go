/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:22:28
 * @LastEditTime: 2022-03-14 17:22:29
 * @LastEdit: McPlus
 * @Description: VideoNote结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\VideoNote.go
 */
package Telegram

type VideoNote struct {
	FileID       string    `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string    `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int       `json:"length"`         // Video width and height (diameter of the video message) as defined by sender
	Duration     int       `json:"duration"`       // Duration of the video in seconds as defined by sender
	Thumb        PhotoSize `json:"thumb"`          // Optional. Video thumbnail
	FileSize     int       `json:"file_size"`      // Optional. File size in bytes
}
