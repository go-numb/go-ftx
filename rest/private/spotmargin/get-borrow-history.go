package spotmargin

import (
	"encoding/json"
	"net/http"
	"time"
)

type RequestForBorrowHistory struct {
	StartTime int64 `json:"start_time,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`
}

type ResponseForBorrowHistory []BorrowHistory

type BorrowHistory struct {
	Coin string    `json:"coin"`
	Cost float64   `json:"cost"`
	Rate float64   `json:"rate"`
	Size float64   `json:"size"`
	Time time.Time `json:"time"`
}

func (req *RequestForBorrowHistory) Path() string {
	return "/spot_margin/borrow_history"
}

func (req *RequestForBorrowHistory) Method() string {
	return http.MethodGet
}

func (req *RequestForBorrowHistory) Query() string {
	return ""
}

func (req *RequestForBorrowHistory) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
