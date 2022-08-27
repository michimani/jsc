package jsc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/michimani/jsc/slack/types"
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

type TypedSlackClient[P types.IParameter, R types.IResponse] struct {
	Base *SlackClient
}

func NewTypedClient[P types.IParameter, R types.IResponse](c *SlackClient) *TypedSlackClient[P, R] {
	return &TypedSlackClient[P, R]{
		Base: c,
	}
}

const contentType = "application/x-www-form-urlencoded"

func (c *TypedSlackClient[P, R]) CallAPI(ctx context.Context, method string, endpoint string, p types.IParameter) (*R, error) {
	body := c.urlencodedParameters(p)

	if method == "GET" && body != nil {
		queryBytes, _ := io.ReadAll(body)
		endpoint = fmt.Sprintf("%s?%s", endpoint, string(queryBytes))
		body = strings.NewReader("")
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Base.Token()))
	req.Header.Set("Content-Type", contentType)

	res, err := c.Base.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := new(R)
	if err := json.NewDecoder(res.Body).Decode(r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *TypedSlackClient[P, R]) urlencodedParameters(p any) *strings.Reader {
	uv := url.Values{}

	rv := reflect.ValueOf(*p.(*P))
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.FieldByName(field.Name)

		if value.IsNil() {
			continue
		}

		apiArgKey := field.Tag.Get("paramkey")
		apiArgValue := ""

		switch reflect.TypeOf(value.Interface()).String() {
		case "*string":
			apiArgValue = *value.Interface().(*string)
		case "*int":
			apiArgValue = strconv.Itoa(*value.Interface().(*int))
		case "*bool":
			apiArgValue = strconv.FormatBool(*value.Interface().(*bool))
		default:
			// noop
		}

		uv.Add(apiArgKey, apiArgValue)
	}

	return strings.NewReader(uv.Encode())
}
