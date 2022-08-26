package types

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type ConversationsHistoryParameter struct {
	// Required
	// Conversation ID to fetch history for.
	Channel string

	// Optional
	Cursor             *string
	IncludeAllMetadata *bool
	Inclusive          *bool
	Latest             *string
	Limit              *int
	Oldest             *string

	token string
}

func (p *ConversationsHistoryParameter) SetToken(token string) {
	if p == nil {
		return
	}
	p.token = token
}

func (p *ConversationsHistoryParameter) Body() (*strings.Reader, error) {
	if p == nil {
		return nil, errors.New("ConversationsHistoryParameter is nil.")
	}

	uv := url.Values{}

	if p.Channel == "" {
		return nil, errors.New("Channel is required.")
	}

	uv.Add("channel", p.Channel)

	if p.Cursor != nil {
		uv.Add("cursor", *p.Cursor)
	}
	if p.IncludeAllMetadata != nil {
		uv.Add("include_all_metadata", strconv.FormatBool(*p.IncludeAllMetadata))
	}
	if p.Inclusive != nil {
		uv.Add("inclusive", strconv.FormatBool(*p.Inclusive))
	}
	if p.Latest != nil {
		uv.Add("latest", *p.Latest)
	}
	if p.Limit != nil {
		uv.Add("limit", strconv.Itoa(*p.Limit))
	}
	if p.Oldest != nil {
		uv.Add("oldest", *p.Oldest)
	}

	return strings.NewReader(uv.Encode()), nil
}
