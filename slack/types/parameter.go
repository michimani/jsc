package types

type IParameter interface {
}

// ConversationsHistoryParameter is struct for parameter of conversations.history api.
// https://api.slack.com/methods/conversations.history
type ConversationsHistoryParameter struct {
	// Required
	// Conversation ID to fetch history for.
	Channel *string `paramkey:"channel"`

	// Optional
	Cursor             *string `paramkey:"cursor"`
	IncludeAllMetadata *bool   `paramkey:"include_all_metadata"`
	Inclusive          *bool   `paramkey:"inclusive"`
	Latest             *string `paramkey:"latest"`
	Limit              *int    `paramkey:"limit"`
	Oldest             *string `paramkey:"oldest"`
}

// GetUsersInfoParameter is struct for parameter of users.info api.
// https://api.slack.com/methods/users.info
type GetUsersInfoParameter struct {
	// Required
	// User to get info on.
	User *string `paramkey:"user"`

	// Optional
	IncludeLocale *bool `paramkey:"include_locale"`
}

// GetChatPermalinkParameter is struct for parameter of chat.getPermalink api.
// https://api.slack.com/methods/chat.getPermalink
type GetChatPermalinkParameter struct {
	// Required
	// The ID of the conversation or channel containing the message
	Channel *string `paramkey:"channel"`
	// A message's ts value, uniquely identifying it within a channel
	MessageTs *string `paramkey:"message_ts"`
}
