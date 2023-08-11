package intistelecom

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestOriginators(t *testing.T) {
	var data []Originator
	data = append(data, Originator{
		CreatedAt:  time.Now().String(),
		LastUsedAt: time.Now().String(),
		Originator: "test",
		State:      "COMPLETED",
		Id:         1,
		Default:    true,
	})
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	_, err := Originators()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateOriginator(t *testing.T) {
	tpl := Originator{
		CreatedAt:  time.Now().String(),
		LastUsedAt: time.Now().String(),
		Originator: "test",
		State:      "COMPLETED",
		Id:         1,
		Default:    true,
	}
	baseOriginator := &BaseOriginator{
		Originator: "test",
		Default:    false,
	}
	json, _ := client.JSONEncoder(tpl)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	originator, err := CreateOriginator(baseOriginator)
	if err != nil {
		t.Error(err)
	}
	if originator.Originator != baseOriginator.Originator {
		t.Errorf("Originator not equals `%s`", baseOriginator.Originator)
	}
}

func TestSetDefaultOriginator(t *testing.T) {
	tpl := Originator{
		CreatedAt:  time.Now().String(),
		LastUsedAt: time.Now().String(),
		Originator: "test",
		State:      "COMPLETED",
		Id:         1,
		Default:    true,
	}
	json, _ := client.JSONEncoder(tpl)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	_, err := SetDefaultOriginator("test")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOriginator(t *testing.T) {
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString("")),
			Header:     make(http.Header),
		}
	})

	_, err := DeleteOriginator("test")
	if err != nil {
		t.Error(err)
	}
}
