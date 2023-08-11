package intistelecom

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestTemplates(t *testing.T) {
	var data []Template
	data = append(data, Template{
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		Template:  "some template",
		Title:     "some title",
		Id:        1,
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

	_, err := Templates()
	if err != nil {
		t.Error(err)
	}
}

func TestEditTemplate(t *testing.T) {
	tpl := Template{
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		Template:  "some template",
		Title:     "some title",
		Id:        1,
	}
	baseTemplate := &BaseTemplate{
		Template: "some template",
		Title:    "some title",
		Id:       1,
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

	template, err := EditTemplate(baseTemplate)
	if err != nil {
		t.Error(err)
	}
	if template.Title != baseTemplate.Title {
		t.Errorf("Template title not equals `%s`", baseTemplate.Title)
	}
}

func TestCreateTemplate(t *testing.T) {
	tpl := Template{
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		Template:  "some template",
		Title:     "some title",
		Id:        1,
	}
	baseTemplate := &BaseTemplate{
		Template: "some template",
		Title:    "some title",
		Id:       1,
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

	template, err := CreateTemplate(baseTemplate)
	if err != nil {
		t.Error(err)
	}
	if template.Title != baseTemplate.Title {
		t.Errorf("Template title not equals `%s`", baseTemplate.Title)
	}
}

func TestDeleteTemplate(t *testing.T) {
	client := getClient()
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString("")),
			Header:     make(http.Header),
		}
	})

	_, err := DeleteTemplate(123456)
	if err != nil {
		t.Error(err)
	}
}
