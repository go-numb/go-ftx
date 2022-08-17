package orders

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForModifyOrder struct {
	OrderID  string  `json:"-"`
	ClientID string  `json:"-"`
	Price    float64 `json:"price,omitempty"`
	Size     float64 `json:"size,omitempty"`
}

type ResponseForModifyOrder struct {
	ClientID string `json:"clientId"`
	Future   string `json:"future"`
	Market   string `json:"market"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Side     string `json:"side"`

	Price         float64 `json:"price"`
	Size          float64 `json:"size"`
	FilledSize    float64 `json:"filledSize"`
	RemainingSize float64 `json:"remainingSize"`

	ID         int       `json:"id"`
	ReduceOnly bool      `json:"reduceOnly"`
	Ioc        bool      `json:"ioc"`
	PostOnly   bool      `json:"postOnly"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (req *RequestForModifyOrder) Path() string {
	// prioritize ClientID
	if req.ClientID != "" {
		return fmt.Sprintf("/orders/by_client_id/%s/modify", req.ClientID)
	}
	return fmt.Sprintf("/orders/%s/modify", req.OrderID)
}

func (req *RequestForModifyOrder) Method() string {
	return http.MethodPost
}

func (req *RequestForModifyOrder) Query() string {
	return ""
}

func (req *RequestForModifyOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
