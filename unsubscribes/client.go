package unsubscribes

import (
	"net/http"
	"strconv"

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
			S:   c.S.Copy("unsubscribes"),
		},
	}
}

func (c *Client) List(options *models.SortQuery) (*models.Unsubscribes, error) {
	var res models.Unsubscribes
	err := c.S.Send(&service.Request{
		Method: http.MethodGet,
		Params: toParams(options),
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func toParams(options *models.SortQuery) map[string]string {
	return map[string]string{
		"sort":          options.Sort,
		"sortAscending": strconv.FormatBool(options.Ascending),
	}
}
