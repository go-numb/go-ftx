package account

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForLeverage struct {
	Leverage int `json:"leverage"`
}

type ResponseForLeverage interface{}

func (req *RequestForLeverage) Path() string {
	return "/account/leverage"
}

func (req *RequestForLeverage) Method() string {
	return http.MethodPost
}

func (req *RequestForLeverage) Query() string {
	return ""
}

func (req *RequestForLeverage) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
