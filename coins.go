package ftx

import (
	"net/http"
)

type Coin struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Collateral bool   `json:"collateral"`
}

func (p *Client) Coins() (coins []Coin, err error) {
	res, err := p.sendRequest(http.MethodGet, "/coins", nil, nil)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}
