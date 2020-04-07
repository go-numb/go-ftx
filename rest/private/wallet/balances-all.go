package wallet

import (
	"net/http"
)

type RequestForBalancesAll struct {
}

type ResponseForBalancesAll map[string][]Balance

func (req *RequestForBalancesAll) Path() string {
	return "/wallet/all_balances"
}

func (req *RequestForBalancesAll) Method() string {
	return http.MethodGet
}

func (req *RequestForBalancesAll) Query() string {
	return ""
}

func (req *RequestForBalancesAll) Payload() []byte {
	return nil
}
