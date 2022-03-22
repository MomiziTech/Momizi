/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:18:41
 * @LastEditTime: 2022-03-14 17:18:43
 * @LastEdit: McPlus
 * @Description: Document结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Document.go
 */
package Telegram

type Document struct {
	FileID       string    `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string    `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumb        PhotoSize `json:"thumb"`          // Optional. Document thumbnail as defined by sender
	FileName     string    `json:"file_name"`      // Optional. Original filename as defined by sender
	MimeType     string    `json:"mime_type"`      // Optional. MIME type of the file as defined by sender
	FileSize     int       `json:"file_size"`      // Optional. File size in bytes
}
