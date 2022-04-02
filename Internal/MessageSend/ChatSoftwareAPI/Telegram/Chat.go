/*
 * @Author: McPlus
 * @Date: 2022-03-14 16:11:54
 * @LastEditTime: 2022-04-02 09:24:40
 * @LastEdit: McPlus
 * @Description: Chat功能
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Chat.go
 */
package Telegram

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/MomiziTech/Momizi/Tools/ReadConfig"
	"github.com/nyancatda/HttpRequest"
)

type Chat struct {
	ID                    string          `json:"id"`                       // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Type                  string          `json:"type"`                     // Type of chat, can be either "private", "group", "supergroup" or "channel"
	Title                 string          `json:"title"`                    // 	Optional. Title, for supergroups, channels and group chats
	UserName              string          `json:"username"`                 // Optional. Username, for private chats, supergroups and channels if available
	FirstName             string          `json:"first_name"`               // Optional. First name of the other party in a private chat
	LastName              string          `json:"last_name"`                // 	Optional. Last name of the other party in a private chat
	Photo                 ChatPhoto       `json:"photo"`                    // Optional. Chat photo. Returned only in getChat.
	Bio                   string          `json:"bio"`                      // Optional. Bio of the other party in a private chat. Returned only in getChat.
	HasPrivateForwards    bool            `json:"has_private_forwards"`     // Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	Description           string          `json:"description"`              // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink            string          `json:"invite_link"`              // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
	Permissions           ChatPermissions `json:"permissions"`              // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay         int             `json:"slow_mode_delay"`          // 	Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
	MessageAutoDeleteTime int             `json:"message_auto_delete_time"` // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
	HasProtectedContent   bool            `json:"has_protected_content"`    // Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
	StickerSetName        string          `json:"sticker_set_name"`         // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet      bool            `json:"can_set_sticker_set"`      // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
	LinkedChatID          int             `json:"linked_chat_id"`           // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
	Location              ChatLocation    `json:"location"`                 // 	Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
	PinnedMessage         PinnedMessage   `json:"pinned_message"`           // Optional. The most recent pinned message (by sending date). Returned only in getChat.
}

type GetChatReturn struct {
	BasicReturn

	Chat *Chat `json:"result"`
}

/**
 * @description: 构建Chat
 * @param {int} ID ChatID
 * @return {*Chat}
 */
func NewChat(ID int) *Chat {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAddress := ConfigTelegram.APILink + "bot" + ConfigTelegram.APIToken + "/getChat"

	DataMap := map[string]string{
		"chat_id": strconv.Itoa(ID),
	}

	Buffer, _, _ := HttpRequest.PostRequestXWWWForm(APIAddress, []string{}, DataMap)
	var JsonData GetChatReturn
	json.Unmarshal(Buffer, &JsonData)
	if JsonData.Success {
		return JsonData.Chat
	}
	return &Chat{}
}

type GetAdministratorsReturn struct {
	BasicReturn

	Administrators []ChatMemberAdministrator `json:"result"`
}

/**
 * @description: 获取管理员列表
 * @param {*}
 * @return {[]ChatMemberAdministrator, error}
 */
func (Chat Chat) GetAdministrators() ([]ChatMemberAdministrator, error) {
	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAddress := ConfigTelegram.APILink + "bot" + ConfigTelegram.APIToken + "/getChatAdministrators"

	DataMap := map[string]string{
		"chat_id": Chat.ID,
	}

	Buffer, _, Error := HttpRequest.PostRequestXWWWForm(APIAddress, []string{}, DataMap)
	var JsonData GetAdministratorsReturn
	json.Unmarshal(Buffer, &JsonData)

	if Error != nil {
		return []ChatMemberAdministrator{}, Error
	}

	if JsonData.Success {
		return JsonData.Administrators, nil
	}

	return []ChatMemberAdministrator{}, errors.New(string(Buffer))

}

type SendMessageReturn struct {
	BasicReturn

	Message Message `json:"result"`
}

/**
 * @description:
 * @param {string} Text
 * @param {string} ParseMode *可选 MarkdownV2/HTML/Markdown -- Mode for parsing entities in the message text
 * @param {[]MessageEntity} Entities *可选 -- A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
 * @param {bool} DisableWebPagePreview *可选 -- Disables link previews for links in this message
 * @param {bool} DisableNotification *可选 -- Sends the message silently. Users will receive a notification with no sound.
 * @param {bool} ProtectContent *可选 -- Protects the contents of the sent message from forwarding and saving
 * @param {int} ReplyToMessaggeID *可选 -- If the message is a reply, ID of the original message
 * @param {bool} AllowSendingWithoutReply *可选 -- Pass True, if the message should be sent even if the specified replied-to message is not found
 * @return {*}
 */
