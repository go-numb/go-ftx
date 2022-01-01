package orders

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForOpenOrder struct {
	ProductCode string `url:"market,omitempty"`
}

type ResponseForOpenOrder []OpenOrder

type OpenOrder struct {
	Type     string `json:"type"`
	Market   string `json:"market"`
	Future   string `json:"future"`
	ClientID string `json:"clientId"`
	Status   string `json:"status"`
	Side     string `json:"side"`

	Price         float64 `json:"price"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	Size          float64 `json:"size"`
	FilledSize    float64 `json:"filledSize"`
	RemainingSize float64 `json:"remainingSize"`

	ID         int64 `json:"id"`
	ReduceOnly bool  `json:"reduceOnly"`
	Ioc        bool  `json:"ioc"`
	PostOnly   bool  `json:"postOnly"`

	CreatedAt time.Time `json:"createdAt"`
}

type Order struct {
	ID            int64     `json:"id"`
	ClientID      string    `json:"clientId"`
	Market        string    `json:"market"`
	Type          string    `json:"type"`
	Side          string    `json:"side"`
	Size          float64   `json:"size"`
	Price         float64   `json:"price"`
	ReduceOnly    bool      `json:"reduceOnly"`
	Ioc           bool      `json:"ioc"`
	PostOnly      bool      `json:"postOnly"`
	Status        string    `json:"status"`
	FilledSize    float64   `json:"filledSize"`
	RemainingSize float64   `json:"remainingSize"`
	AvgFillPrice  float64   `json:"avgFillPrice"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (req *RequestForOpenOrder) Path() string {
	return "/orders"
}

func (req *RequestForOpenOrder) Method() string {
	return http.MethodGet
}

func (req *RequestForOpenOrder) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForOpenOrder) Payload() []byte {
	return nil
}
