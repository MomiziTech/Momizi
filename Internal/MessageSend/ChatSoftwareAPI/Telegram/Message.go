/*
 * @Author: McPlus
 * @Date: 2022-03-14 16:46:55
 * @LastEditTime: 2022-03-17 23:54:17
 * @LastEdit: McPlus
 * @Description: Message结构体
 * @FilePath: \Momizi\Internal\MessageSend\ChatSoftwareAPI\Telegram\Message.go
 */
package Telegram

type Message struct {
	MessageID                     int                           `json:"message_id"`                        // Unique message identifier inside this chat
	From                          User                          `json:"from"`                              // Optional. Sender of the message; empty for messages sent to channels. For backward compatibility, the field contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	SenderChat                    Chat                          `json:"sender_chat"`                       // Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for channel posts, the supergroup itself for messages from anonymous group administrators, the linked channel for messages automatically forwarded to the discussion group. For backward compatibility, the field from contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
	Date                          int                           `json:"date"`                              // Date the message was sent in Unix time
	Chat                          Chat                          `json:"chat"`                              // Conversation the message belongs to
	ForwardFrom                   User                          `json:"forward_from"`                      // Optional. For forwarded messages, sender of the original message
	ForwardFromChat               Chat                          `json:"forward_from_chat"`                 // Optional. For messages forwarded from channels or from anonymous administrators, information about the original sender chat
	ForwardFromMessageID          int                           `json:"forward_from_message_id"`           // Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature              string                        `json:"forward_signature"`                 // Optional. For forwarded messages that were originally sent in channels or by an anonymous chat administrator, signature of the message sender if present
	ForwardSenderName             string                        `json:"forward_sender_name"`               // Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardDate                   int                           `json:"forward_date"`                      // Optional. For forwarded messages, date the original message was sent in Unix time
	IsAutomaticForward            bool                          `json:"is_automatic_forward"`              // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	ViaBot                        User                          `json:"via_bot"`                           // Optional. Bot through which the message was sent
	EditDate                      int                           `json:"edit_date"`                         // Optional. Date the message was last edited in Unix time
	HasProtectedContent           bool                          `json:"has_protected_content"`             // Optional. True, if the message can't be forwarded
	MediaGroupID                  string                        `json:"media_group_id"`                    // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               string                        `json:"author_signature"`                  // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          string                        `json:"text"`                              // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Entities                      []MessageEntity               `json:"entities"`                          // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Animation                     Animation                     `json:"animation"`                         // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         Audio                         `json:"audio"`                             // Optional. Message is an audio file, information about the file
	Document                      Document                      `json:"document"`                          // Optional. Message is a general file, information about the file
	Photo                         []PhotoSize                   `json:"photo"`                             // Optional. Message is a photo, available sizes of the photo
	Sticker                       Sticker                       `json:"sticker"`                           // Optional. Message is a sticker, information about the sticker
	Video                         Video                         `json:"video"`                             // Optional. Message is a video, information about the video
	VideoNote                     VideoNote                     `json:"video_note"`                        // Optional. Message is a video note, information about the video message
	Voice                         Voice                         `json:"voice"`                             // Optional. Message is a voice message, information about the file
	Caption                       string                        `json:"caption"`                           // Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	CaptionEntities               []MessageEntity               `json:"caption_entities"`                  // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Contact                       Contact                       `json:"contact"`                           // Optional. Message is a shared contact, information about the contact
	Dice                          Dice                          `json:"dice"`                              // Optional. Message is a dice with random value
	Game                          Game                          `json:"game"`                              // Optional. Message is a game, information about the game. More about games »
	Poll                          Poll                          `json:"poll"`                              // Optional. Message is a native poll, information about the poll
	Venue                         Venue                         `json:"venue"`                             // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      Location                      `json:"location"`                          // Optional. Message is a shared location, information about the location
	NewChatMembers                []User                        `json:"new_chat_members"`                  // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                User                          `json:"left_chat_member"`                  // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  string                        `json:"new_chat_title"`                    // Optional. A chat title was changed to this value
	NewChatPhoto                  []PhotoSize                   `json:"new_chat_photo"`                    // Optional. A chat photo was change to this value
	DeleteChatPhoto               bool                          `json:"delete_chat_photo"`                 // Optional. Service message: the chat photo was deleted
	GroupChatCreated              bool                          `json:"group_chat_created"`                // Optional. Service message: the group has been created
	SupergroupChatCreated         bool                          `json:"supergroup_chat_created"`           // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            bool                          `json:"channel_chat_created"`              // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"` // Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatID               int                           `json:"migrate_to_chat_id"`                // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID             int                           `json:"migrate_from_chat_id"`              // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Invoice                       Invoice                       `json:"invoice"`                           // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment             SuccessfulPayment             `json:"successful_payment"`                // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	ConnectedWebsite              string                        `json:"connected_website"`                 // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	PassportData                  PassportData                  `json:"passport_data"`                     // Optional. Telegram Passport data
	ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered"`         // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	VoiceChatScheduled            VoiceChatScheduled            `json:"voice_chat_scheduled"`              // Optional. Service message: voice chat scheduled
	VoiceChatStarted              VoiceChatStarted              `json:"voice_chat_started"`                // Optional. Service message: voice chat started
	VoiceChatEnded                VoiceChatEnded                `json:"voice_chat_ended"`                  // Optional. Service message: voice chat ended
	VoiceChatParticipantsInvited  VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited"`   // Optional. Service message: new participants invited to a voice chat
	ReplyMarkup                   InlineKeyboardMarkup          `json:"reply_markup"`                      // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	PinnedMessage                 PinnedMessage                 `json:"pinned_message"`                    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	// ReplyToMessage                Message                       `json:"reply_to_message"`                  // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
}

type PinnedMessage struct {
	Id            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Username      string `json:"username"`
	Type          string `json:"type"`
	PinnedMessage struct {
		MessageId int  `json:"message_id"`
		From      User `json:"from"`
		Chat      struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"pinned_message"`
}
