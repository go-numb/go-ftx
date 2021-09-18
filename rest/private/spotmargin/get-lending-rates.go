package spotmargin

import (
	"encoding/json"
	"net/http"
)

type ResponseForLendingRates []LendingRates

type LendingRates struct {
	Coin     string  `json:"coin"`
	Estimate float64 `json:"estimate"`
	Previous float64 `json:"previous"`
}

type RequestForLendingRates struct{}

func (req *RequestForLendingRates) Path() string {
	return "/spot_margin/lending_rates"
}

func (req *RequestForLendingRates) Method() string {
	return http.MethodGet
}

func (req *RequestForLendingRates) Query() string {
	return ""
}

func (req *RequestForLendingRates) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
