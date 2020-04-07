package options

import (
	"net/http"
)

type RequestForMyOpQuotes struct {
}

type ResponseForMyOpQuotes []Quote

func (req *RequestForMyOpQuotes) Path() string {
	return "/options/my_quotes"
}

func (req *RequestForMyOpQuotes) Method() string {
	return http.MethodGet
}

func (req *RequestForMyOpQuotes) Query() string {
	return ""
}

func (req *RequestForMyOpQuotes) Payload() []byte {
	return nil
}
