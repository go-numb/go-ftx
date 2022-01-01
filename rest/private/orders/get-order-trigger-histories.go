package orders

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForOrderTriggerHistories struct {
	ProductCode string `url:"market,omitempty"`
	Type        string `url:"type,omitempty"`
	Side        string `url:"side,omitempty"`
	OrderType   string `url:"orderType,omitempty"`

	Start int64 `url:"start_time,omitempty"`
	End   int64 `url:"end_time,omitempty"`
}

type ResponseForOrderTriggerHistories []TriggerHistory

type TriggerHistory struct {
	Status       string  `json:"status"`
	Error        string  `json:"error"`
	Future       string  `json:"future"`
	Market       string  `json:"market"`
	Type         string  `json:"type"`
	OrderType    string  `json:"orderType"`
	OrderStatus  string  `json:"orderStatus"`
	Side         string  `json:"side"`
	OrderPrice   float64 `json:"orderPrice"`
	TriggerPrice float64 `json:"triggerPrice"`
	AvgFillPrice float64 `json:"avgFillPrice"`
	TrailStart   float64 `json:"trailStart"`
	TrailValue   float64 `json:"trailValue"`

	ID               int       `json:"id"`
	OrderID          int       `json:"orderId"`
	Size             float64   `json:"size"`
	FilledSize       float64   `json:"filledSize"`
	RetryUntilFilled bool      `json:"retryUntilFilled"`
	ReduceOnly       bool      `json:"reduceOnly"`
	TriggeredAt      time.Time `json:"triggeredAt"`
	CreatedAt        time.Time `json:"createdAt"`
}

func (req *RequestForOrderTriggerHistories) Path() string {
	return "/conditional_orders/history"
}

func (req *RequestForOrderTriggerHistories) Method() string {
	return http.MethodGet
}

func (req *RequestForOrderTriggerHistories) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForOrderTriggerHistories) Payload() []byte {
	return nil
}
