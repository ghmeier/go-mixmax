package events

import (
	"net/http"

	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-mixmax/models"
	"github.com/ghmeier/go-service"
)

type Client struct {
	*client.Client
}

func New(c *client.Client) *Client {
	return &Client{
		Client: &client.Client{
			Key: c.Key,
			S:   c.S.Copy("events"),
		},
	}
}

func (c *Client) List() (*models.Events, error) {
	var res models.Events
	err := c.S.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
