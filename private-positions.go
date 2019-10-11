package ftx

import "net/http"

type Position struct {
	Cost                         float64 `json:"cost"`
	EntryPrice                   float64 `json:"entryPrice"`
	Future                       string  `json:"future"`
	InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
	LongOrderSize                float64 `json:"longOrderSize"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`
	NetSize                      float64 `json:"netSize"`
	OpenSize                     float64 `json:"openSize"`
	RealizedPnl                  float64 `json:"realizedPnl"`
	ShortOrderSize               float64 `json:"shortOrderSize"`
	Side                         string  `json:"side"`
	Size                         float64 `json:"size"`
	UnrealizedPnl                int     `json:"unrealizedPnl"`
}

func (p *Client) Positions() (positions []Position, err error) {
	res, err := p.sendRequest(
		http.MethodGet,
		"/positions",
		nil, nil)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &positions)
	if err != nil {
		return nil, err
	}

	return positions, nil
}
