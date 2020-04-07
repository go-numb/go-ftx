package rest

import (
	"time"

	"github.com/go-numb/go-ftx/auth"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

const ENDPOINT = "https://ftx.com/api"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	Auth *auth.Config

	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func New(auth *auth.Config) *Client {
	hc := new(fasthttp.Client)

	return &Client{
		Auth:        auth,
		HTTPC:       hc,
		HTTPTimeout: 5 * time.Second,
	}
}
