package rest

import (
	"github.com/sanychtasher/go-ftx/rest/private/leveraged"
	"github.com/sanychtasher/go-ftx/rest/private/options"
)

/*
	# Leveraged Tokens
*/

func (p *Client) LvTokens(req *leveraged.RequestForLvTokens) (*leveraged.ResponseForLvTokens, error) {
	results := new(leveraged.ResponseForLvTokens)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) LvToken(req *leveraged.RequestForLvToken) (*leveraged.ResponseForLvToken, error) {
	results := new(leveraged.ResponseForLvToken)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CreatedLvTokens(req *leveraged.RequestForCreatedLvTokens) (*leveraged.ResponseForCreatedLvTokens, error) {
	results := new(leveraged.ResponseForCreatedLvTokens)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CreatedLvToken(req *leveraged.RequestForCreatedLvToken) (*leveraged.ResponseForCreatedLvToken, error) {
	results := new(leveraged.ResponseForCreatedLvToken)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) RedemptionLvTokens(req *leveraged.RequestForRedemptionLvTokens) (*leveraged.ResponseForRedemptionLvTokens, error) {
	results := new(leveraged.ResponseForRedemptionLvTokens)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) RedemptionLvToken(req *leveraged.RequestForRedemptionLvToken) (*leveraged.ResponseForRedemptionLvToken, error) {
	results := new(leveraged.ResponseForRedemptionLvToken)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) LvBalances(req *leveraged.RequestForLvBalances) (*leveraged.ResponseForLvBalances, error) {
	results := new(leveraged.ResponseForLvBalances)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

/*
	# Opstions
*/

func (p *Client) OpQuoteRequests(req *options.RequestForOpQuoteRequests) (*options.ResponseForOpQuoteRequests, error) {
	results := new(options.ResponseForOpQuoteRequests)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) MyOpQuoteRequests(req *options.RequestForMyOpQuoteRequests) (*options.ResponseForMyOpQuoteRequests, error) {
	results := new(options.ResponseForMyOpQuoteRequests)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) MyOpQuoteRequest(req *options.RequestForMyOpQuoteRequest) (*options.ResponseForMyOpQuoteRequest, error) {
	results := new(options.ResponseForMyOpQuoteRequest)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CreateOpQuoteRequest(req *options.RequestForCreateOpQuoteRequest) (*options.ResponseForCreateOpQuoteRequest, error) {
	results := new(options.ResponseForCreateOpQuoteRequest)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) ModifyOpQuoteRequest(req *options.RequestForModifyOpQuoteRequest) (*options.ResponseForModifyOpQuoteRequest, error) {
	results := new(options.ResponseForModifyOpQuoteRequest)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CancelOpQuoteRequest(req *options.RequestForCancelOpQuoteRequest) (*options.ResponseForCancelOpQuoteRequest, error) {
	results := new(options.ResponseForCancelOpQuoteRequest)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) MyOpQuotes(req *options.RequestForMyOpQuotes) (*options.ResponseForMyOpQuotes, error) {
	results := new(options.ResponseForMyOpQuotes)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// // Create Quote is Modify Quote. document is incorrect?
// func (p *Client) CreateOpQuote(req *options.RequestForCreateOpQuote) (*options.ResponseForCreateOpQuote, error) {
// 	results := new(options.ResponseForCreateOpQuote)
// 	if err := p.request(req, results); err != nil {
// 		return nil, err
// 	}
// 	return results, nil
// }

func (p *Client) CancelOpQuote(req *options.RequestForCancelOpQuote) (*options.ResponseForCancelOpQuote, error) {
	results := new(options.ResponseForCancelOpQuote)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) AcceptOpQuote(req *options.RequestForAcceptOpQuote) (*options.ResponseForAcceptOpQuote, error) {
	results := new(options.ResponseForAcceptOpQuote)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OpPositions(req *options.RequestForOpPositions) (*options.ResponseForOpPositions, error) {
	results := new(options.ResponseForOpPositions)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OpTrades(req *options.RequestForOpTrades) (*options.ResponseForOpTrades, error) {
	results := new(options.ResponseForOpTrades)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OpFills(req *options.RequestForOpFills) (*options.ResponseForOpFills, error) {
	results := new(options.ResponseForOpFills)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
