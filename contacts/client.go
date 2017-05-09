package contacts

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
			S:   c.S.Copy("contacts"),
		},
	}
}

func (c *Client) List() (*models.Contacts, error) {
	//TODO: sorting.
	var res models.Contacts
	err := c.S.Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) New(contacts ...*models.ContactParams) error {
	err := c.S.Send(&service.Request{
		Method:  http.MethodPost,
		Headers: map[string]string{"X-API-Token": c.Key},
		Data:    map[string][]*models.ContactParams{"contacts": contacts},
	}, nil)

	return err
}

func (c *Client) Contact(id string) (*models.Contact, error) {
	var res models.Contact
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Update(id string, params *models.ContactParams) error {
	err := c.S.Copy(id).Send(&service.Request{
		Method: http.MethodPatch,
		Data:   map[string]*models.ContactParams{"contact": params},
	}, nil)

	return err
}

func (c *Client) Delete(id string) error {
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodDelete}, nil)

	return err
}

func (c *Client) Notes(id string) (*models.Notes, error) {
	var res models.Notes
	// TODO: populate expand field
	err := c.S.Copy(id, "notes").Send(&service.Request{
		Method: http.MethodGet,
		Params: map[string]string{"expand": ""},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewNote(id, text string) error {
	err := c.S.Copy(id, "notes").Send(&service.Request{
		Method: http.MethodPost,
		Params: map[string]string{"expand": ""},
		Data:   map[string]string{"text": text},
	}, nil)
	return err
}

func (c *Client) UpdateNote(id, noteId, text string) error {
	err := c.S.Copy(id, "notes", noteId).Send(&service.Request{
		Method: http.MethodPatch,
		Params: map[string]string{"expand": ""},
		Data:   map[string]string{"text": text},
	}, nil)
	return err
}

func (c *Client) DeleteNote(id, noteId string) error {
	err := c.S.Copy(id, "notes", noteId).Send(&service.Request{
		Method: http.MethodDelete,
	}, nil)
	return err
}
