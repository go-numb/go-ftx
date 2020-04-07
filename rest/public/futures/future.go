package futures

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForFuture struct {
	ProductCode string `url:"-"`
}

type ResponseForFuture Future

type Future struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Underlying string `json:"underlying"`

	Index     float64 `json:"index"`
	Last      float64 `json:"last"`
	Mark      float64 `json:"mark"`
	Ask       float64 `json:"ask"`
	Bid       float64 `json:"bid"`
	Change1H  float64 `json:"change1h"`
	Change24H float64 `json:"change24h"`

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

func (req *RequestForFuture) Path() string {
	return fmt.Sprintf("/futures/%s", req.ProductCode)
}

func (req *RequestForFuture) Method() string {
	return http.MethodGet
}

func (req *RequestForFuture) Query() string {
	return ""
}

func (req *RequestForFuture) Payload() []byte {
	return nil
}
