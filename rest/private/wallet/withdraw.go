package wallet

import (
	"net/http"
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
	Coin    string    `json:"coin"`
	Address string    `json:"address"`
	Tag     string    `json:"tag"`
	Fee     float64   `json:"fee"`
	ID      int64     `json:"id"`
	Size    float64   `json:"size"`
	Status  string    `json:"status"` // one of "requested", "processing", "complete", or "cancelled"
	Time    time.Time `json:"time"`
	TxID    string    `json:"txid"`
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
