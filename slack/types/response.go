package types

import "strconv"

type IResponse interface{}

// TsString is type of timestamp string.
type TsString string

func (ts *TsString) String() string {
	if ts == nil {
		return ""
	}

	return string(*ts)
}

func (ts *TsString) Float64() (float64, error) {
	if ts == nil {
		return 0, nil
	}

	f, err := strconv.ParseFloat(ts.String(), 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

// ConversationsHistoryResponse is struct representing response of `GET conversations.history`.
// https://api.slack.com/methods/conversations.history
type ConversationsHistoryResponse struct {
	OK *bool `json:"ok,omitempty"`

	// If 'OK' is true, following fields will included.
	// https://api.slack.com/methods/conversations.history#examples
	Messages            []ConversationsHistoryMessage `json:"messages,omitempty"`
	HasMore             *bool                         `json:"has_more,omitempty"`
	PinCount            *int64                        `json:"pin_count,omitempty"`
	ChannelActionsTs    *TsString                     `json:"channel_actions_ts,omitempty"`
	ChannelActionsCount *int64                        `json:"channel_actions_count,omitempty"`
	ResponseMetadata    *ResponseMetadata             `json:"response_metadata,omitempty"`

	// If 'OK' is false, Error fields will included.
	// https://api.slack.com/methods/conversations.history#errors
	Error *string `json:"error,omitempty"`
}

type ResponseMetadata struct {
	NextCursor *string
}

type ConversationsHistoryMessage struct {
	Type        *string             `json:"type,omitempty"`
	Subtype     *string             `json:"subtype,omitempty"`
	Text        *string             `json:"text,omitempty"`
	Ts          *TsString           `json:"ts,omitempty"`
	Username    *string             `json:"username,omitempty"`
	Icons       *MessageIcon        `json:"icons,omitempty"`
	BotID       *string             `json:"bot_id,omitempty"`
	Attachments []MessageAttachment `json:"attachments,omitempty"`
}

type MessageIcon struct {
	Emoji *string `json:"emoji,omitempty"`
}

type MessageAttachment struct {
	ImageURL    *string           `json:"image_url,omitempty"`
	ImageWidth  *int64            `json:"image_width,omitempty"`
	ImageHeight *int64            `json:"image_height,omitempty"`
	ImageBytes  *int64            `json:"image_bytes,omitempty"`
	ID          *int              `json:"id,omitempty"`
	Color       *string           `json:"color,omitempty"`
	Fallback    *string           `json:"fallback,omitempty"`
	Title       *string           `json:"title,omitempty"`
	AuthorName  *string           `json:"author_name,omitempty"`
	Footer      *string           `json:"footer,omitempty"`
	Fields      []AttachmentField `json:"fields,omitempty"`
}

type AttachmentField struct {
	Value *string `json:"value,omitempty"`
	Title *string `json:"title,omitempty"`
	Short *bool   `json:"short,omitempty"`
}
