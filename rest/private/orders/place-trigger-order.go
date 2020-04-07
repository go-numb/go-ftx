package orders

import (
	"net/http"
	"time"
)

type RequestForPlaceTriggerOrder struct {
	Market string `json:"market"`
	// Type stop, trailingStop, takeProfit. default is stop
	Type             string  `json:"type"`
	Side             string  `json:"side"`
	TriggerPrice     float64 `json:"triggerPrice"`
	Size             float64 `json:"size"`
	ReduceOnly       bool    `json:"reduceOnly,omitempty"`
	RetryUntilFilled bool    `json:"retryUntilFilled,omitempty"`
}

type ResponseForPlaceTriggerOrder struct {
	Type      string `json:"type"`
	Future    string `json:"future"`
	Market    string `json:"market"`
	OrderID   string `json:"orderId"`
	Status    string `json:"status"`
	Error     string `json:"error"`
	OrderType string `json:"orderType"`
	Side      string `json:"side"`

	TriggerPrice float64 `json:"triggerPrice"`
	Size         float64 `json:"size"`

	OrderPrice float64 `json:"orderPrice"`

	ID               int  `json:"id"`
	ReduceOnly       bool `json:"reduceOnly"`
	RetryUntilFilled bool `json:"retryUntilFilled"`

	TriggeredAt time.Time `json:"triggeredAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (req *RequestForPlaceTriggerOrder) Path() string {
	return "/conditional_orders"
}

func (req *RequestForPlaceTriggerOrder) Method() string {
	return http.MethodPost
}

func (req *RequestForPlaceTriggerOrder) Query() string {
	return ""
}

func (req *RequestForPlaceTriggerOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
