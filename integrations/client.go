package integrations

import (
	"net/http"

	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-mixmax/models"
	"github.com/ghmeier/go-service"
)

type Client struct {
	Key          string
	commands     service.Service
	enhancements service.Service
	linkresolver service.Service
	sidebar      service.Service
}

func New(c *client.Client) *Client {
	return &Client{
		Key:          c.Key,
		commands:     c.S.Copy("commands"),
		enhancements: c.S.Copy("enhancements"),
		linkresolver: c.S.Copy("linkresolver"),
		sidebar:      c.S.Copy("sidebar"),
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
	err := c.commands.Copy(id).Send(&service.Request{Method: http.MethodDelete}, nil)
	return err
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
	err := c.commands.Copy(id).Send(&service.Request{Method: http.MethodDelete}, nil)
	return nil, err
}
