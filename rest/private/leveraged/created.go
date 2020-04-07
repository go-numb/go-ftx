package leveraged

import (
	"fmt"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForCreatedLvToken struct {
	ProductCode string `json:"-"`
	// to body
	Size string `json:"size"`
}

type ResponseForCreatedLvToken struct {
	Token         string    `json:"token"`
	RequestedSize float64   `json:"requestedSize"`
	Cost          float64   `json:"cost"`
	RequestedAt   time.Time `json:"requestedAt"`

	Pending bool `json:"pending"`
	ID      int  `json:"id"`
}

func (req *RequestForCreatedLvToken) Path() string {
	return fmt.Sprintf("/lt/%s/create", req.ProductCode)
}

func (req *RequestForCreatedLvToken) Method() string {
	return http.MethodPost
}

func (req *RequestForCreatedLvToken) Query() string {
	return ""
}

func (req *RequestForCreatedLvToken) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
