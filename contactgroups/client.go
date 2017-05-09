package contactgroups

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
			S:   c.S.Copy("contactgroups"),
		},
	}
}

func (c *Client) List() (*models.ContactGroups, error) {
	var res models.ContactGroups
	err := c.S.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) New(name string, contacts []*models.Contact) (*models.ContactGroup, error) {
	var res models.ContactGroup
	err := c.S.Send(&service.Request{
		Method: http.MethodPost,
		Data:   map[string]interface{}{"name": name, "contacts": contacts},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.ContactGroup, error) {
	var res models.ContactGroup
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Update(id, name string) (*models.ContactGroup, error) {
	var res models.ContactGroup
	err := c.S.Copy(id).Send(&service.Request{
		Method: http.MethodPatch,
		Data:   map[string]interface{}{"name": name},
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

func (c *Client) Contacts(id string) (*models.Contacts, error) {
	var res models.Contacts
	err := c.S.Copy(id, "contacts").Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) AddContact(id string, params *models.ContactParams) (*models.Contact, error) {
	var res models.Contact
	err := c.S.Copy(id, "contacts").Send(&service.Request{
		Method: http.MethodPost,
		Data:   params,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteContact(id, contactID string) error {
	err := c.S.Copy(id, "contacts", contactID).Send(&service.Request{Method: http.MethodDelete}, nil)
	return err
}
