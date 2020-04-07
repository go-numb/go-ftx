package subaccount

import (
	"fmt"
	"net/http"
)

type RequestForBalanceSubAccount struct {
	NickName string
}

type ResponseForBalanceSubAccount []Balance

type Balance struct {
	Coin  string  `json:"coin"`
	Free  float64 `json:"free"`
	Total float64 `json:"total"`
}

func (req *RequestForBalanceSubAccount) Path() string {
	return fmt.Sprintf("/subaccounts/%s/balances", req.NickName)
}

func (req *RequestForBalanceSubAccount) Method() string {
	return http.MethodGet
}

func (req *RequestForBalanceSubAccount) Query() string {
	return ""
}

func (req *RequestForBalanceSubAccount) Payload() []byte {
	return nil
}
