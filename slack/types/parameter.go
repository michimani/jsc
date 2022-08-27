package types

type IParameter interface {
}

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

func (p *ConversationsHistoryParameter) Value() ConversationsHistoryParameter {
	if p == nil {
		return ConversationsHistoryParameter{}
	}
	return *p
}
