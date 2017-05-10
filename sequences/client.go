package sequences

import (
	"net/http"
	"time"

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
			S:   c.S.Copy("sequences"),
		},
	}
}

func (c *Client) List(options *models.SequenceFilterParams) (*models.Sequences, error) {
	params := map[string]string{
		"name":   options.Name,
		"expand": options.Expand,
		"folder": options.Folder,
	}

	var res models.Sequences
	err := c.S.Send(&service.Request{Method: http.MethodGet, Params: params}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Cancel(id string, emails []string) ([]string, error) {
	var cancelled models.SequenceCancel
	err := c.S.Copy(id, "cancel").Send(&service.Request{
		Method: http.MethodPost,
		Data:   map[string]interface{}{"emails": emails},
	}, &cancelled)

	if err != nil {
		return nil, err
	}

	return cancelled.Recipients, nil
}

func (c *Client) Recipients(id string) (*models.SequenceRecipients, error) {
	var res models.SequenceRecipients
	err := c.S.Copy(id, "recipients").Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewRecipients(id string, recipients []*models.SequenceRecipientParams, scheduled time.Time) ([]*models.SequenceRecipient, error) {
	res := make([]*models.SequenceRecipient, 0)
	err := c.S.Copy(id, "recipients").Send(&service.Request{
		Method: http.MethodPost,
		Data: map[string]interface{}{
			"recipients":  recipients,
			"scheduledAt": scheduled.Unix(),
		},
	}, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) CancelRecipient(id, recipientID string) error {
	err := c.S.Copy(id, "recipients", recipientID, "cancel").Send(&service.Request{
		Method: http.MethodPost,
	}, nil)
	return err
}

func (c *Client) Folders() ([]*models.Folder, error) {
	res := make([]*models.Folder, 0)
	err := c.S.Copy("folders").Send(&service.Request{Method: http.MethodGet}, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ListFolder(id string) (*models.Sequences, error) {
	var res models.Sequences
	err := c.S.Copy("folders", id, "sequences").Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Sent() (*models.Sent, error) {
	var res models.Sent
	err := c.S.Copy("sent").Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
