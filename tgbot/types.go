package tgbot

// User object represents a Telegram user or bot.
type User struct {
	// Unique identifier for this user or bot
	Id 						int 	`json:"id"`
	// True, if this user is a bot
	IsBot 					bool 	`json:"is_bot"`
	// 	User‘s or bot’s first name
	FirstName 				string 	`json:"first_name"`
	// Optional. User‘s or bot’s last name
	LastName 				string 	`json:"last_name,omitempty"`
	// Optional. User‘s or bot’s username
	Username 				string 	`json:"username,omitempty"`
	// Optional. IETF language tag of the user's language
	LanguageCode 			string 	`json:"language_code,omitempty"`
	// Optional. True, if the bot can be invited to groups.
	CanJoinGroups 			bool 	`json:"can_join_groups,omitempty"`
	// Optional. True, if privacy mode is disabled for the bot.
	CanReadAllGroupMessages bool 	`json:"can_read_all_group_messages,omitempty"`
	// Optional. True, if the bot supports inline queries.
	SupportsInlineQueries 	bool 	`json:"supports_inline_queries,omitempty"`
}

type Message struct {
	MessageId int `json:"message_id"`
	From *User `json:"from,omitempty"`
	Date int `json:"date"`
	Chat *Chat `json:"chat"`
	ForwardFrom *User `json:"forward_from,omitempty"`
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId int `json:"forward_from_message_id,omitempty"`
	ForwardSignature string `json:"forward_signature,omitempty"`
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	ForwardDate int `json:"forward_date,omitempty"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	EditDate int `json:"edit_date,omitempty"`
	MediaGroupId string `json:"media_group_id,omitempty"`
	AuthorSignature string `json:"author_signature,omitempty"`
	Text string `json:"text,omitempty"`
	Entities *[]MessageEntity `json:"entities,omitempty"`
	CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`

}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// Optional. True, if the user is allowed to send text messages, contacts, locations and venues.
	CanSendMessages 		bool `json:"can_send_messages,omitempty"`
	// Optional. rue, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes,
	// implies CanSendMessages.
	CanSendMediaMessages 	bool `json:"can_send_media_messages,omitempty"`
	// Optional. True, if the user is allowed to send polls, implies CanSendMessages.
	CanSendPolls 			bool `json:"can_send_polls,omitempty"`
	// Optional. True, if the user is allowed to send animations, games, stickers and use inline bots,
	//implies CanSendMediaMessages.
	CanSendOtherMessages 	bool `json:"can_send_other_messages,omitempty"`
	// Optional. True, if the user is allowed to add web page previews to their messages, implies CanSendMediaMessages.
	CanAddWebPagePreviews 	bool `json:"can_add_web_page_previews,omitempty"`
	// Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups.
	CanChangeInfo 			bool `json:"can_change_info,omitempty"`
	// Optional. True, if the user is allowed to invite new users to the chat.
	CanInviteUsers 			bool `json:"can_invite_users,omitempty"`
	// Optional. True, if the user is allowed to pin messages. Ignored in public supergroups.
	CanPinMessages 			bool `json:"can_pin_messages,omitempty"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo.
	// This FileID can be used only for photo download and only for as long as the photo is not changed.
	SmallFileId string `json:"small_file_id"`
	// Unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	//Can't be used to download or reuse the file.
	SmallFileUniqueId string `json:"small_file_unique_id"`
	// File identifier of big (640x640) chat photo.
	// This FileId can be used only for photo download and only for as long as the photo is not changed.
	BigFileId string `json:"big_file_id"`
	// Unique file identifier of big (640x640) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueId string `json:"big_file_unique_id"`
}

// Chat represents a chat.
type Chat struct {
	// Unique identifier for this chat.
	// This number may be greater than 32 bits and some programming languages may have difficulty/silent defects
	// in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision
	// float type are safe for storing this identifier.
	Id 					int64 				`json:"id"`
	// 	Type of chat, can be either “private”, “group”, “supergroup” or “channel”.
	Type 				string 				`json:"type"`
	// Optional. Title, for supergroups, channels and group chats
	Title 				string 				`json:"title,omitempty"`
	// Optional. Username, for private chats, supergroups and channels if available
	Username 			string 				`json:"username,omitempty"`
	// Optional. First name of the other party in a private chat
	FirstName 			string 				`json:"first_name,omitempty"`
	// Optional. Last name of the other party in a private chat
	LastName 			string 				`json:"last_name,omitempty"`
	// Optional. Chat photo. Returned only in GetChat.
	Photo 				*ChatPhoto 			`json:"photo,omitempty"`
	// Optional. Description, for groups, supergroups and channel chats. Returned only in GetChat.
	Description 		string 				`json:"description,omitempty"`
	// Optional. Chat invite link, for groups, supergroups and channel chats. Each administrator in a chat
	// generates their own invite links, so the bot must first generate the link using ExportChatInviteLink.
	// Returned only in GetChat.
	InviteLink 			string 				`json:"invite_link,omitempty"`
	// Optional. Pinned message, for groups, supergroups and channels. Returned only in GetChat.
	PinnedMessage 		*Message 			`json:"pinned_message,omitempty"`
	// Optional. Default chat member permissions, for groups and supergroups. Returned only in GetChat.
	Permissions 		*ChatPermissions 	`json:"permissions,omitempty"`
	// Optional. For supergroups, the minimum allowed delay between consecutive messages
	// sent by each unpriviledged user. Returned only in GetChat.
	SlowModeDelay 		int 				`json:"slow_mode_delay,omitempty"`
	// Optional. For supergroups, name of group sticker set. Returned only in GetChat.
	StickerSetName 		string 				`json:"sticker_set_name,omitempty"`
	// Optional. True, if the bot can change the group sticker set. Returned only in GetChat.
	CanSetStickerSet 	bool 				`json:"can_set_sticker_set,omitempty"`
}

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// 	Type of the entity. Can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD),
	//	“bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org),
	//	“phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text),
	//	“strikethrough” (strikethrough text), “code” (monowidth string), “pre” (monowidth block),
	//	“text_link” (for clickable text URLs), “text_mention” (for users without usernames)
	Type 		string 	`json:"type"`
	// 	Offset in UTF-16 code units to the start of the entity.
	Offset		int 	`json:"offset"`
	// Length of the entity in UTF-16 code units.
	Length 		int 	`json:"length"`
	// Optional. For “text_link” only, url that will be opened after user taps on the text.
	Url 		string 	`json:"url,omitempty"`
	// Optional. For “text_mention” only, the mentioned user.
	User 		*User 	`json:"user,omitempty"`
	// Optional. For “pre” only, the programming language of the entity text.
	Language 	string `json:"language,omitempty"`
}
