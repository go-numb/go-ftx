package markets

import (
	"fmt"
	"math"
	"net/http"
	"sort"

	"github.com/dustin/go-humanize"
)

// RequestForMarkets arg productcode can be blank
type RequestForMarkets struct {
	ProductCode string `url:"-"`
}

type ResponseForMarkets []Market

type Market struct {
	Type          string `json:"type"`
	Name          string `json:"name"`
	BaseCurrency  string `json:"baseCurrency"`
	QuoteCurrency string `json:"quoteCurrency"`
	Underlying    string `json:"underlying"`

	Last           float64 `json:"last"`
	Ask            float64 `json:"ask"`
	Bid            float64 `json:"bid"`
	Price          float64 `json:"price"`
	PriceIncrement float64 `json:"priceIncrement"`

	SizeIncrement  float64 `json:"sizeIncrement"`
	ChangeBod      float64 `json:"changeBod"`
	Change1H       float64 `json:"change1h"`
	Change24H      float64 `json:"change24h"`
	QuoteVolume24H float64 `json:"quoteVolume24h"`
	VolumeUsd24H   float64 `json:"volumeUsd24h"`

	Enabled bool `json:"enabled"`
}

func (req *RequestForMarkets) Path() string {
	if req.ProductCode != "" {
		 // fmt.Println("/markets/%s", req.ProductCode)
		return fmt.Sprintf("/markets/%s", req.ProductCode)
	}
	return "/markets"
}

func (req *RequestForMarkets) Method() string {
	return http.MethodGet
}

func (req *RequestForMarkets) Query() string {
	return ""
}

func (req *RequestForMarkets) Payload() []byte {
	return nil
}

func (markets ResponseForMarkets) List() []string {
	list := make([]string, len(markets)+1)
	list[0] = fmt.Sprint("TYPE	NAME	BASE	QUOTE	UNDER")
	for i := range markets {
		list[i+1] = fmt.Sprintf("%s	%s	%s	%s	%s", markets[i].Type, markets[i].Name, markets[i].BaseCurrency, markets[i].QuoteCurrency, markets[i].Underlying)
	}
	sort.Strings(list)
	return list
}

const (
	ALL = iota
	BASESPOT
	BASEFUTURE
)

// Ranking return USD volume in 24hours.
func (markets ResponseForMarkets) Ranking(base int) []string {
	var list []string

	sort.Sort(sort.Reverse(markets))

	switch base {
	case BASESPOT:
		for i := range markets {
			if markets[i].Type != "spot" {
				continue
			}
			list = append(list, fmt.Sprintf("$%s	%s", humanize.Commaf(math.Round(markets[i].VolumeUsd24H)), markets[i].Name))
		}

	case BASEFUTURE:
		for i := range markets {
			if markets[i].Type != "future" {
				continue
			}
			list = append(list, fmt.Sprintf("$%s	%s", humanize.Commaf(math.Round(markets[i].VolumeUsd24H)), markets[i].Name))
		}

	default:
		ll := make([]string, len(markets))
		for i := range markets {
			ll[i] = fmt.Sprintf("$%s	%s", humanize.Commaf(math.Round(markets[i].VolumeUsd24H)), markets[i].Name)
		}
		list = ll
	}

	return list
}

func (a ResponseForMarkets) Len() int           { return len(a) }
func (a ResponseForMarkets) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ResponseForMarkets) Less(i, j int) bool { return a[i].VolumeUsd24H < a[j].VolumeUsd24H }
