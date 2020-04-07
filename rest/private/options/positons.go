package options

import (
	"net/http"
)

type RequestForOpPositions struct {
}

type ResponseForOpPositions []OpPosition

type OpPosition struct {
	Side       string  `json:"side"`
	EntryPrice float64 `json:"entryPrice"`
	Size       float64 `json:"size"`
	NetSize    float64 `json:"netSize"`
	Option     Option  `json:"option"`
}

func (req *RequestForOpPositions) Path() string {
	return "/options/positions"
}

func (req *RequestForOpPositions) Method() string {
	return http.MethodGet
}

func (req *RequestForOpPositions) Query() string {
	return ""
}

func (req *RequestForOpPositions) Payload() []byte {
	return nil
}
