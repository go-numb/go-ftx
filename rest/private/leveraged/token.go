package leveraged

import (
	"fmt"
	"net/http"
)

type RequestForLvToken struct {
	ProductCode string
}

type ResponseForLvToken LvToken

func (req *RequestForLvToken) Path() string {
	return fmt.Sprintf("/lt/%s", req.ProductCode)
}

func (req *RequestForLvToken) Method() string {
	return http.MethodGet
}

func (req *RequestForLvToken) Query() string {
	return ""
}

func (req *RequestForLvToken) Payload() []byte {
	return nil
}
