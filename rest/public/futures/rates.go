package futures

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RequestForRates struct {
	ProductCode string `json:"future,omitempty"`
	Start       int64  `json:"start_time,omitempty"`
	End         int64  `json:"end_time,omitempty"`
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
	return fmt.Sprintf("/funding_rates")
}

func (req *RequestForRates) Method() string {
	return http.MethodGet
}

func (req *RequestForRates) Query() string {
	// value, _ := query.Values(req)
	// return value.Encode()
	return ""
}

func (req *RequestForRates) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}

func (a ResponseForRates) Len() int           { return len(a) }
func (a ResponseForRates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResponseForRates) Less(i, j int) bool { return a[i].Rate < a[j].Rate }
