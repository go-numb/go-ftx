package wallet

import (
	"net/http"
)

type RequestForWithdraw struct {
	Coin    string  `url:"coin"`
	Size    float64 `url:"size"`
	Address string  `url:"address"`
	// Optionals
	Tag      string `url:"tag,omitempty"`
	Password string `url:"password,omitempty"`
	Code     int    `url:"code,omitempty"`
}

type ResponseForWithdraw Withdraw

func (req *RequestForWithdraw) Path() string {
	return "/wallet/withdrawals"
}

func (req *RequestForWithdraw) Method() string {
	return http.MethodPost
}

func (req *RequestForWithdraw) Query() string {
	return ""
}

func (req *RequestForWithdraw) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
