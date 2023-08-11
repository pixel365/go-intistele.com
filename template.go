package intistelecom

import (
	"bytes"
	"fmt"
)

type BaseTemplate struct {
	Template string `json:"template"`
	Title    string `json:"title"`
	Id       int32  `json:"id"`
}

type Template struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Template  string `json:"template"`
	Title     string `json:"title"`
	Id        int32  `json:"id"`
}

func Templates() ([]Template, error) {
	cl := getClient()
	var result []Template
	response, err := cl.Get("/template")
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func EditTemplate(template *BaseTemplate) (Template, error) {
	cl := getClient()
	result := Template{}
	payload, err := cl.JSONEncoder(template)
	if err != nil {
		return result, err
	}
	response, err := cl.Put("/template", bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func CreateTemplate(template *BaseTemplate) (Template, error) {
	cl := getClient()
	result := Template{}
	payload, err := cl.JSONEncoder(template)
	if err != nil {
		return result, err
	}
	response, err := cl.Post("/template", bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func DeleteTemplate(id int32) (bool, error) {
	cl := getClient()
	_, err := cl.Delete(fmt.Sprintf("/template/%d", id))
	return err == nil, err
}
