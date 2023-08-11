package intistelecom

import (
	"bytes"
)

type BaseOriginator struct {
	Originator string `json:"originator"`
	Default    bool   `json:"default"`
}

type Originator struct {
	CreatedAt  string `json:"created_at"`
	LastUsedAt string `json:"last_used_at"`
	Originator string `json:"originator"`
	State      string `json:"state"`
	Id         int32  `json:"id"`
	Default    bool   `json:"default"`
}

func Originators() ([]Originator, error) {
	cl := getClient()
	var result []Originator
	response, err := cl.Get("/originator")
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func CreateOriginator(originator *BaseOriginator) (Originator, error) {
	cl := getClient()
	result := Originator{}
	payload, err := cl.JSONEncoder(originator)
	if err != nil {
		return result, err
	}
	response, err := cl.Post("/originator", bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func SetDefaultOriginator(originator string) (Originator, error) {
	cl := getClient()
	result := Originator{}
	response, err := cl.Put("/originator/default/"+originator, nil)
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func DeleteOriginator(originator string) (bool, error) {
	cl := getClient()
	_, err := cl.Delete("/originator/" + originator)
	return err == nil, err
}
