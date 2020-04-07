package leveraged

import (
	"net/http"
	"time"
)

type RequestForRedemptionLvTokens struct {
}

type ResponseForRedemptionLvTokens []Redemption

type Redemption struct {
	ID          int       `json:"id"`
	Token       string    `json:"token"`
	Size        float64   `json:"size"`
	Price       float64   `json:"price"`
	Proceeds    float64   `json:"proceeds"`
	Fee         float64   `json:"fee"`
	RequestedAt time.Time `json:"requestedAt"`
	FulfilledAt time.Time `json:"fulfilledAt"`

	Pending bool `json:"pending"`
}

func (req *RequestForRedemptionLvTokens) Path() string {
	return "/lt/redemptions"
}

func (req *RequestForRedemptionLvTokens) Method() string {
	return http.MethodGet
}

func (req *RequestForRedemptionLvTokens) Query() string {
	return ""
}

func (req *RequestForRedemptionLvTokens) Payload() []byte {
	return nil
}
