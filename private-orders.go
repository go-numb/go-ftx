package ftx

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type RequestForOrder struct {
	Market     string      `json:"market"`
	Side       string      `json:"side"`
	Price      float64     `json:"price"`
	Type       string      `json:"type"`
	Size       float64     `json:"size"`
	ReduceOnly bool        `json:"reduceOnly,omitempty"`
	Ioc        bool        `json:"ioc,omitempty"`
	PostOnly   bool        `json:"postOnly,omitempty"`
	ClientID   interface{} `json:"clientId,omitempty"`
}

type ResponseByOrder struct {
	CreatedAt     time.Time   `json:"createdAt"`
	FilledSize    int         `json:"filledSize"`
	Future        string      `json:"future"`
	ID            int         `json:"id"`
	Market        string      `json:"market"`
	Price         float64     `json:"price"`
	RemainingSize int         `json:"remainingSize"`
	Side          string      `json:"side"`
	Size          int         `json:"size"`
	Status        string      `json:"status"`
	Type          string      `json:"type"`
	ReduceOnly    bool        `json:"reduceOnly"`
	Ioc           bool        `json:"ioc"`
	PostOnly      bool        `json:"postOnly"`
	ClientID      interface{} `json:"clientId"`
}

func (p *Client) Order(o *RequestForOrder) (order []ResponseByOrder, err error) {
	body, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	res, err := p.sendRequest(
		http.MethodPost,
		"/orders",
		bytes.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	// in Close()
	err = decode(res, &order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (p *Client) CancelAll() (status string, err error) {
	res, err := p.sendRequest(
		http.MethodDelete,
		"/orders",
		nil, nil)
	if err != nil {
		return "", err
	}

	// in Close()
	err = decode(res, &status)
	if err != nil {
		return "", err
	}

	return status, nil
}

func (p *Client) CancelByID(oid int) (status string, err error) {
	res, err := p.sendRequest(
		http.MethodDelete,
		fmt.Sprintf("/orders/%d", oid),
		nil, nil)
	if err != nil {
		return "", err
	}

	// in Close()
	err = decode(res, &status)
	if err != nil {
		return "", err
	}

	return status, nil
}
