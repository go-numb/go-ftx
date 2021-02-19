package spotmargin

import (
	"encoding/json"
	"net/http"
)

type ResponseForLendingOffer struct{}

type RequestForLendingOffer struct {
	Coin string  `json:"coin"`
	Size float64 `json:"size"`
	Rate float64 `json:"rate"`
}

func (req *RequestForLendingOffer) Path() string {
	return "/spot_margin/offers"
}

func (req *RequestForLendingOffer) Method() string {
	return http.MethodPost
}

func (req *RequestForLendingOffer) Query() string {
	return ""
}

func (req *RequestForLendingOffer) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
