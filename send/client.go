package send

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
			S:   c.S.Copy("send"),
		},
	}
}

func (c *Client) New(send *models.Send) error {
	err := c.S.Send(&service.Request{Method: http.MethodPost, Data: send}, nil)
	return err
}
