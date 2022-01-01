package futures

import (
	"net/http"
	"time"
)

type RequestForFutures struct {
}

type ResponseForFutures []FutureForList

type FutureForList struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Underlying string `json:"underlying"`

	Index        float64 `json:"index"`
	Mark         float64 `json:"mark"`
	Last         float64 `json:"last"`
	Ask          float64 `json:"ask"`
	Bid          float64 `json:"bid"`
	Change1H     float64 `json:"change1h"`
	Change24H    float64 `json:"change24h"`
	ChangeBod    float64 `json:"changeBod"`
	Volume       float64 `json:"volume"`
	VolumeUsd24H float64 `json:"volumeUsd24h"`

	OpenInterestUsd     float64 `json:"openInterestUsd"`
	PositionLimitWeight float64 `json:"positionLimitWeight"`

	PriceIncrement float64 `json:"priceIncrement"`
	SizeIncrement  float64 `json:"sizeIncrement"`

	UpperBound float64 `json:"upperBound"`
	LowerBound float64 `json:"lowerBound"`

	Description string    `json:"description"`
	Expiry      time.Time `json:"expiry"`
	Enabled     bool      `json:"enabled"`
	Expired     bool      `json:"expired"`
	Perpetual   bool      `json:"perpetual"`
	PostOnly    bool      `json:"postOnly"`
}

func (req *RequestForFutures) Path() string {
	return "/futures"
}

func (req *RequestForFutures) Method() string {
	return http.MethodGet
}

func (req *RequestForFutures) Query() string {
	return ""
}

func (req *RequestForFutures) Payload() []byte {
	return nil
}

func (futures ResponseForFutures) Products() []string {
	list := make([]string, len(futures))
	for i := range futures {
		list[i] = futures[i].Name
	}
	return list
}

// Sort by alphabetical order (by Name)
func (a ResponseForFutures) Len() int           { return len(a) }
func (a ResponseForFutures) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResponseForFutures) Less(i, j int) bool { return a[i].Name < a[j].Name }
