package account

import (
	"net/http"
)

type RequestForPositions struct {
}

type ResponseForPositions []Position

type Position struct {
	Future string `json:"future"`
	Side   string `json:"side"`

	InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`

	EntryPrice                float64 `json:"entryPrice"`
	EstimatedLiquidationPrice float64 `json:"estimatedLiquidationPrice,omitempty"`

	Size           float64 `json:"size"`
	NetSize        float64 `json:"netSize"`
	OpenSize       float64 `json:"openSize"`
	LongOrderSize  float64 `json:"longOrderSize"`
	ShortOrderSize float64 `json:"shortOrderSize"`

	Cost          float64 `json:"cost"`
	UnrealizedPnl float64 `json:"unrealizedPnl"`
	RealizedPnl   float64 `json:"realizedPnl"`
}

func (req *RequestForPositions) Path() string {
	return "/positions"
}

func (req *RequestForPositions) Method() string {
	return http.MethodGet
}

func (req *RequestForPositions) Query() string {
	return ""
}

func (req *RequestForPositions) Payload() []byte {
	return nil
}
