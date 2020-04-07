package leveraged

import (
	"net/http"
)

type RequestForLvTokens struct {
}

type ResponseForLvTokens []LvToken

type LvToken struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Underlying      string `json:"underlying"`
	ContractAddress string `json:"contractAddress"`

	Leverage         float64 `json:"leverage"`
	Outstanding      float64 `json:"outstanding"`
	PricePerShare    float64 `json:"pricePerShare"`
	PositionPerShare float64 `json:"positionPerShare"`
	UnderlyingMark   float64 `json:"underlyingMark"`
	Change1H         float64 `json:"change1h"`
	Change24H        float64 `json:"change24h"`
}

func (req *RequestForLvTokens) Path() string {
	return "/lt/tokens"
}

func (req *RequestForLvTokens) Method() string {
	return http.MethodGet
}

func (req *RequestForLvTokens) Query() string {
	return ""
}

func (req *RequestForLvTokens) Payload() []byte {
	return nil
}

func (tokens ResponseForLvTokens) Products() []string {
	list := make([]string, len(tokens))
	for i := range tokens {
		list[i] = tokens[i].Name
	}
	return list
}
