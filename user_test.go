package intistelecom

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestGetBalance(t *testing.T) {
	data := map[string]any{
		"amount":   1,
		"currency": "USD",
	}
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	balance, err := GetBalance()
	if err != nil {
		t.Error(err)
	}

	if balance.Amount != 1 {
		t.Errorf("Balance value must be equals to %d", 1)
	}
}

func TestGetMe(t *testing.T) {
	data := map[string]any{
		"id":       1,
		"username": "John_Doe",
	}
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	me, err := GetMe()
	if err != nil {
		t.Error(err)
	}

	if me.Id != 1 {
		t.Errorf("User ID must be equals to %d", 1)
	}
}

func TestGetBalanceError(t *testing.T) {
	var e []string
	e = append(e, "Some Error")
	data := map[string]any{
		"errors":    e,
		"timestamp": time.Now().String(),
	}
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 400,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	_, err := GetBalance()
	if err != nil {
		if err.Error() != "some error" {
			t.Error(err)
		}
	}
}
