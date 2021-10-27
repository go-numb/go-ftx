package funding

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type Request struct {
	ProductCode string `url:"future,omitempty"`
	Start       int64  `url:"start_time,omitempty"`
	End         int64  `url:"end_time,omitempty"`
}

type Response []Funding

type Funding struct {
	Future string `json:"future"`

	Payment float64 `json:"payment"`
	Rate    float64 `json:"rate"`

	Time time.Time `json:"time"`
	ID   int       `json:"id"`
}

func (req *Request) Path() string {
	return "/funding_payments"
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
