package rest

import (
	"github.com/sanychtasher/go-ftx/rest/public/futures"
	"github.com/sanychtasher/go-ftx/rest/public/markets"
)

func (p *Client) Markets(req *markets.RequestForMarkets) (*markets.ResponseForMarkets, error) {
	results := new(markets.ResponseForMarkets)
	if req.ProductCode == "" {
		if err := p.request(req, results); err != nil {
			return nil, err
		}
		return results, nil
	}

	r := new(markets.Market)
	if err := p.request(req, r); err != nil {
		return nil, err
	}

	*results = append(*results, *r)

	return results, nil
}

func (p *Client) Orderbook(req *markets.RequestForOrderbook) (*markets.ResponseForOrderbook, error) {
	results := new(markets.ResponseForOrderbook)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Trades(req *markets.RequestForTrades) (*markets.ResponseForTrades, error) {
	results := new(markets.ResponseForTrades)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Candles(req *markets.RequestForCandles) (*markets.ResponseForCandles, error) {
	results := new(markets.ResponseForCandles)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Futures(req *futures.RequestForFutures) (*futures.ResponseForFutures, error) {
	results := new(futures.ResponseForFutures)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Future(req *futures.RequestForFuture) (*futures.ResponseForFuture, error) {
	results := new(futures.ResponseForFuture)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Stats(req *futures.RequestForStats) (*futures.ResponseForStats, error) {
	results := new(futures.ResponseForStats)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Rates(req *futures.RequestForRates) (*futures.ResponseForRates, error) {
	results := new(futures.ResponseForRates)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
