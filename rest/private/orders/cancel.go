package orders

import (
	"fmt"
	"net/http"
)

type RequestForCancelByID struct {
	ClientID       string
	OrderID        int
	TriggerOrderID string
}

type ResponseForCancelByID string

func (req *RequestForCancelByID) Path() string {
	if req.TriggerOrderID != "" {
		return fmt.Sprintf("/conditional_orders/%s", req.TriggerOrderID)
	} else if req.ClientID != "" {
		return fmt.Sprintf("/orders/by_client_id/%s", req.ClientID)
	}
	return fmt.Sprintf("/orders/%d", req.OrderID)
}

func (req *RequestForCancelByID) Method() string {
	return http.MethodDelete
}

func (req *RequestForCancelByID) Query() string {
	return ""
}

func (req *RequestForCancelByID) Payload() []byte {
	return nil
}
