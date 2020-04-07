package options

import (
	"net/http"
	"time"
)

type RequestForOpQuoteRequests struct {
}

type ResponseForOpQuoteRequests []OpQuote

type OpQuote struct {
	ID     int `json:"id"`
	Option struct {
		Underlying string    `json:"underlying"`
		Type       string    `json:"type"`
		Strike     float64   `json:"strike"`
		Expiry     time.Time `json:"expiry"`
	} `json:"option"`
	Side   string  `json:"side"`
	Size   float64 `json:"size"`
	Status string  `json:"status"`

	Time          time.Time `json:"time"`
	RequestExpiry time.Time `json:"requestExpiry"`
}

func (req *RequestForOpQuoteRequests) Path() string {
	return "/options/requests"
}

func (req *RequestForOpQuoteRequests) Method() string {
	return http.MethodGet
}

func (req *RequestForOpQuoteRequests) Query() string {
	return ""
}

func (req *RequestForOpQuoteRequests) Payload() []byte {
	return nil
}
