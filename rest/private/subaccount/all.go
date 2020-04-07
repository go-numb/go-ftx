package subaccount

import (
	"net/http"
)

type RequestForSubAccounts struct {
}

type ResponseForSubAccounts []SubAccount

type SubAccount struct {
	Nickname  string `json:"nickname"`
	Deletable bool   `json:"deletable"`
	Editable  bool   `json:"editable"`
}

func (req *RequestForSubAccounts) Path() string {
	return "/subaccounts"
}

func (req *RequestForSubAccounts) Method() string {
	return http.MethodGet
}

func (req *RequestForSubAccounts) Query() string {
	return ""
}

func (req *RequestForSubAccounts) Payload() []byte {
	return nil
}
