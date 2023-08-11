package intistelecom

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestCost(t *testing.T) {
	data := Routing{
		Currency: "USD",
		Mcc:      "250",
		Mnc:      "01",
		Price:    0.1,
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

	_, err := Cost("17778889966")
	if err != nil {
		t.Error(err)
	}
}
