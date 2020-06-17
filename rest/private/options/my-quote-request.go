package options

import (
	"fmt"
	"net/http"
)

type RequestForMyOpQuoteRequest struct {
	RequestID int
}

type ResponseForMyOpQuoteRequest []Quote

func (req *RequestForMyOpQuoteRequest) Path() string {
	return fmt.Sprintf("/options/requests/%d/quotes", req.RequestID)
}

func (req *RequestForMyOpQuoteRequest) Method() string {
	return http.MethodGet
}

func (req *RequestForMyOpQuoteRequest) Query() string {
	return ""
}

func (req *RequestForMyOpQuoteRequest) Payload() []byte {
	return nil
}
