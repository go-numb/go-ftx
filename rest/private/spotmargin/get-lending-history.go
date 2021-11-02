package spotmargin

import (
	"encoding/json"
	"net/http"
	"time"
)

type RequestForLendingHistory struct {
	StartTime int64 `json:"start_time,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`
}

type ResponseForLendingHistory []LendingHistory

type LendingHistory struct {
	Coin string    `json:"coin"`
	Cost float64   `json:"cost"`
	Rate float64   `json:"rate"`
	Size float64   `json:"size"`
	Time time.Time `json:"time"`
}

func (req *RequestForLendingHistory) Path() string {
	return "/spot_margin/lending_history"
}

func (req *RequestForLendingHistory) Method() string {
	return http.MethodGet
}

func (req *RequestForLendingHistory) Query() string {
	return ""
}

func (req *RequestForLendingHistory) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
