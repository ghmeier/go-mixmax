package integrations

import (
	"net/http"

	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-mixmax/models"
	"github.com/ghmeier/go-service"
)

type Client struct {
	Key           string
	commands      service.Service
	enhancements  service.Service
	linkresolvers service.Service
	sidebars      service.Service
}

func New(c *client.Client) *Client {
	return &Client{
		Key:           c.Key,
		commands:      c.S.Copy("commands"),
		enhancements:  c.S.Copy("enhancements"),
		linkresolvers: c.S.Copy("linkresolvers"),
		sidebars:      c.S.Copy("sidebars"),
	}
}

func (c *Client) Commands() (*models.Commands, error) {
	var res models.Commands
	err := c.commands.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewCommand(command *models.Command) (*models.Command, error) {
	var res models.Command
	err := c.commands.Send(&service.Request{Method: http.MethodPost, Data: command}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Command(id string) (*models.Command, error) {
	var res models.Command
	err := c.commands.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCommand(id string) error {
	return c.delete(c.commands, id)
}

func (c *Client) Enhancements() (*models.Enhancements, error) {
	var res models.Enhancements
	err := c.enhancements.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewEnhancement(enhancement *models.Enhancement) (*models.Enhancement, error) {
	var res models.Enhancement
	err := c.enhancements.Send(&service.Request{Method: http.MethodPost, Data: enhancement}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Enhancement(id string) (*models.Enhancement, error) {
	var res models.Enhancement
	err := c.enhancements.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteEnhancement(id string) error {
	return c.delete(c.enhancements, id)
}

func (c *Client) LinkResolvers() (*models.LinkResolvers, error) {
	var res models.LinkResolvers
	err := c.linkresolvers.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewLinkResolver(resolver *models.LinkResolver) (*models.LinkResolver, error) {
	var res models.LinkResolver
	err := c.linkresolvers.Send(&service.Request{Method: http.MethodPost, Data: resolver}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) LinkResolver(id string) (*models.LinkResolver, error) {
	var res models.LinkResolver
	err := c.linkresolvers.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteLinkResolver(id string) error {
	return c.delete(c.linkresolvers, id)
}

func (c *Client) Sidebars() (*models.Sidebars, error) {
	var res models.Sidebars
	err := c.sidebars.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewSidebar(sidebar *models.Sidebar) (*models.Sidebar, error) {
	var res models.Sidebar
	err := c.sidebars.Send(&service.Request{Method: http.MethodPost, Data: sidebar}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Sidebar(id string) (*models.Sidebar, error) {
	var res models.Sidebar
	err := c.sidebars.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteSidebar(id string) error {
	return c.delete(c.sidebars, id)
}

func (c *Client) delete(s service.Service, id string) error {
	return s.Copy(id).Send(&service.Request{Method: http.MethodDelete}, nil)
}
