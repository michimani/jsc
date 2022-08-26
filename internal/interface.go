package internal

import "strings"

type IParameter interface {
	Body() (*strings.Reader, error)
	SetToken(token string)
}

type IResponse interface {
}
