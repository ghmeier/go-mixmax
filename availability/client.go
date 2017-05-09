package availability

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
			S:   c.S.Copy("availability"),
		},
	}
}

func (c *Client) List() (*models.Availabilities, error) {
	var res models.Availabilities
	err := c.S.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.Availability, error) {
	var res models.Availability
	err := c.S.Copy(id).Send(&service.Request{
		Method: http.MethodGet,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
