package account

import (
	"net/http"
)

type RequestForPositions struct {
	ShowAvgPrice bool `json:"showAvgPrice"`
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

	CollateralUsed         float64 `json:"collateralUsed,omitempty"`
	RecentAverageOpenPrice float64 `json:"recentAverageOpenPrice,omitempty"`
	RecentPnl              float64 `json:"recentPnl,omitempty"`
	RecentBreakEvenPrice   float64 `json:"recentBreakEvenPrice,omitempty"`
	CumulativeBuySize      float64 `json:"cumulativeBuySize,omitempty"`
	CumulativeSellSize     float64 `json:"cumulativeSellSize,omitempty"`
}

func (req *RequestForPositions) Path() string {
	return "/positions"
}

func (req *RequestForPositions) Method() string {
	return http.MethodGet
}

func (req *RequestForPositions) Query() string {
	if req.ShowAvgPrice {
		return "showAvgPrice=true"
	}
	return ""
}

func (req *RequestForPositions) Payload() []byte {
	return nil
}
