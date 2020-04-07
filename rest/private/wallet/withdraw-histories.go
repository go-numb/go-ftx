package wallet

import (
	"net/http"
	"time"
)

type RequestForWithdrawHistories struct {
}

type ResponseForWithdrawHistories []Withdraw

type Withdraw struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Status  string `json:"status"`
	Txid    string `json:"txid"`

	Fee  float64 `json:"fee"`
	Size float64 `json:"size,string"`

	Time time.Time `json:"time"`

	ID int `json:"id"`
}

func (req *RequestForWithdrawHistories) Path() string {
	return "/wallet/withdrawals"
}

func (req *RequestForWithdrawHistories) Method() string {
	return http.MethodGet
}

func (req *RequestForWithdrawHistories) Query() string {
	return ""
}

func (req *RequestForWithdrawHistories) Payload() []byte {
	return nil
}
