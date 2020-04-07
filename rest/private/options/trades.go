package options

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForOpTrades struct {
	Limit int   `url:"limit,omitempty"`
	Start int64 `url:"start_time,omitempty"`
	End   int64 `url:"end_time,omitempty"`
}

type ResponseForOpTrades []OpTrade

type OpTrade struct {
	ID     int       `json:"id"`
	Price  float64   `json:"price"`
	Size   float64   `json:"size"`
	Option Option    `json:"option"`
	Time   time.Time `json:"time"`
}

func (req *RequestForOpTrades) Path() string {
	return "/options/trades"
}

func (req *RequestForOpTrades) Method() string {
	return http.MethodGet
}

func (req *RequestForOpTrades) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForOpTrades) Payload() []byte {
	return nil
}
