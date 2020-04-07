package futures

import (
	"net/http"
	"time"
)

type RequestForRates struct {
}

type ResponseForRates []Rate

type Rate struct {
	Future string    `json:"future"`
	Rate   float64   `json:"rate"`
	Time   time.Time `json:"time"`
}

func (req *RequestForRates) Path() string {
	return "/funding_rates"
}

func (req *RequestForRates) Method() string {
	return http.MethodGet
}

func (req *RequestForRates) Query() string {
	return ""
}

func (req *RequestForRates) Payload() []byte {
	return nil
}

func (a ResponseForRates) Len() int           { return len(a) }
func (a ResponseForRates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResponseForRates) Less(i, j int) bool { return a[i].Rate < a[j].Rate }
