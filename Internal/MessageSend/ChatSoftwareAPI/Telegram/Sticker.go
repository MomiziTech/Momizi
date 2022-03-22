/*
 * @Author: McPlus
 * @Date: 2022-03-14 17:19:24
 * @LastEditTime: 2022-03-14 17:47:35
 * @LastEdit: McPlus
 * @Description: Sticker结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Sticker.go
 */
package Telegram

type Sticker struct {
	FileID       string       `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string       `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int          `json:"width"`          // Sticker width
	Height       int          `json:"height"`         // Sticker height
	IsAnimated   bool         `json:"is_animated"`    // True, if the sticker is animated
	IsVideo      bool         `json:"is_video"`       // True, if the sticker is a video sticker
	Thumb        PhotoSize    `json:"thumb"`          // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji        string       `json:"emoji"`          // Optional. Emoji associated with the sticker
	SetName      string       `json:"set_name"`       // Optional. Name of the sticker set to which the sticker belongs
	MaskPosition MaskPosition `json:"mask_position"`  // Optional. For mask stickers, the position where the mask should be placed
	FileSize     int          `json:"file_size"`      // Optional. File size in bytes
}
