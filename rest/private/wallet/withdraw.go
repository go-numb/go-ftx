package wallet

import (
	"net/http"
	"time"
)

type RequestForWithdraw struct {
	Coin    string  `json:"coin"`
	Size    float64 `json:"size"`
	Address string  `json:"address"`
	// Optionals
	Tag      string `json:"tag,omitempty"`
	Methods  string `json:"method,omitempty"`
	Password string `json:"password,omitempty"`
	Code     string `json:"code,omitempty"`
}

type ResponseForWithdraw struct {
	Coin    string    `json:"coin,omitempty"`
	Address string    `json:"address,omitempty"`
	Tag     string    `json:"tag,omitempty"`
	Fee     float64   `json:"fee,omitempty"`
	ID      int64     `json:"id,omitempty"`
	Size    float64   `json:"size,omitempty"`
	Status  string    `json:"status,omitempty"` // one of "requested", "processing", "complete", or "cancelled"
	Time    time.Time `json:"time,omitempty"`
	TxID    string    `json:"txid,omitempty"`
}

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
