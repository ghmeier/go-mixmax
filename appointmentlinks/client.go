package appointmentlinks

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
			S:   c.S.Copy("appointmentlinks"),
		},
	}
}

func (c *Client) Me() (*models.AppointmentLinks, error) {
	var res models.AppointmentLinks
	err := c.S.Copy("me").Send(&service.Request{
		Method: http.MethodGet,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Set(name string) error {
	req := &models.AppointmentLinks{Name: name}
	err := c.S.Copy("me").Send(&service.Request{
		Method: http.MethodPatch,
		Data:   req,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Get(name string) (*models.AppointmentLinks, error) {
	var res models.AppointmentLinks
	err := c.S.Send(&service.Request{
		Method: http.MethodGet,
		Params: map[string]string{"name": name},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
