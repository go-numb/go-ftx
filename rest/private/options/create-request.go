package options

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForCreateOpQuoteRequest struct {
	Underlying string  `json:"underlying"`
	Type       string  `json:"type"`
	Side       string  `json:"side"`
	Strike     float64 `json:"strike"`
	Size       float64 `json:"size"`
	Expiry     int64   `json:"expiry"`
	// Options
	LimitPrice     float64 `json:"limitPrice,omitempty"`
	HideLimitPrice bool    `json:"hideLimitPrice,omitempty"`
	RequestExpiry  int64   `json:"requestExpiry,omitempty"`
	CounterpartyID int     `json:"counterpartyId,omitempty"`
}

type ResponseForCreateOpQuoteRequest struct {
	ID     int `json:"id"`
	Option struct {
		Strike     float64   `json:"strike"`
		Type       string    `json:"type"`
		Underlying string    `json:"underlying"`
		Expiry     time.Time `json:"expiry"`
	} `json:"option"`
	Strike        int       `json:"strike"`
	Type          string    `json:"type"`
	Underlying    string    `json:"underlying"`
	Side          string    `json:"side"`
	Status        string    `json:"status"`
	Size          float64   `json:"size"`
	Expiry        time.Time `json:"expiry"`
	RequestExpiry time.Time `json:"requestExpiry"`
	Time          time.Time `json:"time"`
}

func (req *RequestForCreateOpQuoteRequest) Path() string {
	return "/options/requests"
}

func (req *RequestForCreateOpQuoteRequest) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateOpQuoteRequest) Query() string {
	return ""
}

func (req *RequestForCreateOpQuoteRequest) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
