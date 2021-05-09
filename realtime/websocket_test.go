package realtime_test

import (
	"context"
	"fmt"
	"time"
	"testing"

	"github.com/go-numb/go-ftx/realtime"
	// "github.com/tuanito/go-ftx/realtime"
)

func TestConnect(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rtQuotesCh := make(chan realtime.Response)
	// errCh := make (chan error)
	go realtime.Connect(ctx, rtQuotesCh, []string{"ticker"}, []string{"BTC-PERP", "ETH-PERP"}, nil)

	layout := "2006-01-02 15:04:05"
	for {
		currentTime := time.Now()
		select {
		case v := <-rtQuotesCh:
			switch v.Type {
			case realtime.TICKER:
				fmt.Printf("%s %s	%+v\n", currentTime.Format(layout), v.Symbol, v.Ticker)

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
			}
			// Error handling for auto-reconnect
			if v.Results  != nil { 
			fmt.Println("Connection error :", v.Results)
			}
		}
	}

}

func TestConnectForPrivate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan realtime.Response)
	go realtime.ConnectForPrivate(ctx, ch, "", "", []string{"orders", "fills"}, nil)

	for {
		select {
		case v := <-ch:
			switch v.Type {
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
