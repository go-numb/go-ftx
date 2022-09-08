package convert

import (
	"fmt"
	"net/http"
)

type RequestForAcceptQuote struct {
	QuoteId int `json:"quoteId"`
}

type ResponseForAcceptQuote struct {
	// Filled bool `json:"filled"`
}

func (req *RequestForAcceptQuote) Path() string {
	return fmt.Sprintf("/otc/quotes/%d/accept", req.QuoteId)
}

func (req *RequestForAcceptQuote) Method() string {
	return http.MethodPost
}

func (req *RequestForAcceptQuote) Query() string {
	return ""
}

func (req *RequestForAcceptQuote) Payload() []byte {
	return nil
}
