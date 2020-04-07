package leveraged

import (
	"net/http"
)

type RequestForLvBalances struct {
}

type ResponseForLvBalances []TokenBalance

type TokenBalance struct {
	Token   string  `json:"token"`
	Balance float64 `json:"balance"`
}

func (req *RequestForLvBalances) Path() string {
	return "/lt/balances"
}

func (req *RequestForLvBalances) Method() string {
	return http.MethodGet
}

func (req *RequestForLvBalances) Query() string {
	return ""
}

func (req *RequestForLvBalances) Payload() []byte {
	return nil
}
