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
