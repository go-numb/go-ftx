package ftx

import (
	"net/http"
	"time"
)

type Future struct {
	Future string    `json:"future"`
	Rate   float64   `json:"rate"`
	Time   time.Time `json:"time"`
}

func (p *Client) Futures() (futures []Future, err error) {
	res, err := p.sendRequest(http.MethodGet, "/funding_rates", nil, nil)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &futures)
	if err != nil {
		return nil, err
	}

	return futures, nil
}
