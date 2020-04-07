package options

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForOpFills struct {
	Limit int   `url:"limit,omitempty"`
	Start int64 `url:"start_time,omitempty"`
	End   int64 `url:"end_time,omitempty"`
}

type ResponseForOpFills []OpTrade

type OpFill struct {
	ID        int       `json:"id"`
	Liquidity string    `json:"liquidity"`
	Side      string    `json:"side"`
	Price     float64   `json:"price"`
	QuoteID   float64   `json:"quoteId"`
	Size      float64   `json:"size"`
	Fee       float64   `json:"fee"`
	FeeRate   float64   `json:"feeRate"`
	Option    Option    `json:"option"`
	Time      time.Time `json:"time"`
}

func (req *RequestForOpFills) Path() string {
	return "/options/fills"
}

func (req *RequestForOpFills) Method() string {
	return http.MethodGet
}

func (req *RequestForOpFills) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForOpFills) Payload() []byte {
	return nil
}
