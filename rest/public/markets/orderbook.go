package markets

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// query
// ?depth={depth}
// max 100, default 20
type RequestForOrderbook struct {
	ProductCode string `url:"-"`
	Depth       int    `url:"depth,omitempty"`
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
	if req.Depth == 0 {
		req.Depth = 20
	}
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *RequestForOrderbook) Payload() []byte {
	return nil
}
