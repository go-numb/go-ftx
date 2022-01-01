package markets

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// query
// ?depth={depth}
type RequestForOrderbook struct {
	ProductCode string `json:"-"`
	Depth       int    `json:"depth,omitempty"`
}

type ResponseForOrderbook Orderbook

type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

func (req *RequestForOrderbook) Path() string {
	return fmt.Sprintf("/markets/%s/orderbook", req.ProductCode)
}

func (req *RequestForOrderbook) Method() string {
	return http.MethodGet
}

func (req *RequestForOrderbook) Query() string {
	// values, _ := query.Values(req)
	// return values.Encode()
	return ""
}

func (req *RequestForOrderbook) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
