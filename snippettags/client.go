package snippettags

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
			S:   c.S.Copy("snippetags"),
		},
	}
}

func (c *Client) List(options *models.SnippetTagFilterParams) (*models.Snippets, error) {
	params := map[string]string{
		"sort":          options.Sort,
		"sortAscending": strconv.FormatBool(options.Ascending),
		"expand":        options.Expand,
		"search":        options.Search,
	}

	var res models.Snippets
	err := c.S.Send(&service.Request{Method: http.MethodGet, Params: params}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) New(name string) (*models.SnippetTag, error) {
	var res models.SnippetTag
	err := c.S.Send(&service.Request{
		Method: http.MethodPost,
		Data:   map[string]string{"name": name},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.SnippetTag, error) {
	var res models.SnippetTag
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Update(id, name string) (*models.SnippetTag, error) {
	var res models.SnippetTag
	err := c.S.Copy(id).Send(&service.Request{
		Method: http.MethodPatch,
		Data:   map[string]string{"name": name},
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

func (c *Client) Snippets(id string, options *models.SnippetTagFilterParams) (*models.Snippets, error) {
	params := map[string]string{
		"sort":          options.Sort,
		"sortAscending": strconv.FormatBool(options.Ascending),
		"expand":        options.Expand,
	}

	var res models.Snippets
	err := c.S.Copy(id, "snippets").Send(&service.Request{
		Method: http.MethodGet,
		Params: params,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewSnippet(id string, snippet *models.SnippetParams) (*models.SnippetTag, error) {
	var res models.SnippetTag
	err := c.S.Copy(id, "snippets").Send(&service.Request{
		Method: http.MethodPost,
		Data:   snippet,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
