package leveraged

import (
	"net/http"
	"time"
)

type RequestForCreatedLvTokens struct {
}

type ResponseForCreatedLvTokens []Creation

type Creation struct {
	ID    int    `json:"id"`
	Token string `json:"token"`

	Price         float64   `json:"price"`
	RequestedSize float64   `json:"requestedSize"`
	CreatedSize   float64   `json:"createdSize"`
	Cost          float64   `json:"cost"`
	Fee           float64   `json:"fee"`
	RequestedAt   time.Time `json:"requestedAt"`
	FulfilledAt   time.Time `json:"fulfilledAt"`

	Pending bool `json:"pending"`
}

func (req *RequestForCreatedLvTokens) Path() string {
	return "/lt/creations"
}

func (req *RequestForCreatedLvTokens) Method() string {
	return http.MethodGet
}

func (req *RequestForCreatedLvTokens) Query() string {
	return ""
}

func (req *RequestForCreatedLvTokens) Payload() []byte {
	return nil
}
