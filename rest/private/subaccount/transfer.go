package subaccount

import (
	"net/http"
	"time"
)

type RequestForTransferSubAccount struct {
	Coin string  `json:"coin"`
	Size float64 `json:"size"`
	// out account name
	Source string `json:"source"`
	// in account name
	Destination string `json:"destination"`
}

type ResponseForTransferSubAccount struct {
	ID     int    `json:"id"`
	Coin   string `json:"coin"`
	Status string `json:"status"`
	Notes  string `json:"notes"`

	Size float64 `json:"size"`

	Time time.Time `json:"time"`
}

func (req *RequestForTransferSubAccount) Path() string {
	return "/subaccounts/transfer"
}

func (req *RequestForTransferSubAccount) Method() string {
	return http.MethodPost
}

func (req *RequestForTransferSubAccount) Query() string {
	return ""
}

func (req *RequestForTransferSubAccount) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
