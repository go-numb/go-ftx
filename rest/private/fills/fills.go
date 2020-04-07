package fills

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode string `url:"market,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	Start       int64  `url:"start,omitempty"`
	End         int64  `url:"end,omitempty"`
}

type Response []Fill

type Fill struct {
	Future    string `json:"future"`
	Market    string `json:"market"`
	Type      string `json:"type"`
	Liquidity string `json:"liquidity"`

	// only rest follow 2factor
	BaseCurrency  string `json:"baseCurrency"`
	QuoteCurrency string `json:"quoteCurrency"`

	Side string `json:"side"`

	Price   float64 `json:"price"`
	Size    float64 `json:"size"`
	Fee     float64 `json:"fee"`
	FeeRate float64 `json:"feeRate"`

	Time time.Time `json:"time"`

	ID      int `json:"id"`
	OrderID int `json:"orderId"`
	TradeID int `json:"tradeId"`
}

func (req *Request) Path() string {
	return "/fills"
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
