package orders

import (
	"net/http"
)

type RequestForCancelAll struct {
	ProductCode           string `json:"market,omitempty"`
	Side                  string `json:"side,omitempty"`
	ConditionalOrdersOnly bool   `json:"conditionalOrdersOnly,omitempty"`
	LimitOrdersOnly       bool   `json:"limitOrdersOnly,omitempty"`
}

type ResponseForCancelAll string

func (req *RequestForCancelAll) Path() string {
	return "/orders"
}

func (req *RequestForCancelAll) Method() string {
	return http.MethodDelete
}

func (req *RequestForCancelAll) Query() string {
	return ""
}

func (req *RequestForCancelAll) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
