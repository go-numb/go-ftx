package markets

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// query
// ?resolution={resolution}&limit={limit}&start_time={start_time}&end_time={end_time}
type RequestForCandles struct {
	ProductCode string `url:"-"`
	// Resolution is sec length, window length in seconds. options: 15, 60(1m), 300(5m), 900(15m), 3600(60m), 14400(4h), 86400(1D)
	Resolution int   `url:"resolution"`
	Start      int64 `url:"start_time,omitempty"`
	End        int64 `url:"end_time,omitempty"`
}

type ResponseForCandles []Candle

type Candle struct {
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Open      float64   `json:"open"`
	Volume    float64   `json:"volume"`
	StartTime time.Time `json:"startTime"`
}

func (req *RequestForCandles) Path() string {
	return fmt.Sprintf("/markets/%s/candles", req.ProductCode)
}

func (req *RequestForCandles) Method() string {
	return http.MethodGet
}

func (req *RequestForCandles) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *RequestForCandles) Payload() []byte {
	return nil
}

// Sort by timestamp
func (a ResponseForCandles) Len() int           { return len(a) }
func (a ResponseForCandles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResponseForCandles) Less(i, j int) bool { return a[i].StartTime.Before(a[j].StartTime) }
