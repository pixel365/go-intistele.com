package intistelecom

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestSend(t *testing.T) {
	data := MessageId{Id: "585587529782352"}
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	message := MessageBody{
		Destination:    "17778889966",
		Originator:     "test",
		Text:           "some text",
		TimeToSend:     "",
		CallbackURL:    "https://some-callback.com",
		ValidityPeriod: 60,
		UseLocaltime:   false,
	}
	result, err := Send(&message)
	if err != nil {
		t.Error(err)
	}

	if result.Id == "" {
		t.Error("Invalid ID")
	}
}

func TestBatch(t *testing.T) {
	var data []MessageId
	data = append(data, MessageId{Id: "585587529782352"})
	data = append(data, MessageId{Id: "585587529782352"})
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	var payload []MessageBody
	message := MessageBody{
		Destination:    "17778889966",
		Originator:     "test",
		Text:           "some text",
		TimeToSend:     "",
		CallbackURL:    "https://some-callback.com",
		ValidityPeriod: 60,
		UseLocaltime:   false,
	}
	payload = append(payload, message)
	payload = append(payload, message)
	result, err := Batch(&payload)
	if err != nil {
		t.Error(err)
	}

	if len(result) != 2 {
		t.Error("Batch error")
	}
}

func TestCancelMessage(t *testing.T) {
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString("")),
			Header:     make(http.Header),
		}
	})

	_, err := CancelMessage("6729263692823")
	if err != nil {
		t.Error(err)
	}
}

func TestMessageStatus(t *testing.T) {
	var data []MessageDetails
	details := MessageDetails{
		MessageId:        "234532723462897",
		PartId:           "234532723462897",
		State:            "deliver",
		Sender:           "test",
		Network:          "network",
		Dcs:              "8",
		ErrorDescription: "",
		TimeChangeState:  "",
		Cost:             0.5,
		ErrorId:          0,
		Ported:           false,
	}
	data = append(data, details)
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	_, err := MessageStatus("234532723462897")
	if err != nil {
		t.Error(err)
	}
}

func TestMessagePartStatus(t *testing.T) {
	var data []MessageDetails
	details := MessageDetails{
		MessageId:        "234532723462897",
		PartId:           "234532723462897",
		State:            "deliver",
		Sender:           "test",
		Network:          "network",
		Dcs:              "8",
		ErrorDescription: "",
		TimeChangeState:  "",
		Cost:             0.5,
		ErrorId:          0,
		Ported:           false,
	}
	data = append(data, details)
	json, _ := client.JSONEncoder(data)
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(json))),
			Header:     make(http.Header),
		}
	})

	_, err := MessagePartStatus("234532723462897")
	if err != nil {
		t.Error(err)
	}
}
