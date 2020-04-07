package orders

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForOrderTriggers struct {
	CID string `url:"-"`
}

type ResponseForOrderTriggers []Trigger

type Trigger struct {
	FilledSize float64 `json:"filledSize"`
	OrderSize  float64 `json:"orderSize"`
	Error      string  `json:"error"`

	OrderID int       `json:"orderId"`
	Time    time.Time `json:"time"`
}

func (req *RequestForOrderTriggers) Path() string {
	return fmt.Sprintf("/conditional_orders/%s/triggers", req.CID)
}

func (req *RequestForOrderTriggers) Method() string {
	return http.MethodGet
}

func (req *RequestForOrderTriggers) Query() string {
	return ""
}

func (req *RequestForOrderTriggers) Payload() []byte {
	return nil
}
