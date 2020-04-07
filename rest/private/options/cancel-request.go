package options

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForCancelOpQuoteRequest struct {
	RequestID int
}

type ResponseForCancelOpQuoteRequest struct {
	ID     int `json:"id"`
	Option struct {
		Expiry     time.Time `json:"expiry"`
		Strike     float64   `json:"strike"`
		Type       string    `json:"type"`
		Underlying string    `json:"underlying"`
	} `json:"option"`
	Status        string    `json:"status"`
	Side          string    `json:"side"`
	Size          float64   `json:"size"`
	RequestExpiry time.Time `json:"requestExpiry"`
	Time          time.Time `json:"time"`
}

func (req *RequestForCancelOpQuoteRequest) Path() string {
	return fmt.Sprintf("/options/requests/%d", req.RequestID)
}

func (req *RequestForCancelOpQuoteRequest) Method() string {
	return http.MethodDelete
}

func (req *RequestForCancelOpQuoteRequest) Query() string {
	return ""
}

func (req *RequestForCancelOpQuoteRequest) Payload() []byte {
	return nil
}
