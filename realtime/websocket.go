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
	ERROR
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

func subscribe(conn *websocket.Conn, channels []string, symbols []string) error {
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

func unsubscribe(conn *websocket.Conn, channels []string, symbols []string) error {
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

/*
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("artificial errror at %v, %s",
	e.When, e.What)
}

func createFakeError() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
*/

func ConnectWithRetries(...., numRetries)

// Maybe should return the connection so we can auto reconnect later on
// Or alternatively extract the conn out of this function
// and pass it as a parameter to Connect.
func Connect(ctx context.Context, ch chan Response, channels []string, symbols []string, l *log.Logger ) error {
	var outputErr error
	outputErr = nil

/*
	snapshottedTime := time.Now()
	var delay float64
	var artificialError error
	artificialError = run() 
	outputErr = artificialError
*/
	if l == nil {
		l = log.New(os.Stdout, "ftx websocket", log.Llongfile)
	}

	// outer loop 1..numRErreies

	conn, _, err := websocket.DefaultDialer.Dial("wss://ftx.com/ws/", nil)
	if err != nil {
		outputErr = err
		return outputErr 
	}

	if err := subscribe(conn, channels, symbols); err != nil {
		outputErr = err
		return outputErr 
	}

	// ping each 15sec for exchange
	go ping(conn)

	go func() {
		defer conn.Close()
		defer unsubscribe(conn, channels, symbols)

	RESTART:
		for {
			var res Response
			_, msg, err := conn.ReadMessage()
			if err != nil {
				l.Printf("[ERROR]: msg error: %+v", err)
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", err)
				ch <- res
				break RESTART
			}

			typeMsg, err := jsonparser.GetString(msg, "type")
			if typeMsg == "error" {
				l.Printf("[ERROR]: error: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			channel, err := jsonparser.GetString(msg, "channel")
			if err != nil {
				l.Printf("[ERROR]: channel error: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			market, err := jsonparser.GetString(msg, "market")
			if err != nil {
				l.Printf("[ERROR]: market err: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			res.Symbol = market

			data, _, _, err := jsonparser.Get(msg, "data")
			if err != nil {
				if isSubscribe, _ := jsonparser.GetString(msg, "type"); isSubscribe == "subscribed" {
					l.Printf("[SUCCESS]: %s %+v", isSubscribe, string(msg))
					continue
				} else {
					err = fmt.Errorf("[ERROR]: data err: %v %s", err, string(msg))
					l.Println(err)
					res.Type = ERROR
					res.Results = err
					ch <- res
					break RESTART
				}
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
			// Create artificial error after a certain time
/*
			currentTime := time.Now()
			delay:= currentTime.Sub(snapshottedTime ).Seconds()
			fmt.Println(" delay = ", delay)
			if  delay > 10.0 {
				break RESTART
			}
*/
		}
	}()

/*

	if  delay > 10.0 {
		fmt.Println("Creating artificial error ! : delay = ", delay)
		artificialError = createFakeError() 
		outputErr = artificialError
	}
*/
	// There was a fatal error on the websocket, let's close the connection gracefully before reconnecting later @Tuan
	if outputErr != nil {
		fmt.Println("Closing connection...")
		conn.Close()
		fmt.Println("Unsubscribing connection...")
		unsubscribe(conn, channels, symbols)
		// Check for unsubscribed message before gracefully exiting : we need an infinite loop that listens to "unsubscribed" incoming
		// messages. Once all have arived, then exit this function with the right error message.

		// Check for unsubscribed according to this protocol : 
		// Excerpts from FTX API documentation : 
		// 	Websocket connections go through the following lifecycle: - Establish a websocket connection with wss://ftx.com/ws/ - (Optional) 
		// 1. Authenticate with {'op': 'login', 'args': {'key': <api_key>, 'sign': <signature>, 'time': <ts>}} - 
		// 2. Send pings at regular intervals (every 15 seconds): {'op': 'ping'}. You will see an {'type': 'pong'} response. - 
		// 3. Subscribe to a channel with {'op': 'subscribe', 'channel': 'trades', 'market': 'BTC-PERP'} - 
		// 4. Receive subscription response {'type': 'subscribed', 'channel': 'trades', 'market': 'BTC-PERP'} - 
		// 5. Receive data {'type': 'update', 'channel': 'trades', 'market': 'BTC-PERP', 'data': {'bid': 5230.5, 'ask': 5231.0, 'ts': 1557133490.4047449, 'last': 5230.5}} - 
		// 6. Unsubscribe {'op': 'unsubscribe', 'channel': 'trades', 'market': 'BTC-PERP'} - 
		// 7. Receive unsubscription response {'type': 'unsubscribed', 'channel': 'trades', 'market': 'BTC-PERP'}
		fmt.Println("Closed connection...")
		fmt.Println("Unsubscribed connection...")
	}
	return outputErr // <-- Now outputErr is not systematically nil, therefore it is possible to handle the error and create an auto-reconnect if there is an issue
	// return nil // Original return
}

// CheckUnsubscribed check wether the feedhandler is unsubscribed before a reconnect
/*
func CheckUnsubscribed ( ) {

			data, _, _, err := jsonparser.Get(msg, "data")
			if err != nil {
				if isUnsubscribed, _ := jsonparser.GetString(msg, "type"); isUnsubscribed == "unsubscribed" {
					l.Printf("[SUCCESS]: %s %+v", isUnsubscribe, string(msg))
					continue
				} else {
					err = fmt.Errorf("[ERROR]: data err: %v %s", err, string(msg))
					l.Println(err)
					res.Type = ERROR
					res.Results = err
					ch <- res
					// break RESTART
				}
}
*/
func ConnectForPrivate(ctx context.Context, ch chan Response, key, secret string, channels []string, l *log.Logger, subaccount ...string) error {
	if l == nil {
		l = log.New(os.Stdout, "ftx websocket", log.Llongfile)
	}

	conn, _, err := websocket.DefaultDialer.Dial("wss://ftx.com/ws/", nil)
	if err != nil {
		return err
	}

	// sign up
	if err := signature(conn, key, secret, subaccount); err != nil {
		return err
	}

	if err := subscribe(conn, channels, nil); err != nil {
		return err
	}

	go ping(conn)

	go func() {
		defer conn.Close()
		defer unsubscribe(conn, channels, nil)

	RESTART:
		for {
			var res Response
			_, msg, err := conn.ReadMessage()
			if err != nil {
				l.Printf("[ERROR]: msg error: %+v", err)
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", err)
				ch <- res
				break RESTART
			}

			typeMsg, err := jsonparser.GetString(msg, "type")
			if typeMsg == "error" {
				l.Printf("[ERROR]: error: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			channel, err := jsonparser.GetString(msg, "channel")
			if err != nil {
				l.Printf("[ERROR]: channel error: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			data, _, _, err := jsonparser.Get(msg, "data")
			if err != nil {
				if isSubscribe, _ := jsonparser.GetString(msg, "type"); isSubscribe == "subscribed" {
					l.Printf("[SUCCESS]: %s %+v", isSubscribe, string(msg))
					continue
				} else {
					err = fmt.Errorf("[ERROR]: data err: %v %s", err, string(msg))
					l.Println(err)
					res.Type = ERROR
					res.Results = err
					ch <- res
					break RESTART
				}
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
	}()

	return nil
}

func signature(conn *websocket.Conn, key, secret string, subaccount []string) error {
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

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%dwebsocket_login", msec)))
	args := map[string]interface{}{
		"key":  key,
		"sign": hex.EncodeToString(mac.Sum(nil)),
		"time": msec,
	}
	if len(subaccount) > 0 {
		args["subaccount"] = subaccount[0]
	}

	if err := conn.WriteJSON(&requestForPrivate{
		Op:   "login",
		Args: args,
	}); err != nil {
		return err
	}

	return nil
}
