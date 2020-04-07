package options

import (
	"fmt"
	"net/http"
)

// Document記載はCreate quoteだが、modify request
type RequestForModifyOpQuoteRequest struct {
	RequestID int     `json:"-"`
	Price     float64 `json:"price"`
}

type ResponseForModifyOpQuoteRequest Quote

func (req *RequestForModifyOpQuoteRequest) Path() string {
	return fmt.Sprintf("/options/requests/%s/quotes", req.RequestID)
}

func (req *RequestForModifyOpQuoteRequest) Method() string {
	return http.MethodPost
}

func (req *RequestForModifyOpQuoteRequest) Query() string {
	return ""
}

func (req *RequestForModifyOpQuoteRequest) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
