package realtime_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/boyi/go-ftx/realtime"
)

func TestConnect(t *testing.T) {
	start := time.Now()
	defer func() {
		fmt.Printf("exec time: %f s\n", time.Since(start).Seconds())
	}()

	isReconnect := false

RECONNECT:
	if isReconnect {
		fmt.Println("[ws RE:CONNECT]")
		isReconnect = false
		time.Sleep(15 * time.Second)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan realtime.Response, 10)
	go realtime.Connect(ctx, ch, []string{"trades"}, []string{"BTC-PERP", "ETH-PERP", "SOL-PERP"}, nil)

	for {
		select {
		case v := <-ch:
			switch v.Type {
			case realtime.TICKER:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Ticker)

			case realtime.TRADES:
				// fmt.Printf("%s	%+v\n", v.Symbol, v.Trades)
				for i := range v.Trades {
					if v.Trades[i].Liquidation {
						fmt.Printf("-----------------------------[%s]%+v\n", v.Symbol, v.Trades[i])
					}
				}

			case realtime.ORDERBOOK:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Orderbook)

			case realtime.UNDEFINED:
				fmt.Printf("%s	%s\n", v.Symbol, v.Results.Error())

			case realtime.ERROR:
				fmt.Printf("[ERROR] %s\n", v.Results.Error())
				goto EXIT
			}
		}
	}

EXIT:
	time.AfterFunc(time.Minute, cancel)
	close(ch)
	isReconnect = true

	fmt.Printf("reconnect exec time: %f s\n", time.Since(start).Seconds())

	goto RECONNECT
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
