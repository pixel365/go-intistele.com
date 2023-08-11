package intistelecom

import (
	"bytes"
)

type MessageBody struct {
	Destination    string `json:"destination"`
	Originator     string `json:"originator"`
	Text           string `json:"text"`
	TimeToSend     string `json:"time_to_send"`
	CallbackURL    string `json:"callback_url"`
	ValidityPeriod int32  `json:"validity_period"`
	UseLocaltime   bool   `json:"use_localtime"`
}

type MessageId struct {
	Id string `json:"message_id"`
}

type MessageDetails struct {
	MessageId        string  `json:"message_id"`
	PartId           string  `json:"part_id"`
	State            string  `json:"state"`
	Sender           string  `json:"sender"`
	Network          string  `json:"network"`
	Dcs              string  `json:"dcs"`
	ErrorDescription string  `json:"error_description"`
	TimeChangeState  string  `json:"time_change_state"`
	Cost             float32 `json:"cost"`
	ErrorId          int32   `json:"error_id"`
	Ported           bool    `json:"ported"`
}

func Send(message *MessageBody) (MessageId, error) {
	cl := getClient()
	result := MessageId{}
	payload, err := cl.JSONEncoder(message)
	if err != nil {
		return result, err
	}

	response, err := cl.Post("/message/send", bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func Batch(messages *[]MessageBody) ([]MessageId, error) {
	cl := getClient()
	var result []MessageId
	payload, err := cl.JSONEncoder(messages)
	if err != nil {
		return result, err
	}
	response, err := cl.Post("/message/batch", bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func CancelMessage(messageId string) (bool, error) {
	cl := getClient()
	_, err := cl.Delete("/message/cancel/" + messageId)
	return err == nil, err
}

func MessageStatus(messageId string) ([]MessageDetails, error) {
	cl := getClient()
	var result []MessageDetails
	response, err := cl.Get("/message/status/" + messageId)
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func MessagePartStatus(partId string) ([]MessageDetails, error) {
	cl := getClient()
	var result []MessageDetails
	response, err := cl.Get("/message/status/part/" + partId)
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}
