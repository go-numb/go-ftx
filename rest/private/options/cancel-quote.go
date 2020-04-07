package options

import (
	"fmt"
	"net/http"
)

type RequestForCancelOpQuote struct {
	QuoteID int
}

type ResponseForCancelOpQuote Quote

func (req *RequestForCancelOpQuote) Path() string {
	return fmt.Sprintf("/options/quotes/%d", req.QuoteID)
}

func (req *RequestForCancelOpQuote) Method() string {
	return http.MethodDelete
}

func (req *RequestForCancelOpQuote) Query() string {
	return ""
}

func (req *RequestForCancelOpQuote) Payload() []byte {
	return nil
}
