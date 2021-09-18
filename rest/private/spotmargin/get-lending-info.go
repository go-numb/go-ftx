package spotmargin

import (
	"encoding/json"
	"net/http"
)

type ResponseForLendingInfo []LendingInfo

type LendingInfo struct {
	Coin     string  `json:"coin"`
	Lendable float64 `json:"lendable"`
	Locked   float64 `json:"locked"`
	MinRate  float64 `json:"minRate"`
	Offered  float64 `json:"offered"`
}

type RequestForLendingInfo struct{}

func (req *RequestForLendingInfo) Path() string {
	return "/spot_margin/lending_info"
}

func (req *RequestForLendingInfo) Method() string {
	return http.MethodGet
}

func (req *RequestForLendingInfo) Query() string {
	return ""
}

func (req *RequestForLendingInfo) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
