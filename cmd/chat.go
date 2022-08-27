package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/michimani/jsc/slack"
	"github.com/michimani/jsc/slack/api"
	"github.com/michimani/jsc/slack/types"
)

type DomainPool struct {
	client    *slack.SlackClient
	domainMap map[string]string
}

func NewDomainPool(c *slack.SlackClient) *DomainPool {
	return &DomainPool{
		client:    c,
		domainMap: map[string]string{},
	}
}

func (dp *DomainPool) GetDomain(cid string, mts string) (string, error) {
	if dp == nil {
		return "", errors.New("DomainPool is nil")
	}

	if domain, ok := dp.domainMap[cid]; ok {
		return domain, nil
	}

	domain, err := dp.fetchDomain(cid, mts)
	if err != nil {
		return "", err
	}

	dp.domainMap[cid] = domain

	return domain, nil
}

func (dp *DomainPool) fetchDomain(cid string, mts string) (string, error) {
	if dp == nil {
		return "", errors.New("DomainPool is nil")
	}

	res, err := api.GetChatPermalink(
		context.Background(),
		dp.client,
		&types.GetChatPermalinkParameter{
			Channel:   &cid,
			MessageTs: &mts,
		},
	)

	if err != nil {
		return "", err
	}

	if !*res.OK {
		return "", fmt.Errorf("failed to get domain: %s", *res.Error)
	}

	tmp := (*res.Permalink)[8:]
	domain := tmp[:strings.Index(tmp, "/")]

	return domain, nil
}
