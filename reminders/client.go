package reminders

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
			S:   c.S.Copy("reminders"),
		},
	}
}

func (c *Client) List(options *models.ReminderFilterParams) (*models.Reminders, error) {
	params := map[string]string{
		"service":          options.Service,
		"serviceMessageId": options.ServiceMessageID,
		"serviceThreadId":  options.ServiceThreadID,
		"mixmaxMessageId":  options.MixmaxMessageID,
		"dismissed":        strconv.FormatBool(options.Dismissed),
		"limit":            strconv.Itoa(options.Limit),
	}

	var res models.Reminders
	err := c.S.Send(&service.Request{Method: http.MethodGet, Params: params}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) New(reminder *models.ReminderParams) (*models.Reminder, error) {
	var res models.Reminder
	err := c.S.Send(&service.Request{
		Method: http.MethodPost,
		Data:   map[string]interface{}{"reminder": reminder, "archiveThread": reminder.ArchiveThread},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string) (*models.Reminder, error) {
	var res models.Reminder
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Update(id string, params *models.ReminderParams) (*models.Reminder, error) {
	var res models.Reminder
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodPatch, Data: params}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Delete(id string) error {
	err := c.S.Copy(id).Send(&service.Request{Method: http.MethodDelete}, nil)
	return err
}