func (Chat Chat) SendMessage(Text string, ParseMode string, Entities []MessageEntity, DisableWebPagePreview bool,
	DisableNotification bool, ProtectContent bool, ReplyToMessageID int, AllowSendingWithoutReply bool) (Message, error) {
	DataMap := map[string]string{
		"text": Text,
	}

	if ParseMode != "" {
		DataMap["parse_mode"] = ParseMode
	}

	if Entities != nil {
		data, _ := json.Marshal(Entities)
		DataMap["entities"] = string(data)
	}

	if DisableWebPagePreview {
		DataMap["disable_web_page_preview"] = strconv.FormatBool(DisableWebPagePreview)
	}

	if DisableNotification {
		DataMap["disable_notification"] = strconv.FormatBool(DisableNotification)
	}

	if ProtectContent {
		DataMap["protect_content"] = strconv.FormatBool(ProtectContent)
	}

	if ReplyToMessageID != -1 {
		DataMap["reply_to_message_id"] = strconv.Itoa(ReplyToMessageID)
	}

	if AllowSendingWithoutReply {
		DataMap["allow_sending_without_reply"] = strconv.FormatBool(AllowSendingWithoutReply)
	}

	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAddress := ConfigTelegram.APILink + "bot" + ConfigTelegram.APIToken + "/sendMessage"

	Buffer, _, Error := HttpRequest.PostRequestXWWWForm(APIAddress, []string{}, DataMap)
	var JsonData SendMessageReturn
	json.Unmarshal(Buffer, &JsonData)

	if Error != nil {
		return Message{}, Error
	}

	if JsonData.Success {
		return JsonData.Message, nil
	}

	return Message{}, errors.New(string(Buffer))
}

/**
 * @description: 发送图片消息
 * @param {string} Photo 图片路径
 * @param {string} Caption *可选 -- Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
 * @param {string} ParseMode *可选 MarkdownV2/HTML/Markdown -- Mode for parsing entities in the photo caption
 * @param {[]MessageEntity} CaptionEntities *可选 -- A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
 * @param {bool} DisableNotification *可选 -- Sends the message silently. Users will receive a notification with no sound.
 * @param {bool} ProtectContent *可选 -- Protects the contents of the sent message from forwarding and saving
 * @param {int} ReplyToMessageID *可选 -- If the message is a reply, ID of the original message
 * @param {bool} AllowSendingWithoutReply *可选 -- Pass True, if the message should be sent even if the specified replied-to message is not found
 * @return {*}
 */
func (Chat Chat) SendPhoto(Photo string, Caption string, ParseMode string, CaptionEntities []MessageEntity, DisableNotification bool,
	ProtectContent bool, ReplyToMessageID int, AllowSendingWithoutReply bool) (Message, error) {
	DataMap := make(map[string]string)

	if Caption != "" {
		DataMap["caption"] = Caption
	}

	if ParseMode != "" {
		DataMap["parse_mode"] = ParseMode
	}

	if CaptionEntities != nil {
		data, _ := json.Marshal(CaptionEntities)
		DataMap["disable_web_page_preview"] = string(data)
	}

	if DisableNotification {
		DataMap["disable_notification"] = strconv.FormatBool(DisableNotification)
	}

	if ProtectContent {
		DataMap["protect_content"] = strconv.FormatBool(ProtectContent)
	}

	if ReplyToMessageID != -1 {
		DataMap["reply_to_message_id"] = strconv.Itoa(ReplyToMessageID)
	}

	if AllowSendingWithoutReply {
		DataMap["allow_sending_without_reply"] = strconv.FormatBool(AllowSendingWithoutReply)
	}

	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAddress := ConfigTelegram.APILink + "bot" + ConfigTelegram.APIToken + "/sendPhoto"

	Buffer, _, Error := HttpRequest.PostRequestFormDataFile(APIAddress, []string{}, DataMap, "photo", []string{Photo})
	var JsonData SendMessageReturn
	json.Unmarshal(Buffer, &JsonData)

	if Error != nil {
		return Message{}, Error
	}

	if JsonData.Success {
		return JsonData.Message, nil
	}

	return Message{}, errors.New(string(Buffer))
}

/**
 * @description: 发送音频消息
 * @param {string} Audio 音频文件路径
 * @param {string} Caption *可选 -- Audio caption, 0-1024 characters after entities parsing
 * @param {string} ParseMode *可选 MarkdownV2/HTML/Markdown -- Mode for parsing entities in the audio caption. See formatting options for more details.
 * @param {[]MessageEntity} CaptionEntities *可选 -- A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
 * @param {int} Duration *可选 -- Duration of the audio in seconds
 * @param {string} Performer *可选 -- Performer
 * @param {string} Title *可选 -- Track name
 * @param {bool} DisableNotification *可选 -- Sends the message silently. Users will receive a notification with no sound.
 * @param {bool} ProtectContent *可选 -- Protects the contents of the sent message from forwarding and saving
 * @param {int} ReplyToMessageID *可选 -- If the message is a reply, ID of the original message
 * @param {bool} AllowSendingWithoutReply *可选 -- Pass True, if the message should be sent even if the specified replied-to message is not found
 * @return {*}
 */
