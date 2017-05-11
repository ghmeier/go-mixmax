package snippets

import (
	"net/http"
	"strconv"
	"strings"

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
			S:   c.S.Copy("snippets"),
		},
	}
}

func (c *Client) List() (*models.Snippets, error) {
	var res models.Snippets
	err := c.S.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string, options *models.SnippetFilterParams) (*models.Snippet, error) {
	params := map[string]string{
		"includeDeleted": strconv.FormatBool(options.IncludeDeleted),
		"expand":         options.Expand,
		"fields":         strings.Join(options.Fields, ","),
	}

	var res models.Snippet
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodGet, Params: params}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Update(id string, update *models.SnippetParams, options *models.SnippetFilterParams) (*models.Snippet, error) {
	params := map[string]string{
		"expand": options.Expand,
	}

	var res models.Snippet
	err := c.S.Copy(id).Send(&service.Request{
		Method: http.MethodPatch,
		Params: params,
		Data:   update,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Delete(id string) error {
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodDelete}, nil)
	return err
}

func (c *Client) Send(id string, send *models.SnippetSendParams) error {
	err := c.S.Copy(id, "send").Send(&service.Request{
		Method: http.MethodPost,
		Data:   send,
	}, nil)
	return err
}
