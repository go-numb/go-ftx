package types

import (
	"math"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	BUY    = "buy"
	SELL   = "sell"
	MARKET = "market"
	LIMIT  = "limit"
	STOP = "stop"
)

type FtxTime struct {
	Time time.Time
}

func (p *FtxTime) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	sec, nsec := math.Modf(f)
	p.Time = time.Unix(int64(sec), int64(nsec))
	return nil
}
