package rules

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
			S:   c.S.Copy("rules"),
		},
	}
}

func (c *Client) List() (*models.Rules, error) {
	var res models.Rules
	err := c.S.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) New(rule *models.RuleParams) (*models.Rule, error) {
	var res models.Rule
	err := c.S.Send(&service.Request{Method: http.MethodPost, Data: rule}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.Rule, error) {
	var res models.Rule
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Update(id string, params *models.RuleParams) (*models.Rule, error) {
	var res models.Rule
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodPatch, Data: params}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Delete(id string) error {
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodDelete}, nil)
	return err
}
