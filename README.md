# go-ftx

FTX exchange API version2, renew at 2020/04.

## Description

go-ftx is a go client library for [FTX API Document](https://docs.ftx.com).

**Supported**
- [x] Public & Private
- [x] Orders
- [x] Leveraged tokens
- [x] Options
- [x] Websocket
- [x] SubAccounts 

**Not Supported**
- [ ] FIX API

## Installation

```
$ go get -u github.com/sanychtasher/go-ftx
```

## Usage
``` golang
package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/sanychtasher/go-ftx/auth"
	"github.com/sanychtasher/go-ftx/rest"
	"github.com/sanychtasher/go-ftx/rest/private/orders"
	//"log"
	"github.com/sanychtasher/go-ftx/rest/private/account"
	"github.com/sanychtasher/go-ftx/rest/public/futures"
	"github.com/sanychtasher/go-ftx/rest/public/markets"
	"github.com/sanychtasher/go-ftx/types"
	"github.com/labstack/gommon/log"
	"sort"
	"strings"
)

func main() {
	// Only main account
	client := rest.New(auth.New("<key>", "<secret>"))

	// or
	// UseSubAccounts
	clientWithSubAccounts := rest.New(
		auth.New(
			"<key>",
			"<secret>",
			auth.SubAccount{
				UUID:     1,
				Nickname: "subaccount_1",
			},
			auth.SubAccount{
				UUID:     2,
				Nickname: "subaccount_2",
			},
			// many....
		))
	// switch subaccount
	clientWithSubAccounts.Auth.UseSubAccountID(1) // or 2... this number is key in map[int]SubAccount

	// account informations
	// client or clientWithSubAccounts in this time.
	c := client // or clientWithSubAccounts
	info, err := c.Information(&account.RequestForInformation{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", info)

	// lev, err := c.Leverage(5)
	lev, err := c.Leverage(&account.RequestForLeverage{
		Leverage: 3,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", lev)

	market, err := c.Markets(&markets.RequestForMarkets{
		ProductCode: "XRPBULL/USDT",
	})

	if err != nil {
		log.Fatal(err)
	}

	// products List
	fmt.Printf("%+v\n", strings.Join(market.List(), "\n"))
	// product ranking by USD
	fmt.Printf("%+v\n", strings.Join(market.Ranking(markets.ALL), "\n"))

	// FundingRate
	rates, err := c.Rates(&futures.RequestForRates{})
	if err != nil {
		log.Fatal(err)
	}
	// Sort by FundingRate & Print
	// Custom sort
	sort.Sort(sort.Reverse(rates))
	for _, v := range *rates {
		fmt.Printf("%s			%s		%s\n", humanize.Commaf(v.Rate), v.Future, v.Time.String())
	}

	order, err := c.PlaceOrder(&orders.RequestForPlaceOrder{
		Type:   types.LIMIT,
		Market: "BTC-PERP",
		Side:   types.BUY,
		Price:  6200,
		Size:   0.01,
		// Optionals
		ClientID:   "use_original_client_id",
		Ioc:        false,
		ReduceOnly: false,
		PostOnly:   false,
	})
	if err != nil {
		// client.Logger.Error(err) // Logger does not seem to exist @tuanito
	}

	fmt.Printf("%+v\n", order)

	ok, err := c.CancelByID(&orders.RequestForCancelByID{
		OrderID: 1,
		// either... , prioritize clientID
		ClientID:       "",
		TriggerOrderID: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ok)
	// ok is status comment

}

```


## Websocket
``` golang 
package main

import (
	"context"
	"fmt"
	"github.com/sanychtasher/go-ftx/realtime"
	"github.com/sanychtasher/go-ftx/auth"

	"github.com/labstack/gommon/log"
)


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan realtime.Response)
	go realtime.Connect(ctx, ch, []string{"ticker"}, []string{"BTC-PERP", "ETH-PERP"}, nil)
	go realtime.ConnectForPrivate(ctx, ch, "<key>", "<secret>", []string{"orders", "fills"}, nil)

	for {
		select {
		case v := <-ch:
			switch v.Type {
			case realtime.TICKER:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Ticker)

			case realtime.TRADES:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Trades)
				for i := range v.Trades {
					if v.Trades[i].Liquidation {
						fmt.Printf("-----------------------------%+v\n", v.Trades[i])
					}
				}

			case realtime.ORDERBOOK:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Orderbook)

			case realtime.ORDERS:
				fmt.Printf("%d	%+v\n", v.Type, v.Orders)

			case realtime.FILLS:
				fmt.Printf("%d	%+v\n", v.Type, v.Fills)

			case realtime.UNDEFINED:
				fmt.Printf("UNDEFINED %s	%s\n", v.Symbol, v.Results.Error())
			}
		}
	}
}
```


## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/sanychtasher/go-ftx/blob/master/LICENSE)
