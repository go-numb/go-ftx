package convert

import (
	"fmt"
	"net/http"
)

type RequestForQuoteStatus struct {
	QuoteId int `json:"quoteId"`
}

type ResponseForQuoteStatus struct {
	BaseCoin  string  `json:"baseCoin"`
	Cost      float64 `json:"cost"`
	Expired   bool    `json:"expired"`
	Expiry    float64 `json:"expiry"`
	Filled    bool    `json:"filled"`
	FromCoin  string  `json:"fromCoin"`
	QuoteId   int     `json:"id"`
	Price     float64 `json:"price"`
	Proceeds  float64 `json:"proceeds"`
	QuoteCoin string  `json:"quoteCoin"`
	Side      string  `json:"side"`
	ToCoin    string  `json:"toCoin"`
}

func (req *RequestForQuoteStatus) Path() string {
	return fmt.Sprintf("/otc/quotes/%d", req.QuoteId)
}

func (req *RequestForQuoteStatus) Method() string {
	return http.MethodGet
}

func (req *RequestForQuoteStatus) Query() string {
	return ""
}

func (req *RequestForQuoteStatus) Payload() []byte {
	return nil
}
