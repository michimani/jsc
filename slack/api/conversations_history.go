package api

import (
	"context"

	"github.com/michimani/jsc"
	"github.com/michimani/jsc/slack/types"
)

const (
	getConversationsHistoryEndpoint = "https://slack.com/api/conversations.history"
)

// https://api.slack.com/methods/conversations.history
func GetConversationsHistory(ctx context.Context, c *jsc.SlackClient, p *types.ConversationsHistoryParameter) (*types.ConversationsHistoryResponse, error) {
	tc := jsc.NewTypedClient[types.ConversationsHistoryParameter, types.ConversationsHistoryResponse](c)

	return tc.CallAPI(ctx, "GET", getConversationsHistoryEndpoint, p)
}
