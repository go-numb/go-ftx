package rest_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/0x1be20/go-ftx/auth"
	"github.com/0x1be20/go-ftx/types"

	"github.com/stretchr/testify/assert"

	"github.com/0x1be20/go-ftx/rest"
	"github.com/0x1be20/go-ftx/rest/private/leveraged"
	"github.com/0x1be20/go-ftx/rest/private/options"
)

/*
	# Leveraged tokens
*/
func TestLevTokens(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.LvTokens(&leveraged.RequestForLvTokens{})
	assert.NoError(t, err)

	list := res.Products()
	fmt.Printf("%+v\n", strings.Join(list, "\n"))
}

func TestLvToken(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.LvToken(&leveraged.RequestForLvToken{
		ProductCode: "BULL",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestCreatedLvTokens(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.CreatedLvTokens(&leveraged.RequestForCreatedLvTokens{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestCreatedLvToken(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.CreatedLvToken(&leveraged.RequestForCreatedLvToken{
		ProductCode: "BULL",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestRedemptionLvTokens(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.RedemptionLvTokens(&leveraged.RequestForRedemptionLvTokens{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestRedemptionLvToken(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.RedemptionLvToken(&leveraged.RequestForRedemptionLvToken{
		ProductCode: "BULL",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestLvBalances(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.LvBalances(&leveraged.RequestForLvBalances{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

/*
	# Options
*/

func TestOpQuoteRequests(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.OpQuoteRequests(&options.RequestForOpQuoteRequests{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestMyOpQuoteRequests(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.MyOpQuoteRequests(&options.RequestForMyOpQuoteRequests{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestMyOpQuoteRequest(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.MyOpQuoteRequest(&options.RequestForMyOpQuoteRequest{
		RequestID: 1,
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestCreateOpQuoteRequest(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.CreateOpQuoteRequest(&options.RequestForCreateOpQuoteRequest{
		Underlying: "BTC",
		Type:       "call",
		Strike:     6200,
		Expiry:     time.Now().Add(10 * time.Hour).Unix(),
		Side:       types.BUY,
		Size:       1,
		// Optionals
		// LimitPrice:     6800,
		// HideLimitPrice: true,
		// RequestExpiry:  time.Now().Add(10 * time.Hour).Unix(),
		// CounterpartyID: 1,
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestModifyOpQuoteRequest(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.ModifyOpQuoteRequest(&options.RequestForModifyOpQuoteRequest{
		RequestID: 1,
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestCancelOpQuoteRequest(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.CancelOpQuoteRequest(&options.RequestForCancelOpQuoteRequest{
		RequestID: 1,
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestMyOpQuotes(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.MyOpQuotes(&options.RequestForMyOpQuotes{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

// func TestCreateOpQuote(t *testing.T) {
// 	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

// 	res, err := c.CreateOpQuote(&options.RequestForCreateOpQuote{})
// 	assert.NoError(t, err)

// 	fmt.Printf("%+v\n", res)
// }

func TestCancelOpQuote(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.CancelOpQuote(&options.RequestForCancelOpQuote{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestAcceptOpQuote(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.AcceptOpQuote(&options.RequestForAcceptOpQuote{
		QuoteID: 1,
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestOpPositions(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.OpPositions(&options.RequestForOpPositions{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestOpTrades(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.OpTrades(&options.RequestForOpTrades{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestOpFills(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))

	res, err := c.OpFills(&options.RequestForOpFills{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}
