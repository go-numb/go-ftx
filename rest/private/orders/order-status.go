package orders

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForOrderStatus struct {
	ClientID string
	OrderID  int
}

type ResponseForOrderStatus struct {
	ClientID string `json:"clientId"`
	Future   string `json:"future"`
	Market   string `json:"market"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Side     string `json:"side"`

	Price         float64 `json:"price"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	FilledSize    float64 `json:"filledSize"`
	Size          float64 `json:"size"`
	RemainingSize float64 `json:"remainingSize"`

	ID         int  `json:"id"`
	Ioc        bool `json:"ioc"`
	ReduceOnly bool `json:"reduceOnly"`
	PostOnly   bool `json:"postOnly"`

	CreatedAt time.Time `json:"createdAt"`
}

func (req *RequestForOrderStatus) Path() string {
	if req.ClientID != "" {
		return fmt.Sprintf("/orders/by_client_id/%s", req.ClientID)
	}
	return fmt.Sprintf("/orders/%d", req.OrderID)
}

func (req *RequestForOrderStatus) Method() string {
	return http.MethodGet
}

func (req *RequestForOrderStatus) Query() string {
	return ""
}

func (req *RequestForOrderStatus) Payload() []byte {
	return nil
}
