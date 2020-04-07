package options

import (
	"fmt"
	"net/http"
)

type RequestForAcceptOpQuote struct {
	QuoteID int
}

type ResponseForAcceptOpQuote Quote

func (req *RequestForAcceptOpQuote) Path() string {
	return fmt.Sprintf("/options/quotes/%d/accept", req.QuoteID)
}

func (req *RequestForAcceptOpQuote) Method() string {
	return http.MethodPost
}

func (req *RequestForAcceptOpQuote) Query() string {
	return ""
}

func (req *RequestForAcceptOpQuote) Payload() []byte {
	return nil
}
