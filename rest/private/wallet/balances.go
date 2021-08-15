package wallet

import (
	"net/http"
)

type RequestForBalances struct {
}

type ResponseForBalances []Balance

type Balance struct {
	Coin                   string  `json:"coin"`
	Free                   float64 `json:"free"`
	Total                  float64 `json:"total"`
	USDValue               float64 `json:"usdvalue"`
	StopBorrow             float64 `json:"spotBorrow"`
	AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
}

func (req *RequestForBalances) Path() string {
	return "/wallet/balances"
}

func (req *RequestForBalances) Method() string {
	return http.MethodGet
}

func (req *RequestForBalances) Query() string {
	return ""
}

func (req *RequestForBalances) Payload() []byte {
	return nil
}
