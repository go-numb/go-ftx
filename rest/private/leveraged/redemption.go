package leveraged

import (
	"fmt"
	"net/http"
	"time"
)

type RequestForRedemptionLvToken struct {
	ProductCode string `json:"-"`
	// to body
	Size string `json:"size"`
}

type ResponseForRedemptionLvToken struct {
	Token             string    `json:"token"`
	Size              float64   `json:"size"`
	ProjectedProceeds float64   `json:"projectedProceeds"`
	Pending           bool      `json:"pending"`
	RequestedAt       time.Time `json:"requestedAt"`

	ID int `json:"id"`
}

func (req *RequestForRedemptionLvToken) Path() string {
	return fmt.Sprintf("/lt/%s/redeem", req.ProductCode)
}

func (req *RequestForRedemptionLvToken) Method() string {
	return http.MethodPost
}

func (req *RequestForRedemptionLvToken) Query() string {
	return ""
}

func (req *RequestForRedemptionLvToken) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
