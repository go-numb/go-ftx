package ftx

import (
	"bytes"
	"net/http"
)

type Account struct {
	BackstopProvider             bool    `json:"backstopProvider"`
	Collateral                   float64 `json:"collateral"`
	FreeCollateral               float64 `json:"freeCollateral"`
	InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
	Liquidating                  bool    `json:"liquidating"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`
	MakerFee                     float64 `json:"makerFee"`
	MarginFraction               float64 `json:"marginFraction"`
	OpenMarginFraction           float64 `json:"openMarginFraction"`
	TakerFee                     float64 `json:"takerFee"`
	TotalAccountValue            float64 `json:"totalAccountValue"`
	TotalPositionSize            float64 `json:"totalPositionSize"`
	Username                     string  `json:"username"`
	Positions                    []struct {
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
}

func (p *Client) Account() (account *Account, err error) {
	res, err := p.sendRequest(
		http.MethodGet,
		"/account",
		nil, nil)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (p *Client) Leverage(n int) (leverage int, err error) {
	params := make(map[string]interface{})
	params["leverage"] = n
	body, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}

	res, err := p.sendRequest(
		http.MethodPost,
		"/account/leverage",
		bytes.NewReader(body), nil)
	if err != nil {
		return 0, err
	}

	// in Close()
	err = decode(res, &leverage)
	if err != nil {
		return 0, err
	}

	return leverage, nil
}

type Balance struct {
	Coin  string  `json:"coin"`
	Free  float64 `json:"free"`
	Total float64 `json:"total"`
}

func (p *Client) Balances() (balances []Balance, err error) {
	res, err := p.sendRequest(
		http.MethodGet,
		"/wallet/balances",
		nil, nil)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &balances)
	if err != nil {
		return nil, err
	}

	return balances, nil
}
