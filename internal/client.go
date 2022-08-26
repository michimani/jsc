package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type SlackClient struct {
	client *http.Client
	token  string
}

type NewSlackClientInput struct {
	Client *http.Client
	Token  string
}

var defaultHTTPClient = &http.Client{
	Timeout: time.Duration(30) * time.Second,
}

func NewSlackClient(in *NewSlackClientInput) (*SlackClient, error) {
	if in.Token == "" {
		return nil, errors.New("Token is required for creating Slack client.")
	}

	c := SlackClient{
		token:  in.Token,
		client: defaultHTTPClient,
	}

	if in.Client != nil {
		c.client = in.Client
	}

	return &c, nil
}

func (c *SlackClient) Token() string {
	return c.token
}

const contentType = "application/x-www-form-urlencoded"

func (c *SlackClient) Do(ctx context.Context, method string, endpoint string, p IParameter, r IResponse) error {
	body, err := p.Body()
	if err != nil {
		return err
	}

	if method == "GET" && body != nil {
		queryBytes, _ := io.ReadAll(body)
		endpoint = fmt.Sprintf("%s?%s", endpoint, string(queryBytes))
		body = strings.NewReader("")
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token()))
	req.Header.Set("Content-Type", contentType)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		if err != io.EOF {
			return nil
		}
		return err
	}

	return nil
}
