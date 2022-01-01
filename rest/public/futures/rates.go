package futures

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForRates struct {
	ProductCode string `url:"future,omitempty"`
	Start       int64  `url:"start_time,omitempty"`
	End         int64  `url:"end_time,omitempty"`
}

type ResponseForRates []Rate

type ByDate []Rate

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }

type Rate struct {
	Future string    `json:"future"`
	Rate   float64   `json:"rate"`
	Time   time.Time `json:"time"`
}

// Example : https://ftx.com/api/funding_rates?future=DEFI-PERP&start_time=1597687200&end_time=1597773600
func (req *RequestForRates) Path() string {
	return "/funding_rates"
}

func (req *RequestForRates) Method() string {
	return http.MethodGet
}

func (req *RequestForRates) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForRates) Payload() []byte {
	return nil
}

func (a ResponseForRates) Len() int           { return len(a) }
func (a ResponseForRates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResponseForRates) Less(i, j int) bool { return a[i].Rate < a[j].Rate }
