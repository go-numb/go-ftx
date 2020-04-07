package wallet

import (
	"fmt"
	"net/http"
)

type RequestForDepositAddress struct {
	Coin string `json:"-"`
}

type ResponseForDepositAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

func (req *RequestForDepositAddress) Path() string {
	return fmt.Sprintf("/wallet/deposit_address/%s", req.Coin)
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
