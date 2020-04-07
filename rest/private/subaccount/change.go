package subaccount

import (
	"net/http"
)

type RequestForChangeSubAccount struct {
	NickName    string `json:"nickname"`
	NewNickname string `json:"newNickname"`
}

type ResponseForChangeSubAccount string

func (req *RequestForChangeSubAccount) Path() string {
	return "/subaccounts/update_name"
}

func (req *RequestForChangeSubAccount) Method() string {
	return http.MethodPost
}

func (req *RequestForChangeSubAccount) Query() string {
	return ""
}

func (req *RequestForChangeSubAccount) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
