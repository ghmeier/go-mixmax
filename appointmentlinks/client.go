package appointmentlinks

import (
	"fmt"
	"net/http"

	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-mixmax/models"
	"github.com/ghmeier/go-service"
)

type Client struct {
	*client.Client
}

func (c *Client) Me() (*models.AppointmentLinks, error) {
	var res models.AppointmentLinks
	err := c.S.Send(&service.Request{
		Method:  http.MethodGet,
		URL:     fmt.Sprintf("%s/appointmentlinks/me", c.URL),
		Headers: map[string]string{"X-API-Token": c.Key},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Set(name string) error {
	req := &models.AppointmentLinks{Name: name}
	err := c.S.Send(&service.Request{
		Method:  http.MethodPatch,
		URL:     fmt.Sprintf("%s/me", c.URL),
		Headers: map[string]string{"X-API-Token": c.Key},
		Data:    req,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Get(name string) (*models.AppointmentLinks, error) {
	var res models.AppointmentLinks
	err := c.S.Send(&service.Request{
		Method:  http.MethodGet,
		URL:     fmt.Sprintf("%s/appointmentlinks?name=%s", c.URL, name),
		Headers: map[string]string{"X-API-Token": c.Key},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