func (Chat Chat) SendAudio(Audio string, Caption string, ParseMode string, CaptionEntities []MessageEntity, Duration int,
	Performer string, Title string, DisableNotification bool, ProtectContent bool, ReplyToMessageID int,
	AllowSendingWithoutReply bool) (Message, error) {
	DataMap := make(map[string]string)

	if Caption != "" {
		DataMap["caption"] = Caption
	}

	if ParseMode != "" {
		DataMap["parse_mode"] = ParseMode
	}

	if CaptionEntities != nil {
		data, _ := json.Marshal(CaptionEntities)
		DataMap["disable_web_page_preview"] = string(data)
	}

	if Duration != -1 {
		DataMap["duration"] = strconv.Itoa(Duration)
	}

	if Performer != "" {
		DataMap["performer"] = Performer
	}

	if Title != "" {
		DataMap["title"] = Title
	}

	if DisableNotification {
		DataMap["disable_notification"] = strconv.FormatBool(DisableNotification)
	}

	if ProtectContent {
		DataMap["protect_content"] = strconv.FormatBool(ProtectContent)
	}

	if ReplyToMessageID != -1 {
		DataMap["reply_to_message_id"] = strconv.Itoa(ReplyToMessageID)
	}

	if AllowSendingWithoutReply {
		DataMap["allow_sending_without_reply"] = strconv.FormatBool(AllowSendingWithoutReply)
	}

	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAddress := ConfigTelegram.APILink + "bot" + ConfigTelegram.APIToken + "/sendAudio"

	Buffer, _, Error := HttpRequest.PostRequestFormDataFile(APIAddress, []string{}, DataMap, "audio", []string{Audio})
	var JsonData SendMessageReturn
	json.Unmarshal(Buffer, &JsonData)

	if Error != nil {
		return Message{}, Error
	}

	if JsonData.Success {
		return JsonData.Message, nil
	}

	return Message{}, errors.New(string(Buffer))
}

/**
 * @description: 发送文件消息
 * @param {string} Document 文件路径
 * @param {string} Caption *可选 -- Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
 * @param {string} ParseMode *可选 MarkdownV2/HTML/Markdown -- Mode for parsing entities in the document caption.
 * @param {[]MessageEntity} CaptionEntities *可选 -- A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
 * @param {bool} DisableContentTypeDetection *可选 -- Disables automatic server-side content type detection for files uploaded using multipart/form-data
 * @param {bool} DisableNotification *可选 -- Sends the message silently. Users will receive a notification with no sound.
 * @param {bool} ProtectContent *可选 -- Protects the contents of the sent message from forwarding and saving
 * @param {int} ReplyToMessageID *可选 -- If the message is a reply, ID of the original message
 * @param {bool} AllowSendingWithoutReply *可选 -- Pass True, if the message should be sent even if the specified replied-to message is not found
 * @return {*}
 */
func (Chat Chat) SendDocument(Document string, Caption string, ParseMode string, CaptionEntities []MessageEntity,
	DisableContentTypeDetection bool, DisableNotification bool, ProtectContent bool, ReplyToMessageID int,
	AllowSendingWithoutReply bool) (Message, error) {
	DataMap := make(map[string]string)

	if Caption != "" {
		DataMap["caption"] = Caption
	}

	if ParseMode != "" {
		DataMap["parse_mode"] = ParseMode
	}

	if CaptionEntities != nil {
		data, _ := json.Marshal(CaptionEntities)
		DataMap["disable_web_page_preview"] = string(data)
	}

	if DisableContentTypeDetection {
		DataMap["disable_content_type_detection"] = strconv.FormatBool(DisableContentTypeDetection)
	}

	if DisableNotification {
		DataMap["disable_notification"] = strconv.FormatBool(DisableNotification)
	}

	if ProtectContent {
		DataMap["protect_content"] = strconv.FormatBool(ProtectContent)
	}

	if ReplyToMessageID != -1 {
		DataMap["reply_to_message_id"] = strconv.Itoa(ReplyToMessageID)
	}

	if AllowSendingWithoutReply {
		DataMap["allow_sending_without_reply"] = strconv.FormatBool(AllowSendingWithoutReply)
	}

	Config := ReadConfig.GetConfig

	ConfigTelegram := Config.ChatSoftware.Telegram

	APIAddress := ConfigTelegram.APILink + "bot" + ConfigTelegram.APIToken + "/sendDocument"

	Buffer, _, Error := HttpRequest.PostRequestFormDataFile(APIAddress, []string{}, DataMap, "document", []string{Document})
	var JsonData SendMessageReturn
	json.Unmarshal(Buffer, &JsonData)

	if Error != nil {
		return Message{}, Error
	}

	if JsonData.Success {
		return JsonData.Message, nil
	}

	return Message{}, errors.New(string(Buffer))
}
