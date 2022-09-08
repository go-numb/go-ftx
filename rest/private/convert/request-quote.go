package convert

import (
	"encoding/json"
	"net/http"
)

type RequestForRequestQuote struct {
	FromCoin string  `json:"fromCoin"`
	ToCoin   string  `json:"toCoin"`
	Size     float64 `json:"size"`
	// WaitForPrice bool    `json:"waitForPrice,omitempty"`
}

type ResponseForRequestQuote struct {
	QuoteId int `json:"quoteId"`
}

func (req *RequestForRequestQuote) Path() string {
	return "/otc/quotes"
}

func (req *RequestForRequestQuote) Method() string {
	return http.MethodPost
}

func (req *RequestForRequestQuote) Query() string {
	return ""
}

func (req *RequestForRequestQuote) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
