package intistelecom

import (
	"encoding/base64"
	"errors"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type JSONMarshal func(v interface{}) ([]byte, error)

type JSONUnmarshal func(data []byte, v interface{}) error

type Error struct {
	Errors    []string `json:"errors"`
	TimeStamp string   `json:"timestamp"`
}

type Client struct {
	*http.Client
	JSONEncoder   JSONMarshal
	JSONDecoder   JSONUnmarshal
	BaseUrl       string
	UserName      string
	ApiKey        string
	authorization string
}

func (c *Client) Authorization() string {
	if c.authorization != "" {
		return c.authorization
	}
	if c.UserName != "" && c.ApiKey != "" {
		c.authorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.UserName+":"+c.ApiKey))
	}

	return c.authorization
}

func NewClient() *Client {
	userName := os.Getenv("INTISTELECOM_USERNAME")
	apiKey := os.Getenv("INTISTELECOM_API_KEY")
	return &Client{
		Client: &http.Client{
			Timeout: time.Second * 5,
		},
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		BaseUrl:     "https://api-go2.intistele.com",
		UserName:    userName,
		ApiKey:      apiKey,
	}
}

func getClient() *Client {
	return client
}

func handleError(client *Client, data *[]byte) error {
	e := Error{}
	err := client.JSONDecoder(*data, &e)
	if err != nil {
		return err
	}
	return errors.New(strings.ToLower(e.Errors[0]))
}

func request(client *Client, method string, path string, payload io.Reader) ([]byte, error) {
	req, _ := http.NewRequest(method, client.BaseUrl+path, payload)
	req.Header.Set("Authorization", client.Authorization())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, handleError(client, &body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (c *Client) Get(path string) ([]byte, error) {
	return request(c, http.MethodGet, path, nil)
}

func (c *Client) Post(path string, payload io.Reader) ([]byte, error) {
	return request(c, http.MethodPost, path, payload)
}

func (c *Client) Delete(path string) ([]byte, error) {
	return request(c, http.MethodDelete, path, nil)
}

func (c *Client) Put(path string, payload io.Reader) ([]byte, error) {
	return request(c, http.MethodPut, path, payload)
}
