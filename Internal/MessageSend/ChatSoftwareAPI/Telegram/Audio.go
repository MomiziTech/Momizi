/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:18:09
 * @LastEditTime: 2022-03-14 17:18:10
 * @LastEdit: McPlus
 * @Description: Audio结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Audio.go
 */
package Telegram

type Audio struct {
	FileID       string    `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string    `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int       `json:"duration"`       // Duration of the audio in seconds as defined by sender
	Performer    string    `json:"performer"`      // Optional. Performer of the audio as defined by sender or by audio tags
	Title        string    `json:"title"`          // Optional. Title of the audio as defined by sender or by audio tags
	FileName     string    `json:"file_name"`      // Optional. Original filename as defined by sender
	MimeType     string    `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int       `json:"file_size"`      // Optional. File size in bytes
	Thumb        PhotoSize `json:"thumb"`          // Optional. Thumbnail of the album cover to which the music file belongs
}
