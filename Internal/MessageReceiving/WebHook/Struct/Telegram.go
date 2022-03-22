/*
 * @Author: NyanCatda
 * @Date: 2022-03-09 14:20:49
 * @LastEditTime: 2022-03-09 15:21:03
 * @LastEditors: NyanCatda
 * @Description: Telegram WebHook结构体
 * @FilePath: \Momizi\Internal\MessageReceiving\WebHookStruct\Telegram.go
 */
package Struct

type Telegram struct {
	UpdateID int `json:"update_id"` // 更新ID
	Message  struct {
		MessageID int `json:"message_id"` // 消息ID
		From      struct {
			ID           int    `json:"id"`            // 发送者ID
			IsBot        bool   `json:"is_bot"`        // 发送者是否是机器人
			FirstName    string `json:"first_name"`    // 发送者名
			LastName     string `json:"last_name"`     // 发送者姓
			Username     string `json:"username"`      // 发送者用户名
			LanguageCode string `json:"language_code"` // 发送者使用的语言(zh-hans)
		} `json:"from"`
		Chat struct {
			ID                          int    `json:"id"`                             // 聊天ID
			FirstName                   string `json:"first_name"`                     // 聊天名
			LastName                    string `json:"last_name"`                      // 聊天姓
			Username                    string `json:"username"`                       // 聊天用户名
			Title                       string `json:"title"`                          // 聊天标题
			Type                        string `json:"type"`                           // 聊天类型，可能类型："private", "group", "supergroup"，"channel"
			AllMembersAreAdministrators bool   `json:"all_members_are_administrators"` // 是否所有成员都是管理员
		} `json:"chat"`
		Date int `json:"date"` // 消息发送时间（时间戳）

		// 文字消息
		Text string `json:"text"` // 消息内容

		// 图片消息
		Photo []struct {
			FileID       string `json:"file_id"`        // 图片ID
			FileUniqueID string `json:"file_unique_id"` // 图片唯一ID
			FileSize     int    `json:"file_size"`      // 图片大小(Byte)
			Width        int    `json:"width"`          // 图片宽度(Px)
			Height       int    `json:"height"`         // 图片高度(Px)
		} `json:"photo"` // 图片消息组

		// 文件消息
		Document struct {
			FileName string `json:"file_name"` // 文件名
			MimeType string `json:"mime_type"` // 文件类型(MIME type, 如：image/jpeg)
			Thumb    struct {
				FileID       string `json:"file_id"`        // 缩略图ID
				FileUniqueID string `json:"file_unique_id"` // 缩略图唯一ID
				FileSize     int    `json:"file_size"`      // 缩略图大小(Byte)
				Width        int    `json:"width"`          // 缩略图宽度(Px)
				Height       int    `json:"height"`         // 缩略图高度(Px)
			} `json:"thumb"`
			FileID       string `json:"file_id"`        // 文件ID
			FileUniqueID string `json:"file_unique_id"` // 文件唯一ID
			FileSize     int    `json:"file_size"`      // 文件大小(Byte)
		} `json:"document"` // 文件消息

		// 语音消息
		Voice struct {
			Duration     int    `json:"duration"`       // 语音时长(Seconds)
			MimeType     string `json:"mime_type"`      // 语音类型(MIME type, 如：audio/ogg)
			FileID       string `json:"file_id"`        // 语音ID
			FileUniqueID string `json:"file_unique_id"` // 语音唯一ID
			FileSize     int    `json:"file_size"`      // 语音大小(Byte)
		} `json:"voice"` // 语音消息

		// 被回复消息
		ReplyToMessage struct {
			MessageID int `json:"message_id"` // 消息ID
			From      struct {
				ID           int    `json:"id"`            // 发送者ID
				IsBot        bool   `json:"is_bot"`        // 发送者是否是机器人
				FirstName    string `json:"first_name"`    // 发送者名
				LastName     string `json:"last_name"`     // 发送者姓
				Username     string `json:"username"`      // 发送者用户名
				LanguageCode string `json:"language_code"` // 发送者使用的语言(zh-hans)
			} `json:"from"`
			Chat struct {
				ID                          int    `json:"id"`                             // 聊天ID
				FirstName                   string `json:"first_name"`                     // 聊天名
				LastName                    string `json:"last_name"`                      // 聊天姓
				Username                    string `json:"username"`                       // 聊天用户名
				Title                       string `json:"title"`                          // 聊天标题
				Type                        string `json:"type"`                           // 聊天类型，可能类型："private", "group", "supergroup"，"channel"
				AllMembersAreAdministrators bool   `json:"all_members_are_administrators"` // 是否所有成员都是管理员
			} `json:"chat"`
			Date int `json:"date"` // 消息发送时间（时间戳）

			// 文字消息
			Text string `json:"text"` // 消息内容

			// 图片消息
			Photo []struct {
				FileID       string `json:"file_id"`        // 图片ID
				FileUniqueID string `json:"file_unique_id"` // 图片唯一ID
				FileSize     int    `json:"file_size"`      // 图片大小(Byte)
				Width        int    `json:"width"`          // 图片宽度(Px)
				Height       int    `json:"height"`         // 图片高度(Px)
			} `json:"photo"` // 图片消息组

			// 文件消息
			Document struct {
				FileName string `json:"file_name"` // 文件名
				MimeType string `json:"mime_type"` // 文件类型(MIME type, 如：image/jpeg)
				Thumb    struct {
					FileID       string `json:"file_id"`        // 缩略图ID
					FileUniqueID string `json:"file_unique_id"` // 缩略图唯一ID
					FileSize     int    `json:"file_size"`      // 缩略图大小(Byte)
					Width        int    `json:"width"`          // 缩略图宽度(Px)
					Height       int    `json:"height"`         // 缩略图高度(Px)
				} `json:"thumb"`
				FileID       string `json:"file_id"`        // 文件ID
				FileUniqueID string `json:"file_unique_id"` // 文件唯一ID
				FileSize     int    `json:"file_size"`      // 文件大小(Byte)
			} `json:"document"` // 文件消息

			// 语音消息
			Voice struct {
				Duration     int    `json:"duration"`       // 语音时长(Seconds)
				MimeType     string `json:"mime_type"`      // 语音类型(MIME type, 如：audio/ogg)
				FileID       string `json:"file_id"`        // 语音ID
				FileUniqueID string `json:"file_unique_id"` // 语音唯一ID
				FileSize     int    `json:"file_size"`      // 语音大小(Byte)
			} `json:"voice"` // 语音消息
		} `json:"reply_to_message"`
	} `json:"message"` // 消息
}
