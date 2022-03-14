/*
 * @Author: McPlus
 * @Date: 2022-03-14 16:44:18
 * @LastEditTime: 2022-03-14 16:46:27
 * @LastEdit: McPlus
 * @Description: ChatPhoto结构体
 * @FilePath: \Momizi\Controller\MessageSend\ChatSoftwareAPI\Telegram\ChatPhoto.go
 */
package Telegram

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`        // File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileUniqueID string `json:"small_file_unique_id"` // Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileID         string `json:"big_file_id"`          // File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileUniqueID   string `json:"big_file_unique_id"`   // Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}
