package api

import (
	"context"
	"jsc/internal"
	"jsc/slack/types"
)

const (
	getConversationsHistoryEndpoint = "https://slack.com/api/conversations.history"
)

// https://api.slack.com/methods/conversations.history
func GetConversationsHistory(ctx context.Context, c *internal.SlackClient, p *types.ConversationsHistoryParameter) (*types.ConversationsHistoryResponse, error) {
	res := &types.ConversationsHistoryResponse{}
	if err := c.Do(ctx, "GET", getConversationsHistoryEndpoint, p, res); err != nil {
		return nil, err
	}

	return res, nil
}
