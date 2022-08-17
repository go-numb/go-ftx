package rest_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/boyi/go-ftx/rest"
	"github.com/boyi/go-ftx/rest/public/futures"
	"github.com/boyi/go-ftx/rest/public/markets"
	"github.com/stretchr/testify/assert"
)

/*
	# Markets
*/
func TestMarkets(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Markets(&markets.RequestForMarkets{
		ProductCode: "MNGO/USD",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", strings.Join(res.List(), "\n"))
	fmt.Printf("%+v\n", strings.Join(res.Ranking(markets.ALL), "\n"))
}

func TestOrderbook(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Orderbook(&markets.RequestForOrderbook{
		ProductCode: "BTC-PERP",
		Depth:       10,
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestTrades(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Trades(&markets.RequestForTrades{
		ProductCode: "BTC-PERP",
		Limit:       100,
		Start:       time.Now().UTC().Add(-1 * time.Minute).Unix(),
		End:         time.Now().UTC().Unix(),
	})
	assert.NoError(t, err)

	trades := *res
	for i := 0; i < len(trades); i++ {
		fmt.Printf("%#v\n", trades[i].Time.String())
	}
}

func TestCandles(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Candles(&markets.RequestForCandles{
		ProductCode: "BTC-PERP",
		Resolution:  900,                                       // optional
		Start:       time.Now().Add(-900 * time.Second).Unix(), // optional
		End:         time.Now().Unix(),                         // optional
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

/*
	# Futures
*/
func TestFutures(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Futures(&futures.RequestForFutures{})
	assert.NoError(t, err)

	for i, v := range res.Products() {
		fmt.Printf("%d: %+v\n", i, v)
	}

	for _, v := range *res {
		fmt.Printf("%+v	%.2f\n", v.Name, v.Ask)
	}
}

func TestFuture(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Future(&futures.RequestForFuture{
		ProductCode: "BTC-PERP",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestStats(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Stats(&futures.RequestForStats{
		ProductCode: "BTC-PERP",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v\n", res)
}

func TestGetFundingRate(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Futures(&futures.RequestForFutures{})
	assert.NoError(t, err)

	list := res.Products()
	futureList := make([]futures.Stats, len(list))

	for i := range list {
		future, err := c.Stats(&futures.RequestForStats{
			ProductCode: list[i],
		})
		if err != nil {
			t.Fatal(err)
		}
		future.Name = list[i]
		futureList[i] = futures.Stats(*future)

		time.Sleep(100 * time.Millisecond)
	}

	sort.Sort(sort.Reverse(futures.StatsList(futureList)))

	for _, v := range futureList {
		if v.NextFundingRate == 0 {
			continue
		}
		fmt.Printf("%s			%s\n", humanize.Commaf(v.NextFundingRate*100), v.Name)
	}

	/*
		<<
		0.000368		BTMX-PERP
		0.000142		MATIC-PERP
		0.000141		TOMO-PERP
		0.000074		PAXG-PERP
		0.000061		BSV-PERP
		0.000059		PRIV-PERP
		0.000042		TRYB-PERP
		0.000039		EOS-PERP
		0.000034		MID-PERP
		0.00003			OKB-PERP
		0.000025		DRGN-PERP
		0.000024		SHIT-PERP
		0.000023		HT-PERP
		0.000021		LTC-PERP
		0.00002			USDT-PERP
		0.000019		XTZ-PERP
		0.000018		XAUT-PERP
		0.000007		DOGE-PERP
		0.000005		BTC-PERP
		0.000002		BNB-PERP
		0.000001		ADA-PERP
		-0.000007		ATOM-PERP
		-0.000018		BCH-PERP
		-0.000019		ALGO-PERP
		-0.000025		TRX-PERP
		-0.00003		XRP-PERP
		-0.000031		EXCH-PERP
		-0.000034		ALT-PERP
		-0.00004		ETC-PERP
		-0.000044		ETH-PERP
		-0.000096		LINK-PERP
		-0.000162		LEO-PERP
	*/
}

func TestMultiUnderlyingsRates(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Rates(&futures.RequestForRates{})
	assert.NoError(t, err)

	sort.Sort(sort.Reverse(res))
	for _, v := range *res {
		fmt.Printf("%s			%s		%s\n", humanize.Commaf(v.Rate), v.Future, v.Time.String())
	}

	/*
		>>
		0.000404			BTMX-PERP		2020-04-05 10:00:00 +0000 +0000
		0.000367			BTMX-PERP		2020-04-05 11:00:00 +0000 +0000
		0.000357			BTMX-PERP		2020-04-05 07:00:00 +0000 +0000
		0.000338			BTMX-PERP		2020-04-05 09:00:00 +0000 +0000
		0.00029				BTMX-PERP		2020-04-05 08:00:00 +0000 +0000

		(omitting...)

		-0.000277			BTMX-PERP		2020-04-04 23:00:00 +0000 +0000
		-0.000293			DOGE-PERP		2020-04-05 03:00:00 +0000 +0000
		-0.000304			XAUT-PERP		2020-04-04 23:00:00 +0000 +0000
		-0.000322			DOGE-PERP		2020-04-05 04:00:00 +0000 +0000
		-0.000379			BTMX-PERP		2020-04-05 00:00:00 +0000 +0000
		-0.000398			XAUT-PERP		2020-04-05 00:00:00 +0000 +0000
	*/
}

func TestSingleUnderlyingRates(t *testing.T) {
	c := rest.New(nil)
	res, err := c.Rates(&futures.RequestForRates{
		ProductCode: "BTC-PERP",
		Start:       time.Now().Add(-900 * time.Second).Unix(), // optional
		End:         time.Now().Unix(),                         // optional

	})
	assert.NoError(t, err)

	sort.Sort(sort.Reverse(res))
	for _, v := range *res {
		fmt.Printf("%s			%s		%s\n", humanize.Commaf(v.Rate), v.Future, v.Time.String())
	}

	/*
		>>
		0.000404			BTC-PERP		2020-04-05 10:00:00 +0000 +0000
		0.000367			BTC-PERP		2020-04-05 11:00:00 +0000 +0000
		0.000357			BTC-PERP		2020-04-05 07:00:00 +0000 +0000
		0.000338			BTC-PERP		2020-04-05 09:00:00 +0000 +0000
		0.00029				BTC-PERP		2020-04-05 08:00:00 +0000 +0000

		(omitting...)

		-0.000277			BTC-PERP		2020-04-04 23:00:00 +0000 +0000
		-0.000293			BTC-PERP		2020-04-05 03:00:00 +0000 +0000
		-0.000304			BTC-PERP		2020-04-04 23:00:00 +0000 +0000
		-0.000322			BTC-PERP		2020-04-05 04:00:00 +0000 +0000
		-0.000379			BTC-PERP		2020-04-05 00:00:00 +0000 +0000
		-0.000398			BTC-PERP		2020-04-05 00:00:00 +0000 +0000
	*/
}
