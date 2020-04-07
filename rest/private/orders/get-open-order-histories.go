package orders

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForHistories struct {
	ProductCode string `url:"market,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	Start       int64  `url:"start_time,omitempty"`
	End         int64  `url:"end_time,omitempty"`
}

type ResponseForHistories []Histories

type Histories struct {
	Type     string `json:"type"`
	ClientID string `json:"clientId"`
	Future   string `json:"future"`
	Market   string `json:"market"`
	Status   string `json:"status"`
	Side     string `json:"side"`

	Price         float64 `json:"price"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	FilledSize    float64 `json:"filledSize"`
	Size          float64 `json:"size"`
	RemainingSize float64 `json:"remainingSize"`

	ID         int  `json:"id"`
	Ioc        bool `json:"ioc"`
	PostOnly   bool `json:"postOnly"`
	ReduceOnly bool `json:"reduceOnly"`

	CreatedAt time.Time `json:"createdAt"`
}

func (req *RequestForHistories) Path() string {
	return "/orders/history"
}

func (req *RequestForHistories) Method() string {
	return http.MethodGet
}

func (req *RequestForHistories) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForHistories) Payload() []byte {
	return nil
}
