package convert_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/boyi/go-ftx/auth"
	"github.com/boyi/go-ftx/rest"
	"github.com/boyi/go-ftx/rest/private/account"
	"github.com/boyi/go-ftx/rest/private/convert"
	"github.com/stretchr/testify/assert"
)

func getAuth() *rest.Client {
	subaccountNickname := os.Getenv("FTXSUBACCOUNT")

	if subaccountNickname == "" {
		return rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET")))
	}

	c := rest.New(auth.New(os.Getenv("FTXKEY"), os.Getenv("FTXSECRET"),
		auth.SubAccount{
			UUID:     1,
			Nickname: subaccountNickname,
		}))
	c.Auth.UseSubAccountID(1)
	return c
}

func TestInformation(t *testing.T) {
	c := getAuth()

	res, err := c.Information(&account.RequestForInformation{})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func GetQuote() (resp *convert.ResponseForRequestQuote, err error) {
	c := getAuth()

	res, err := c.RequestConvertQuote(&convert.RequestForRequestQuote{
		FromCoin: "USDT",
		ToCoin:   "USD",
		Size:     2,
	})
	return res, err
}

func TestRequestQuote(t *testing.T) {
	quote, err := GetQuote()
	assert.NoError(t, err)
	assert.NotZero(t, quote.QuoteId)
	quoteId := quote.QuoteId

	fmt.Printf("QuoteID: %+v\n", quoteId)
}

func GetQuoteStatus(quoteId int) (resp *convert.ResponseForQuoteStatus, err error) {
	c := getAuth()

	res, err := c.GetConvertQuoteStatus(&convert.RequestForQuoteStatus{
		QuoteId: quoteId,
	})

	return res, err
}

func TestQuoteStatus(t *testing.T) {
	quote, err := GetQuote()
	assert.NoError(t, err)
	assert.NotZero(t, quote.QuoteId)
	quoteId := quote.QuoteId
	fmt.Printf("Got quote with ID: %+v\n", quoteId)

	res, err := GetQuoteStatus(quoteId)

	assert.NoError(t, err)
	assert.Equal(t, res.QuoteId, quoteId)
	assert.NotZero(t, res.Price)

	fmt.Printf("Result: %+v\n", res)
}

func AcceptQuote(quoteId int) (resp *convert.ResponseForAcceptQuote, err error) {
	c := getAuth()

	res, err := c.AcceptConvertQuote(&convert.RequestForAcceptQuote{
		QuoteId: quoteId,
	})

	return res, err
}

func TestAcceptQuote(t *testing.T) {
	quote, err := GetQuote()
	assert.NoError(t, err)
	assert.NotZero(t, quote.QuoteId)
	quoteId := quote.QuoteId
	fmt.Printf("Got quote with ID: %+v\n", quoteId)

	res_status, err := GetQuoteStatus(quoteId)
	assert.NoError(t, err)
	assert.Equal(t, res_status.QuoteId, quoteId)
	assert.NotZero(t, res_status.Cost)
	fmt.Printf("Got quote status: %+v\n", res_status)

	res_accept, err := AcceptQuote(quoteId)
	assert.NoError(t, err)
	fmt.Printf("Accepted quote: %+v\n", res_accept)
}
