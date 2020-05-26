package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type Response struct {
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
	Success bool        `json:"success"`
}

func (p *Client) request(req Requester, results interface{}) error {
	res, err := p.do(req)
	if err != nil {
		return err
	}

	if err := decode(res, results); err != nil {
		return err
	}
	return nil
}

func signture(secret, body string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

func (p *Client) newRequest(r Requester) *fasthttp.Request {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(ENDPOINT)
	u.Path = u.Path + r.Path()
	u.RawQuery = r.Query()

	// fmt.Printf("%+v\n", u.String())

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(u.String())
	body := r.Payload()
	req.SetBody(body)

	if p.Auth != nil {
		nonce := fmt.Sprintf("%d", int64(time.Now().UTC().UnixNano()/int64(time.Millisecond)))
		payload := nonce + r.Method() + u.Path + u.RawQuery + string(body)
		// fmt.Printf("%+v\n", payload)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("FTX-KEY", p.Auth.Key)
		req.Header.Set("FTX-SIGN", p.Auth.Signture(payload))
		req.Header.Set("FTX-TS", nonce)

		// set id is there UseSubAccountID
		subaccount := p.Auth.SubAccount()
		if subaccount.Nickname != "" {
			req.Header.Set("FTX-SUBACCOUNT", url.PathEscape(subaccount.Nickname))
		}
	}

	return req
}

func (c *Client) do(r Requester) (*fasthttp.Response, error) {
	req := c.newRequest(r)

	// fasthttp for http2.0
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", string(res.Body()))
	// no usefull headers

	if res.StatusCode() != 200 {
		var r = new(Response)
		if err := json.Unmarshal(res.Body(), r); err != nil {
			return nil, fmt.Errorf("faild to get data. status: %d", res.StatusCode())
		}

		if !r.Success {
			return nil, fmt.Errorf("faild to get data. status: %d - %s", res.StatusCode(), r.Error)
		}
	}
	return res, nil
}

func decode(res *fasthttp.Response, out interface{}) error {
	var r = new(Response)
	r.Result = out

	if err := json.Unmarshal(res.Body(), r); err != nil {
		return err
	}
	if !r.Success {
		return fmt.Errorf("decode error")
	}
	return nil
}
