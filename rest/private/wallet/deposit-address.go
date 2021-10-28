package wallet

import (
	"fmt"
	"net/http"
)

type RequestForDepositAddress struct {
	Coin    string `json:"-"`
	Methods string `json:"-"`
}

type ResponseForDepositAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

func (req *RequestForDepositAddress) Path() string {
	path := fmt.Sprintf("/wallet/deposit_address/%s", req.Coin)
	if req.Methods != "" {
		path = fmt.Sprintf("%s?method=%s", path, req.Methods)
	}
	return path
}

func (req *RequestForDepositAddress) Method() string {
	return http.MethodGet
}

func (req *RequestForDepositAddress) Query() string {
	return ""
}

func (req *RequestForDepositAddress) Payload() []byte {
	return nil
}
