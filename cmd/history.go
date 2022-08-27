package cmd

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/michimani/jsc/slack"
	"github.com/michimani/jsc/slack/api"
	"github.com/michimani/jsc/slack/types"
)

type OutputType string

const (
	OutputTypeTimeline = OutputType("timeline")
	OutputTypeJSON     = OutputType("json")
)

type HistoryInput struct {
	// history limit of each channel
	Limit int
	// channel IDs
	ChannelIDs []string
	// output type (timeline or json)
	OutputType OutputType
}

type HistoryMessage struct {
	ChannelName string    `json:"channel_name"`
	Username    string    `json:"username"`
	MessageURL  string    `json:"message_url"`
	Text        string    `json:"text"`
	PostedAt    time.Time `json:"posted_at"`

	// use for sort
	ts float64 `json:"-"`
}

type JoinedHistory struct {
	Messages []HistoryMessage
}

type ChannelHistory struct {
	Channel  string
	Messages []types.ConversationsHistoryMessage
}

func GetJoinedHistory(ctx context.Context, c *slack.SlackClient, in *HistoryInput) (*JoinedHistory, error) {
	histories := make([]*ChannelHistory, len(in.ChannelIDs))

	for i, channel := range in.ChannelIDs {
		res, err := api.GetConversationsHistory(
			ctx,
			c,
			&types.ConversationsHistoryParameter{
				Channel: &channel,
				Limit:   &in.Limit,
			},
		)

		if err != nil {
			return nil, fmt.Errorf("failed to get history in %s: %v", channel, err)
		}

		if !*res.OK {
			return nil, fmt.Errorf("failed to get history in %s: %s", channel, *res.Error)
		}

		histories[i] = &ChannelHistory{
			Channel:  channel,
			Messages: res.Messages,
		}
	}

	return joined(histories...)
}

func joined(chs ...*ChannelHistory) (*JoinedHistory, error) {
	total := 0
	for _, ch := range chs {
		total += len(ch.Messages)
	}

	messages := make([]HistoryMessage, total)

	idx := 0
	for _, ch := range chs {
		for _, m := range ch.Messages {
			hm, err := toHistoryMessage(ch.Channel, m)
			if err != nil {
				return nil, err
			}
			messages[idx] = *hm
			idx++
		}
	}

	sort.Slice(messages, func(i, j int) bool { return messages[i].ts < messages[j].ts })

	return &JoinedHistory{
		Messages: messages,
	}, nil
}

func toHistoryMessage(channel string, message types.ConversationsHistoryMessage) (*HistoryMessage, error) {
	// TODO: get url by https://api.slack.com/methods/chat.getPermalink
	messageURL := message.Ts.ToID()
	postedAt, err := message.Ts.ToTime()
	if err != nil {
		return nil, err
	}
	ts, err := message.Ts.Float64()
	if err != nil {
		return nil, err
	}

	// TODO: get name by https://api.slack.com/methods/users.info
	username := ""
	if message.Username != nil {
		username = *message.Username
	} else if message.User != nil {
		username = *message.User
	}

	return &HistoryMessage{
		ChannelName: channel,
		Username:    username,
		Text:        *message.Text,
		MessageURL:  messageURL,
		PostedAt:    *postedAt,
		ts:          ts,
	}, nil
}
