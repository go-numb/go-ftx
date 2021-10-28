package orders

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForModifyTriggerOrder struct {
	OrderID      string  `json:"-"`
	TrailValue   float64 `json:"trailValue,omitempty"`
	TriggerPrice float64 `json:"triggerPrice,omitempty"`
	//	OrderPrice only for stop limit orders
	OrderPrice float64 `json:"orderPrice,omitempty"`
	// necessary
	Size float64 `json:"size"`
}

type ResponseForModifyTriggerOrder struct {
	OrderID   int    `json:"orderId"`
	Type      string `json:"type"`
	OrderType string `json:"orderType"`
	Future    string `json:"future"`
	Market    string `json:"market"`
	Status    string `json:"status"`
	Error     string `json:"error"`
	Side      string `json:"side"`

	OrderPrice   float64 `json:"orderPrice"`
	AvgFillPrice float64 `json:"avgFillPrice"`
	TriggerPrice float64 `json:"triggerPrice"`
	Size         float64 `json:"size"`
	FilledSize   float64 `json:"filledSize"`

	ID               int  `json:"id"`
	RetryUntilFilled bool `json:"retryUntilFilled"`
	ReduceOnly       bool `json:"reduceOnly"`

	TriggeredAt time.Time `json:"triggeredAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (req *RequestForModifyTriggerOrder) Path() string {
	return fmt.Sprintf("/conditional_orders/%s/modify", req.OrderID)
}

func (req *RequestForModifyTriggerOrder) Method() string {
	return http.MethodPost
}

func (req *RequestForModifyTriggerOrder) Query() string {
	return ""
}

func (req *RequestForModifyTriggerOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
