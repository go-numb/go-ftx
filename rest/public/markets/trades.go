package markets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-numb/go-ftx/types"
)

// query
// ?limit={limit}&start_time={start_time}&end_time={end_time}
type RequestForTrades struct {
	ProductCode string `json:"-"`
	Limit       int    `json:"limit,omitempty"`
	Start       int64  `json:"start,omitempty"`
	End         int64  `json:"end,omitempty"`
}

type ResponseForTrades []Trade

type Trade struct {
	ID          int       `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        string    `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

type Ticker struct {
	Bid     float64       `json:"bid"`
	Ask     float64       `json:"ask"`
	BidSize float64       `json:"bidSize"`
	AskSize float64       `json:"askSize"`
	Last    float64       `json:"last"`
	Time    types.FtxTime `json:"time"`
}

// This syntax works to request historical prices
// https://ftx.com/api/markets/DEFI-PERP/trades?&start_time=1597687200&end_time=1597773600
func (req *RequestForTrades) Path() string {
	return fmt.Sprintf("/markets/%s/trades", req.ProductCode)
}

func (req *RequestForTrades) Method() string {
	return http.MethodGet
}

func (req *RequestForTrades) Query() string {
	// values, _ := query.Values(req)
	// return values.Encode()
	return ""
}

func (req *RequestForTrades) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
