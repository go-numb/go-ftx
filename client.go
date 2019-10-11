package ftx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	ENDPOINT = "https://ftx.com/api"

	BUY    = "buy"
	SELL   = "sell"
	MARKET = "market"
	LIMIT  = "buy"
)

// Please do not send more than 10 requests per second. Sending requests more frequently will result in HTTP 429 errors.
type Client struct {
	key, secret string
	HTTPC       *http.Client

	Logger *logrus.Entry
}

func New(key, secret string, log *logrus.Logger) *Client {
	hc := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		key:    key,
		secret: secret,
		HTTPC:  hc,
		Logger: logrus.NewEntry(log),
	}
}

type Response struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func (p *Client) newRequest(method, spath string, body io.Reader, params *map[string]string) (*http.Request, error) {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(ENDPOINT)
	u.Path = u.Path + spath

	if params != nil {
		q := u.Query()
		for k, v := range *params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	nonce := fmt.Sprintf("%d", time.Now().UTC().UnixNano()/1000000)
	var q string
	if u.RawQuery != "" {
		q = "?" + u.Query().Encode()
	}
	payload := nonce + method + u.Path + q
	if body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(body)
		payload += buf.String()
	}
	// fmt.Printf("%+v %s, %s\n", payload, p.key, p.secret)
	signture := makeHMAC(p.secret, payload)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", p.key)
	req.Header.Set("FTX-SIGN", signture)
	req.Header.Set("FTX-TS", nonce)

	return req, nil
}

func makeHMAC(secret, body string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

func (c *Client) sendRequest(method, spath string, body io.Reader, params *map[string]string) (*http.Response, error) {
	req, err := c.newRequest(method, spath, body, params)
	c.Logger.Debugf("Request:  %s \n", requestLog(req))
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPC.Do(req)
	c.Logger.Debugf("Response: %s \n", responseLog(res))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}
	return res, nil
}

func decode(res *http.Response, out interface{}) error {
	var r = new(Response)
	r.Result = out
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(r)
	fmt.Printf("%+v\n", r.Result)
	if !r.Success {
		return errors.New("decode error")
	}
	return nil
}

func responseLog(res *http.Response) string {
	b, _ := httputil.DumpResponse(res, true)
	return string(b)
}
func requestLog(req *http.Request) string {
	b, _ := httputil.DumpRequest(req, true)
	return string(b)
}
