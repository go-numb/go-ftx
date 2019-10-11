package ftx

import (
	"fmt"
	"net/http"
	"time"
)

type Market struct {
	Ask            float64     `json:"ask"`
	BaseCurrency   interface{} `json:"baseCurrency"`
	Bid            float64     `json:"bid"`
	Change1H       float64     `json:"change1h"`
	Change24H      float64     `json:"change24h"`
	ChangeBod      float64     `json:"changeBod"`
	Enabled        bool        `json:"enabled"`
	Last           float64     `json:"last"`
	Name           string      `json:"name"`
	Price          float64     `json:"price"`
	PriceIncrement float64     `json:"priceIncrement"`
	QuoteCurrency  interface{} `json:"quoteCurrency"`
	QuoteVolume24H float64     `json:"quoteVolume24h"`
	SizeIncrement  float64     `json:"sizeIncrement"`
	Type           string      `json:"type"`
	Underlying     string      `json:"underlying"`
	VolumeUsd24H   float64     `json:"volumeUsd24h"`
}

func (p *Client) Markets() (markets []Market, err error) {
	res, err := p.sendRequest(http.MethodGet, "/markets", nil, nil)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}

type Trade struct {
	ID          int       `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        string    `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

/*
Trades is gets executions data

- market_name	string	BTC0628	name of the token

- limit	number	35	optional, max 100, min 20, default 20

- start_time	number	1559881511	optional

- end_time	number	1559881711	optional
*/
func (p *Client) Trades(limit int, tokenName string, start, end time.Time) (trades []Trade, err error) {
	params := make(map[string]string)
	params["limit"] = fmt.Sprintf("%d", limit)
	params["start_time"] = fmt.Sprintf("%d", start.Unix())
	params["end_time"] = fmt.Sprintf("%d", end.Unix())

	res, err := p.sendRequest(
		http.MethodGet,
		fmt.Sprintf(
			"/markets/%s/trades", //&start_time=%d&end_time=%d",
			tokenName),
		nil, &params)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &trades)
	if err != nil {
		return nil, err
	}

	return trades, nil
}
