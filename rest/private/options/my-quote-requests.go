package options

import (
	"net/http"
	"time"
)

type RequestForMyOpQuoteRequests struct {
}

type ResponseForMyOpQuoteRequests []MyOpQuote

type MyOpQuote struct {
	ID     int `json:"id"`
	Option struct {
		Underlying string    `json:"underlying"`
		Type       string    `json:"type"`
		Strike     float64   `json:"strike"`
		Expiry     time.Time `json:"expiry"`
	} `json:"option"`
	Status         string  `json:"status"`
	Side           string  `json:"side"`
	Size           float64 `json:"size"`
	LimitPrice     float64 `json:"limitPrice"`
	HideLimitPrice bool    `json:"hideLimitPrice"`
	Quotes         []struct {
		ID          int       `json:"id"`
		Collateral  float64   `json:"collateral"`
		Price       float64   `json:"price"`
		Status      string    `json:"status"`
		QuoteExpiry time.Time `json:"quoteExpiry"`
		Time        time.Time `json:"time"`
	} `json:"quotes"`

	Time          time.Time `json:"time"`
	RequestExpiry time.Time `json:"requestExpiry"`
}

func (req *RequestForMyOpQuoteRequests) Path() string {
	return "/options/my_requests"
}

func (req *RequestForMyOpQuoteRequests) Method() string {
	return http.MethodGet
}

func (req *RequestForMyOpQuoteRequests) Query() string {
	return ""
}

func (req *RequestForMyOpQuoteRequests) Payload() []byte {
	return nil
}
