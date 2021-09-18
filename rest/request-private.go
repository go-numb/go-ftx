package rest

import (
	"github.com/go-numb/go-ftx/rest/private/account"
	"github.com/go-numb/go-ftx/rest/private/fills"
	"github.com/go-numb/go-ftx/rest/private/funding"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"github.com/go-numb/go-ftx/rest/private/spotmargin"
	"github.com/go-numb/go-ftx/rest/private/subaccount"
	"github.com/go-numb/go-ftx/rest/private/wallet"
)

func (p *Client) Information(req *account.RequestForInformation) (*account.ResponseForInformation, error) {
	results := new(account.ResponseForInformation)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Positions(req *account.RequestForPositions) (*account.ResponseForPositions, error) {
	results := new(account.ResponseForPositions)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Leverage(req *account.RequestForLeverage) (*account.ResponseForLeverage, error) {
	results := new(account.ResponseForLeverage)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Coins(req *wallet.RequestForCoins) (*wallet.ResponseForCoins, error) {
	results := new(wallet.ResponseForCoins)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Balances(req *wallet.RequestForBalances) (*wallet.ResponseForBalances, error) {
	results := new(wallet.ResponseForBalances)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) BalancesAll(req *wallet.RequestForBalancesAll) (*wallet.ResponseForBalancesAll, error) {
	results := new(wallet.ResponseForBalancesAll)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) DepositAddress(req *wallet.RequestForDepositAddress) (*wallet.ResponseForDepositAddress, error) {
	results := new(wallet.ResponseForDepositAddress)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) DepositHistories(req *wallet.RequestForDepositHistories) (*wallet.ResponseForDepositHistories, error) {
	results := new(wallet.ResponseForDepositHistories)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) WithdrawHistories(req *wallet.RequestForWithdrawHistories) (*wallet.ResponseForWithdrawHistories, error) {
	results := new(wallet.ResponseForWithdrawHistories)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Withdraw(req *wallet.RequestForWithdraw) (*wallet.ResponseForWithdraw, error) {
	results := new(wallet.ResponseForWithdraw)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OpenOrder(req *orders.RequestForOpenOrder) (*orders.ResponseForOpenOrder, error) {
	results := new(orders.ResponseForOpenOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OrderHistories(req *orders.RequestForHistories) (*orders.ResponseForHistories, error) {
	results := new(orders.ResponseForHistories)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OpenTriggerOrders(req *orders.RequestForOpenTriggerOrders) (*orders.ResponseForOpenTriggerOrders, error) {
	results := new(orders.ResponseForOpenTriggerOrders)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OrderTriggers(req *orders.RequestForOrderTriggers) (*orders.ResponseForOrderTriggers, error) {
	results := new(orders.ResponseForOrderTriggers)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OrderTriggerHistories(req *orders.RequestForOrderTriggerHistories) (*orders.ResponseForOrderTriggerHistories, error) {
	results := new(orders.ResponseForOrderTriggerHistories)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

/*
	# send order
*/
func (p *Client) PlaceOrder(req *orders.RequestForPlaceOrder) (*orders.ResponseForPlaceOrder, error) {
	results := new(orders.ResponseForPlaceOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) PlaceTriggerOrder(req *orders.RequestForPlaceTriggerOrder) (*orders.ResponseForPlaceTriggerOrder, error) {
	results := new(orders.ResponseForPlaceTriggerOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// ModifyOrder use ClientID or OrderID, prioritize ClientID if has ClientID.
// if has clientID, use by_client_id modify. else if has orderID, use modify.
// the order's queue priority will be reset, and the order ID of the modified order will be different from that of the original order.
func (p *Client) ModifyOrder(req *orders.RequestForModifyOrder) (*orders.ResponseForModifyOrder, error) {
	results := new(orders.ResponseForModifyOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) ModifyTriggerOrder(req *orders.RequestForModifyTriggerOrder) (*orders.ResponseForModifyTriggerOrder, error) {
	results := new(orders.ResponseForModifyTriggerOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OrderStatus(req *orders.RequestForOrderStatus) (*orders.ResponseForOrderStatus, error) {
	results := new(orders.ResponseForOrderStatus)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// CancelByID do cancel with triggerOrderID > clientID > orderID.
func (p *Client) CancelByID(req *orders.RequestForCancelByID) (*orders.ResponseForCancelByID, error) {
	results := new(orders.ResponseForCancelByID)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CancelAll(req *orders.RequestForCancelAll) (*orders.ResponseForCancelAll, error) {
	results := new(orders.ResponseForCancelAll)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Fills(req *fills.Request) (*fills.Response, error) {
	results := new(fills.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Funding(req *funding.Request) (*funding.Response, error) {
	results := new(funding.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

/*
	# Spot Margin
*/

func (p *Client) GetLendingInfo(req *spotmargin.RequestForLendingInfo) (*spotmargin.ResponseForLendingInfo, error) {
	results := new(spotmargin.ResponseForLendingInfo)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetLendingRates(req *spotmargin.RequestForLendingRates) (*spotmargin.ResponseForLendingRates, error) {
	results := new(spotmargin.ResponseForLendingRates)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetLendingHistory(req *spotmargin.RequestForLendingHistory) (*spotmargin.ResponseForLendingHistory, error) {
	results := new(spotmargin.ResponseForLendingHistory)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) SubmitLendingOffer(req *spotmargin.RequestForLendingOffer) (*spotmargin.ResponseForLendingOffer, error) {
	results := new(spotmargin.ResponseForLendingOffer)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetBorrowRates(req *spotmargin.RequestForBorrowRates) (*spotmargin.ResponseForBorrowRates, error) {
	results := new(spotmargin.ResponseForBorrowRates)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) GetBorrowHistory(req *spotmargin.RequestForBorrowHistory) (*spotmargin.ResponseForBorrowHistory, error) {
	results := new(spotmargin.ResponseForBorrowHistory)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

/*
	# SubAccount
*/
func (p *Client) SubAccounts(req *subaccount.RequestForSubAccounts) (*subaccount.ResponseForSubAccounts, error) {
	results := new(subaccount.ResponseForSubAccounts)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) CreateSubAccount(req *subaccount.RequestForCreateSubAccount) (*subaccount.ResponseForCreateSubAccount, error) {
	results := new(subaccount.ResponseForCreateSubAccount)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) ChangeSubAccount(req *subaccount.RequestForChangeSubAccount) (*subaccount.ResponseForChangeSubAccount, error) {
	results := new(subaccount.ResponseForChangeSubAccount)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) DeleteSubAccount(req *subaccount.RequestForDeleteSubAccount) (*subaccount.ResponseForDeleteSubAccount, error) {
	results := new(subaccount.ResponseForDeleteSubAccount)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) BalanceSubAccount(req *subaccount.RequestForBalanceSubAccount) (*subaccount.ResponseForBalanceSubAccount, error) {
	results := new(subaccount.ResponseForBalanceSubAccount)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) TransferSubAccount(req *subaccount.RequestForTransferSubAccount) (*subaccount.ResponseForTransferSubAccount, error) {
	results := new(subaccount.ResponseForTransferSubAccount)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
