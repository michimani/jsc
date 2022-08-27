package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/michimani/jsc/slack"
	"github.com/michimani/jsc/slack/api"
	"github.com/michimani/jsc/slack/types"
)

type UserPool struct {
	client  *slack.SlackClient
	userMap map[string]string
}

func NewUserPool(c *slack.SlackClient) *UserPool {
	return &UserPool{
		client:  c,
		userMap: map[string]string{},
	}
}

func (up *UserPool) GetUserName(id string) (string, error) {
	if up == nil {
		return "", errors.New("UserPool is nil")
	}

	if name, ok := up.userMap[id]; ok {
		return name, nil
	}

	user, err := up.fetchUser(id)
	if err != nil {
		return "", err
	}

	up.userMap[id] = *user.RealName

	return up.userMap[id], nil
}

func (up *UserPool) fetchUser(id string) (*types.UserInfo, error) {
	if up == nil {
		return nil, errors.New("UserPool is nil")
	}

	res, err := api.GetUsersInfo(
		context.Background(),
		up.client,
		&types.GetUsersInfoParameter{
			User: &id,
		},
	)

	if err != nil {
		return nil, err
	}

	if !*res.OK {
		return nil, fmt.Errorf("failed to get user info: %s", *res.Error)
	}

	return &res.User, nil
}
