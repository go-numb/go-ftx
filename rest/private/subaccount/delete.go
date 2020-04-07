package subaccount

import (
	"net/http"
)

type RequestForDeleteSubAccount struct {
	NickName string `json:"nickname"`
}

type ResponseForDeleteSubAccount string

func (req *RequestForDeleteSubAccount) Path() string {
	return "/subaccounts"
}

func (req *RequestForDeleteSubAccount) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteSubAccount) Query() string {
	return ""
}

func (req *RequestForDeleteSubAccount) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
