package api

import (
	"context"

	"github.com/michimani/jsc/slack"
	"github.com/michimani/jsc/slack/types"
)

const (
	getUsersInfoEndpoint = "https://slack.com/api/users.info"
)

// https://api.slack.com/methods/users.info
func GetUsersInfo(ctx context.Context, c *slack.SlackClient, p *types.GetUsersInfoParameter) (*types.GetUsersInfoResponse, error) {
	tc := slack.NewTypedClient[types.GetUsersInfoParameter, types.GetUsersInfoResponse](c)

	return tc.CallAPI(ctx, "GET", getUsersInfoEndpoint, p)
}
