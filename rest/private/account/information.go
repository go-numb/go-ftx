package account

import (
	"net/http"
)

type RequestForInformation struct {
}

type ResponseForInformation struct {
	Username string `json:"username"`

	Collateral               float64 `json:"collateral"`
	FreeCollateral           float64 `json:"freeCollateral"`
	TotalAccountValue        float64 `json:"totalAccountValue"`
	TotalPositionSize        float64 `json:"totalPositionSize"`
	InitialMarginRequirement float64 `json:"initialMarginRequirement"`
	Leverage                 float64 `json:"leverage"`

	MakerFee                     float64 `json:"makerFee"`
	TakerFee                     float64 `json:"takerFee"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`

	MarginFraction     float64 `json:"marginFraction"`
	OpenMarginFraction float64 `json:"openMarginFraction"`

	Positions []Position `json:"positions"`

	BackstopProvider bool `json:"backstopProvider"`
	Liquidating      bool `json:"liquidating"`
}

func (req *RequestForInformation) Path() string {
	return "/account"
}

func (req *RequestForInformation) Method() string {
	return http.MethodGet
}

func (req *RequestForInformation) Query() string {
	return ""
}

func (req *RequestForInformation) Payload() []byte {
	return nil
}
