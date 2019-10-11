package ftx

import (
	"fmt"
	"net/http"
)

type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

func (p *Client) OrderBooks(debth int, tokenName string) (orderbook *Orderbook, err error) {
	params := make(map[string]string)
	params["debth"] = fmt.Sprintf("%d", debth)

	res, err := p.sendRequest(
		http.MethodGet,
		fmt.Sprintf(
			"/markets/%s/orderbook", //&start_time=%d&end_time=%d",
			tokenName),
		nil, &params)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &orderbook)
	if err != nil {
		return nil, err
	}

	return orderbook, nil
}
