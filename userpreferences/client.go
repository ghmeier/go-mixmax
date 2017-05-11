package userpreferences

import (
	"net/http"
	"strings"

	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-service"
)

type Client struct {
	*client.Client
}

func New(c *client.Client) *Client {
	return &Client{
		Client: &client.Client{
			Key: c.Key,
			S:   c.S.Copy("userpreferences", "me"),
		},
	}
}

func (c *Client) Me(fields []string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	err := c.S.Send(&service.Request{
		Method: http.MethodGet,
		Params: map[string]string{"fields": strings.Join(fields, ",")},
	}, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Update(fields map[string]string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	err := c.S.Send(&service.Request{
		Method: http.MethodPatch,
		Data:   fields,
	}, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
