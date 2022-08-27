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
