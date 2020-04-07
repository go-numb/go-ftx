package subaccount

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForCreateSubAccount struct {
	NickName string `json:"nickname"`
}

type ResponseForCreateSubAccount SubAccount

func (req *RequestForCreateSubAccount) Path() string {
	return "/subaccounts"
}

func (req *RequestForCreateSubAccount) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateSubAccount) Query() string {
	return ""
}

func (req *RequestForCreateSubAccount) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
