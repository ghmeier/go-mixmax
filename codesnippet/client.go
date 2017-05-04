package codesnippet

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

func (c *Client) List() (*models.CodeSnippets, error) {
	var res models.CodeSnippets
	err := c.S.Send(&service.Request{
		Method:  http.MethodGet,
		URL:     fmt.Sprintf("%s/codesnippets", c.URL),
		Headers: map[string]string{"X-API-Token": c.Key},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.CodeSnippet, error) {
	var res models.CodeSnippet
	err := c.S.Send(&service.Request{
		Method:  http.MethodGet,
		URL:     fmt.Sprintf("%s/codesnippets/%s", c.URL, id),
		Headers: map[string]string{"X-API-Token": c.Key},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
