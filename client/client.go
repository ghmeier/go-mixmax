package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ghmeier/go-service"
)

type Client struct {
	Key string
	URL string
	S   service.Service
}

const (
	version = "v1"
)

func New(key string) *Client {
	c := &Client{
		Key: key,
		S:   service.NewCustom(&responder{}),
		URL: "https://api.mixmax.com/" + version,
	}
	return c
}

type responder struct {
}

type Response struct {
	Message string `json: "message,omitempty"`
	body    []byte
}

func (r *responder) Marshal(res *http.Response) (service.Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var raw Response
	if body == nil || len(body) == 0 {
		raw.body = make([]byte, 0)
		return &raw, nil
	}

	if strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		raw.Message = string(body)
		return &raw, nil

	}
	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, err
	}

	raw.body = body
	return &raw, nil
}

func (r *Response) Error() error {
	if r.Message == "" {
		return nil
	}

	return fmt.Errorf(r.Message)
}

func (r *Response) Body() ([]byte, error) {
	return r.body, nil
}
