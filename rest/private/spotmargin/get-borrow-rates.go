package spotmargin

import (
	"encoding/json"
	"net/http"
)

type ResponseForBorrowRates []BorrowRates

type BorrowRates struct {
	Coin     string  `json:"coin"`
	Estimate float64 `json:"estimate"`
	Previous float64 `json:"previous"`
}

type RequestForBorrowRates struct{}

func (req *RequestForBorrowRates) Path() string {
	return "/spot_margin/borrow_rates"
}

func (req *RequestForBorrowRates) Method() string {
	return http.MethodGet
}

func (req *RequestForBorrowRates) Query() string {
	return ""
}

func (req *RequestForBorrowRates) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
