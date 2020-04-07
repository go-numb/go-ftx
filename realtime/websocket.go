package realtime

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-numb/go-ftx/rest/private/fills"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/go-numb/go-ftx/types"
	"github.com/gorilla/websocket"
)

const (
	UNDEFINED = iota
	TICKER
	TRADES
	ORDERBOOK
	ORDERS
	FILLS
)

type request struct {
	Op      string `json:"op"`
	Channel string `json:"channel"`
	Market  string `json:"market"`
}

// {"op": "login", "args": {"key": "<api_key>", "sign": "<signature>", "time": 1111}}
type requestForPrivate struct {
	Op   string                 `json:"op"`
	Args map[string]interface{} `json:"args"`
}

type Response struct {
	Type   int
	Symbol string

	Ticker    markets.Ticker
	Trades    []markets.Trade
	Orderbook Orderbook

	Orders orders.Order
	Fills  fills.Fill

	Results error
}

type Orderbook struct {
	Bids [][]float64 `json:"bids"`
	Asks [][]float64 `json:"asks"`
	// Action return update/partial
	Action   string        `json:"action"`
	Time     types.FtxTime `json:"time"`
	Checksum int           `json:"checksum"`
}

func subscribe(conn *websocket.Conn, channels, symbols []string) error {
	if symbols != nil {
		for i := range channels {
			for j := range symbols {
				if err := conn.WriteJSON(&request{
					Op:      "subscribe",
					Channel: channels[i],
					Market:  symbols[j],
				}); err != nil {
					return err
				}
			}
		}
	} else {
		for i := range channels {
			if err := conn.WriteJSON(&request{
				Op:      "subscribe",
				Channel: channels[i],
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func unsubscribe(conn *websocket.Conn, channels, symbols []string) error {
	if symbols != nil {
		for i := range channels {
			for j := range symbols {
				if err := conn.WriteJSON(&request{
					Op:      "unsubscribe",
					Channel: channels[i],
					Market:  symbols[j],
				}); err != nil {
					return err
				}
			}
		}
	} else {
		for i := range channels {
			if err := conn.WriteJSON(&request{
				Op:      "unsubscribe",
				Channel: channels[i],
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func ping(conn *websocket.Conn) (err error) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, []byte(`{"op": "pong"}`)); err != nil {
				goto EXIT
			}
		}
	}
EXIT:
	return err
}

func Connect(ctx context.Context, ch chan Response, channels, symbols []string, l *log.Logger) error {
	if l == nil {
		l = log.New(os.Stdout, "ftx websocket", log.Llongfile)
	}

	conn, _, err := websocket.DefaultDialer.Dial("wss://ftx.com/ws/", nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := subscribe(conn, channels, symbols); err != nil {
		l.Println(err)
		return err
	}
	defer unsubscribe(conn, channels, symbols)

	// ping each 15sec for exchange
	go ping(conn)

RESTART:
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			l.Printf("[ERROR]: msg error: %+v\n", err)
			break RESTART
		}

		var res Response
		channel, err := jsonparser.GetString(msg, "channel")
		if err != nil {
			l.Printf("[ERROR]: market err: %+v", string(msg))
			continue
		}

		market, err := jsonparser.GetString(msg, "market")
		if err != nil {
			l.Printf("[ERROR]: market err: %+v", string(msg))
			continue
		}
		res.Symbol = market

		data, _, _, err := jsonparser.Get(msg, "data")
		if err != nil {
			if isSubscribe, _ := jsonparser.GetString(msg, "type"); isSubscribe == "subscribed" {
				l.Printf("[SUCCESS]: %s %+v", isSubscribe, string(msg))
				continue
			}
			return fmt.Errorf("[ERROR]: data err: %+v %s", err, string(msg))
		}

		switch channel {
		case "ticker":
			res.Type = TICKER
			if err := json.Unmarshal(data, &res.Ticker); err != nil {
				l.Printf("[WARN]: cant unmarshal ticker %+v", err)
				continue
			}

		case "trades":
			res.Type = TRADES
			if err := json.Unmarshal(data, &res.Trades); err != nil {
				l.Printf("[WARN]: cant unmarshal trades %+v", err)
				continue
			}

		case "orderbook":
			res.Type = ORDERBOOK
			if err := json.Unmarshal(data, &res.Orderbook); err != nil {
				l.Printf("[WARN]: cant unmarshal orderbook %+v", err)
				continue
			}

		default:
			res.Type = UNDEFINED
			res.Results = fmt.Errorf("%v", string(msg))
		}

		ch <- res

	}

	return nil
}

func ConnectForPrivate(ctx context.Context, ch chan Response, key, secret string, channels []string, l *log.Logger) error {
	if l == nil {
		l = log.New(os.Stdout, "ftx websocket", log.Llongfile)
	}

	conn, _, err := websocket.DefaultDialer.Dial("wss://ftx.com/ws/", nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// sign up
	if err := signeture(conn, key, secret); err != nil {
		l.Fatal(err)
	}

	if err := subscribe(conn, channels, nil); err != nil {
		l.Println(err)
		return err
	}
	defer unsubscribe(conn, channels, nil)

	go ping(conn)

RESTART:
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			l.Printf("[ERROR]: msg error: %+v", string(msg))
			break RESTART
		}

		var res Response
		channel, err := jsonparser.GetString(msg, "channel")
		if err != nil {
			l.Printf("[ERROR]: market err: %s", string(msg))
			continue
		}

		data, _, _, err := jsonparser.Get(msg, "data")
		if err != nil {
			if isSubscribe, _ := jsonparser.GetString(msg, "type"); isSubscribe == "subscribed" {
				l.Printf("[SUCCESS]: %s %+v", isSubscribe, string(msg))
				continue
			}
			return fmt.Errorf("[ERROR]: data err: %v %s", err, string(msg))
		}

		// Private channel has not market name.
		switch channel {
		case "orders":
			res.Type = ORDERS
			if err := json.Unmarshal(data, &res.Orders); err != nil {
				l.Printf("[WARN]: cant unmarshal orders %+v", err)
				continue
			}

		case "fills":
			res.Type = FILLS
			if err := json.Unmarshal(data, &res.Orders); err != nil {
				l.Printf("[WARN]: cant unmarshal fills %+v", err)
				continue
			}

		default:
			res.Type = UNDEFINED
			res.Results = fmt.Errorf("%v", string(msg))
		}

		ch <- res
	}

	return nil
}

func signeture(conn *websocket.Conn, key, secret string) error {
	// key: your API key
	// time: integer current timestamp (in milliseconds)
	// sign: SHA256 HMAC of the following string, using your API secret: <time>websocket_login
	// subaccount: (optional) subaccount name
	// As an example, if:

	// time: 1557246346499
	// secret: 'Y2QTHI23f23f23jfjas23f23To0RfUwX3H42fvN-'
	// sign would be d10b5a67a1a941ae9463a60b285ae845cdeac1b11edc7da9977bef0228b96de9

	// One websocket connection may be logged in to at most one user. If the connection is already authenticated, further attempts to log in will result in 400s.

	msec := time.Now().UTC().UnixNano() / int64(time.Millisecond)

	mac := hmac.New(sha256.New, []byte(fmt.Sprintf("%d%s", msec, secret)))

	if err := conn.WriteJSON(&requestForPrivate{
		Op: "login",
		Args: map[string]interface{}{
			"key":  key,
			"sign": hex.EncodeToString(mac.Sum(nil)),
			"time": msec,
		},
	}); err != nil {
		return err
	}

	return nil
}
