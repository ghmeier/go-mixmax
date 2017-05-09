package codesnippet

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
			S:   c.S.Copy("codesnippets"),
		},
	}
}

func (c *Client) List() (*models.CodeSnippets, error) {
	var res models.CodeSnippets
	err := c.S.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.CodeSnippet, error) {
	var res models.CodeSnippet
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) New(p *models.CodeSnippetParams) (string, error) {
	var snippet models.CodeSnippetResult
	err := c.S.Send(&service.Request{
		Method: http.MethodPost,
		Data:   p,
	}, &snippet)

	return snippet.ID, err
}

func (c *Client) Update(s *models.CodeSnippet) (string, error) {
	var snippet models.CodeSnippetResult
	err := c.S.Copy(s.ID).Send(&service.Request{
		Method: http.MethodPatch,
		Data: &models.CodeSnippetParams{
			HTML:       s.HTML,
			Title:      s.Title,
			Background: s.Background,
			Theme:      s.Theme,
			Language:   s.Language,
		},
	}, &snippet)

	return snippet.ID, err
}
