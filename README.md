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

**Not Supported**
- [ ] FIX API

## Installation

```
$ go get -u github.com/go-numb/go-ftx
```

## Usage
``` golang
package main

import (
 "fmt"
 "github.com/go-numb/go-ftx/rest"
 "github.com/go-numb/go-ftx/auth"
 "github.com/go-numb/go-ftx/types"
 "github.com/go-numb/go-ftx/private/account"

 "github.com/labstack/gommon/log"
)


func main() {
	client := rest.New(auth.New("<key>", "<secret>"))

	// account informations
	info, err := c.Information(&account.RequestForInformation{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", info)

	lev, err := client.Leverage(5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", lev)
	
	market, err := c.Markets(&markets.RequestForMarkets{})
	if err != nil {
		log.Fatal(err)
	}

	// products List
	fmt.Printf("%+v\n", strings.Join(res.List(), "\n"))
	// product ranking by USD
	fmt.Printf("%+v\n", strings.Join(res.Ranking(markets.ALL), "\n"))


	// FundingRate
	rates, err := c.Rates(&futures.RequestForRates{})
	if err != nil {
		log.Fatal(err)
	}
	// Sort by FundingRate & Print
	// Custom sort
	sort.Sort(sort.Reverse(res))
	for _, v := range *res {
		fmt.Printf("%s			%s		%s\n", humanize.Commaf(v.Rate), v.Future, v.Time.String())
	}

	o, err := c.PlaceOrder(&orders.RequestForPlaceOrder{
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
		client.Logger.Error(err)
	}
    

	ok, err := client.Cancel(&orders.RequestForCancelByID{
		OrderID: "erafa",
		// either... , prioritize clientID
		ClientID: "",
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
	"github.com/go-numb/go-ftx/realtime"
	"github.com/go-numb/go-ftx/auth"

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

			case realtime.UNDEFINED:
				fmt.Printf("%s	%s\n", v.Symbol, v.Results.Error())

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

[MIT](https://github.com/go-numb/go-ftx/blob/master/LICENSE)