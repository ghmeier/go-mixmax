package teams

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
			S:   c.S.Copy("teams"),
		},
	}
}

func (c *Client) List(options *models.ExpandQuery) (*models.Teams, error) {
	var res models.Teams
	err := c.S.Send(&service.Request{
		Method: http.MethodGet,
		Params: toParams(options),
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) New(team *models.TeamParams) (*models.Team, error) {
	var res models.Team
	err := c.S.Send(&service.Request{
		Method: http.MethodPost,
		Data:   team,
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Get(id string, options *models.ExpandQuery) (*models.Team, error) {
	var res models.Team
	err := c.S.Copy(id).Send(&service.Request{
		Method: http.MethodGet,
		Params: toParams(options),
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Update(id, name string) (*models.Team, error) {
	var res models.Team
	err := c.S.Copy(id).Send(&service.Request{
		Method: http.MethodPost,
		Data:   map[string]string{"name": name},
	}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) Delete(id string) error {
	err := c.S.Send(&service.Request{Method: http.MethodDelete}, nil)
	return err
}

func (c *Client) Members(id string) (*models.Members, error) {
	var res models.Members
	err := c.S.Copy(id, "members").Send(&service.Request{Method: http.MethodGet}, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) NewMember(id string, members []*models.MemberParams) error {
	err := c.S.Copy(id, "members").Send(&service.Request{
		Method: http.MethodPost,
		Data:   map[string]interface{}{"members": members},
	}, nil)

	return err
}

func (c *Client) DeleteMember(id, memberID string) error {
	err := c.S.Copy(id, "members", memberID).
		Send(&service.Request{Method: http.MethodDelete}, nil)
	return err
}

func toParams(options *models.ExpandQuery) map[string]string {
	return map[string]string{"expand": options.Expand}
}
