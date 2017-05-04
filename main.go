package mixmax

import (
	"github.com/ghmeier/go-mixmax/appointmentlinks"
	"github.com/ghmeier/go-mixmax/availability"
	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-mixmax/codesnippet"
)

type Client struct {
	AppointmentLinks *appointmentlinks.Client
	Availability     *availability.Client
	CodeSnippet      *codesnippet.Client
}

func New(key string) *Client {
	c := new(Client)
	client := client.New(key)
	c.AppointmentLinks = &appointmentlinks.Client{Client: client}
	c.Availability = &availability.Client{Client: client}
	c.CodeSnippet = &codesnippet.Client{Client: client}
	return c
}
