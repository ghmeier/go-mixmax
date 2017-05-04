package availability

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

func (c *Client) List() (*models.Availabilities, error) {
	var res models.Availabilities
	err := c.S.Send(&service.Request{
		Method:  http.MethodGet,
		URL:     fmt.Sprintf("%s/availability", c.URL),
		Headers: map[string]string{"X-API-Token": c.Key},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.Availability, error) {
	var res models.Availability
	err := c.S.Send(&service.Request{
		Method:  http.MethodGet,
		URL:     fmt.Sprintf("%s/availability/%s", c.URL, id),
		Headers: map[string]string{"X-API-Token": c.Key},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
