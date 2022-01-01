package spotmargin

import (
	"net/http"
)

type RequestForLendingRates struct{}

type ResponseForLendingRates []LendingRates

type LendingRates struct {
	Coin     string  `json:"coin"`
	Estimate float64 `json:"estimate"`
	Previous float64 `json:"previous"`
}

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
	return nil
}
