package futures

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForStats struct {
	ProductCode string `url:"-"`
}

type ResponseForStats Stats

type StatsList []Stats
type Stats struct {
	Name                     string  `json:"name,omitempty"`
	Volume                   float64 `json:"volume"`
	NextFundingRate          float64 `json:"nextFundingRate"`
	ExpirationPrice          float64 `json:"expirationPrice"`
	PredictedExpirationPrice float64 `json:"predictedExpirationPrice"`
	StrikePrice              float64 `json:"strikePrice"`
	OpenInterest             float64 `json:"openInterest"`

	NextFundingTime time.Time `json:"nextFundingTime"`
}

func (req *RequestForStats) Path() string {
	return fmt.Sprintf("/futures/%s/stats", req.ProductCode)
}

func (req *RequestForStats) Method() string {
	return http.MethodGet
}

func (req *RequestForStats) Query() string {
	return ""
}

func (req *RequestForStats) Payload() []byte {
	return nil
}

func (a StatsList) Len() int      { return len(a) }
func (a StatsList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a StatsList) Less(i, j int) bool {
	return a[i].NextFundingRate < a[j].NextFundingRate
}
