package api

import (
	"context"

	"github.com/michimani/jsc/slack"
	"github.com/michimani/jsc/slack/types"
)

const (
	getChatPermalinkEndpoint = "https://slack.com/api/chat.getPermalink"
)

// https://api.slack.com/methods/chat.getPermalink
func GetChatPermalink(ctx context.Context, c *slack.SlackClient, p *types.GetChatPermalinkParameter) (*types.GetChatPermalinkResponse, error) {
	tc := slack.NewTypedClient[types.GetChatPermalinkParameter, types.GetChatPermalinkResponse](c)

	return tc.CallAPI(ctx, "GET", getChatPermalinkEndpoint, p)
}
