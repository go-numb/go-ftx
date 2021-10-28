package orders

import (
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForPlaceOrder struct {
	ClientID          string  `json:"clientId,omitempty"`
	Type              string  `json:"type"`
	Market            string  `json:"market"`
	Side              string  `json:"side"`
	Price             float64 `json:"price"`
	Size              float64 `json:"size"`
	ReduceOnly        bool    `json:"reduceOnly,omitempty"`
	Ioc               bool    `json:"ioc,omitempty"`
	PostOnly          bool    `json:"postOnly,omitempty"`
	RejectOnPriceBand bool    `json:"rejectOnPriceBand,omitempty"`
}

type ResponseForPlaceOrder struct {
	ClientID string `json:"clientId"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Future   string `json:"future"`
	Market   string `json:"market"`
	Side     string `json:"side"`

	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	RemainingSize float64 `json:"remainingSize"`
	FilledSize    float64 `json:"filledSize"`

	ID         int  `json:"id"`
	Ioc        bool `json:"ioc"`
	PostOnly   bool `json:"postOnly"`
	ReduceOnly bool `json:"reduceOnly"`

	CreatedAt time.Time `json:"createdAt"`
}

func (req *RequestForPlaceOrder) Path() string {
	return "/orders"
}

func (req *RequestForPlaceOrder) Method() string {
	return http.MethodPost
}

func (req *RequestForPlaceOrder) Query() string {
	return ""
}

func (req *RequestForPlaceOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
