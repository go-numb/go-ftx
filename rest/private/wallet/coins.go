package wallet

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForCoins struct {
}

type ResponseForCoins []Coin

type Coin struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	CanDeposit  bool `json:"canDeposit"`
	CanWithdraw bool `json:"canWithdraw"`
	HasTag      bool `json:"hasTag"`

	Collateral               bool     `json:"collateral"`
	UsdFungible              bool     `json:"usdFungible"`
	IsEtf                    bool     `json:"isEtf"`
	IsToken                  bool     `json:"isToken"`
	Hidden                   bool     `json:"hidden"`
	CanConvert               bool     `json:"canConvert"`
	CollateralWeight         float64  `json:"collateralWeight"`
	Fiat                     bool     `json:"fiat"`
	Methods                  []string `json:"methods"`
	Erc20Contract            string   `json:"erc20Contract"`
	Bep2Asset                string   `json:"bep2Asset"`
	Trc20Contract            string   `json:"trc20Contract"`
	SplMint                  string   `json:"splMint"`
	CreditTo                 string   `json:"creditTo"`
	SpotMargin               bool     `json:"spotMargin"`
	NftQuoteCurrencyEligible bool     `json:"nftQuoteCurrencyEligible"`
	IndexPrice               float64  `json:"indexPrice"`
}

func (req *RequestForCoins) Path() string {
	return "/wallet/coins"
}

func (req *RequestForCoins) Method() string {
	return http.MethodGet
}

func (req *RequestForCoins) Query() string {
	return ""
}

func (req *RequestForCoins) Payload() []byte {
	return nil
}
