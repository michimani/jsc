package types

type IResponse interface{}

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
	User        *string             `json:"user,omitempty"`
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

// GetUsersInfoResponse is struct representing response of `GET https://slack.com/api/users.info`.
// https://api.slack.com/methods/users.info
type GetUsersInfoResponse struct {
	OK *bool `json:"ok,omitempty"`

	// If 'OK' is true, following fields will included.
	// https://api.slack.com/methods/users.info#examples
	User UserInfo `json:"user,omitempty"`

	// If 'OK' is false, Error fields will included.
	// https://api.slack.com/methods/users.info#errors
	Error *string `json:"error,omitempty"`
}

type UserInfo struct {
	ID                *string      `json:"id,omitempty"`
	TeamID            *string      `json:"team_id,omitempty"`
	Name              *string      `json:"name,omitempty"`
	Deleted           *bool        `json:"deleted,omitempty"`
	Color             *string      `json:"color,omitempty"`
	RealName          *string      `json:"real_name,omitempty"`
	Tz                *string      `json:"tz,omitempty"`
	TzLabel           *string      `json:"tz_label,omitempty"`
	TzOffset          *int64       `json:"tz_offset,omitempty"`
	Profile           *UserProfile `json:"profile,omitempty"`
	IsAdmin           *bool        `json:"is_admin,omitempty"`
	IsOwner           *bool        `json:"is_owner,omitempty"`
	IsPrimaryOwner    *bool        `json:"is_primary_owner,omitempty"`
	IsRestricted      *bool        `json:"is_restricted,omitempty"`
	IsUltraRestricted *bool        `json:"is_ultra_restricted,omitempty"`
	IsBot             *bool        `json:"is_bot,omitempty"`
	Updated           *TsInt       `json:"updated,omitempty"`
	IsAppUser         *bool        `json:"is_app_user,omitempty"`
	Has2fa            *bool        `json:"has_2fa,omitempty"`
}

type UserProfile struct {
	AvatarHash            *string `json:"avatar_hash,omitempty"`
	StatusText            *string `json:"status_text,omitempty"`
	StatusEmoji           *string `json:"status_emoji,omitempty"`
	RealName              *string `json:"real_name,omitempty"`
	DisplayName           *string `json:"display_name,omitempty"`
	RealNameNormalized    *string `json:"real_name_normalized,omitempty"`
	DisplayNameNormalized *string `json:"display_name_normalized,omitempty"`
	Email                 *string `json:"email,omitempty"`
	ImageOriginal         *string `json:"image_original,omitempty"`
	Image24               *string `json:"image_24,omitempty"`
	Image32               *string `json:"image_32,omitempty"`
	Image48               *string `json:"image_48,omitempty"`
	Image72               *string `json:"image_72,omitempty"`
	Image192              *string `json:"image_192,omitempty"`
	Image512              *string `json:"image_512,omitempty"`
	Team                  *string `json:"team,omitempty"`
}
