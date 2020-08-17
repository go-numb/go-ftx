package orders

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForOpenTriggerOrders struct {
	ProductCode string `url:"market,omitempty"`
	Type        string `url:"type,omitempty"`
}

type ResponseForOpenTriggerOrders []OpenTriggerOrder

type OpenTriggerOrder struct {
	Type      string `json:"type"`
	OrderType string `json:"orderType"`
	OrderID   string `json:"orderId"`
	Status    string `json:"status"`
	Error     string `json:"error"`
	Market    string `json:"market"`
	Future    string `json:"future"`
	Side      string `json:"side"`

	OrderPrice float64 `json:"orderPrice"`
	Size       float64 `json:"size"`

	TrailStart   float64 `json:"trailStart"`
	TrailValue   float64 `json:"trailValue"`
	TriggerPrice float64 `json:"triggerPrice"`
	AvgFillPrice float64 `json:"avgFillPrice"`

	TriggeredAt time.Time `json:"triggeredAt"`
	CreatedAt   time.Time `json:"createdAt"`

	ReduceOnly       bool    `json:"reduceOnly"`
	RetryUntilFilled bool    `json:"retryUntilFilled"`
	FilledSize       float64 `json:"filledSize"`
	ID               int     `json:"id"`
}

func (req *RequestForOpenTriggerOrders) Path() string {
	return "/conditional_orders"
}

func (req *RequestForOpenTriggerOrders) Method() string {
	return http.MethodGet
}

func (req *RequestForOpenTriggerOrders) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForOpenTriggerOrders) Payload() []byte {
	return nil
}
