package mixmax

import (
	"github.com/ghmeier/go-mixmax/appointmentlinks"
	"github.com/ghmeier/go-mixmax/availability"
	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-mixmax/codesnippet"
	"github.com/ghmeier/go-mixmax/contactgroups"
	"github.com/ghmeier/go-mixmax/contacts"
	"github.com/ghmeier/go-mixmax/events"
	"github.com/ghmeier/go-mixmax/filerequests"
	"github.com/ghmeier/go-mixmax/integrations"
	"github.com/ghmeier/go-mixmax/messages"
	"github.com/ghmeier/go-mixmax/polls"
	"github.com/ghmeier/go-mixmax/qa"
	"github.com/ghmeier/go-mixmax/reminders"
	"github.com/ghmeier/go-mixmax/rules"
	"github.com/ghmeier/go-mixmax/send"
	"github.com/ghmeier/go-mixmax/sequences"
)

type Client struct {
	AppointmentLinks *appointmentlinks.Client
	Availability     *availability.Client
	CodeSnippet      *codesnippet.Client
	Contacts         *contacts.Client
	ContactGroups    *contactgroups.Client
	Events           *events.Client
	FileRequests     *filerequests.Client
	Integrations     *integrations.Client
	Messages         *messages.Client
	Polls            *polls.Client
	QAs              *qa.Client
	Reminders        *reminders.Client
	Rules            *rules.Client
	Send             *send.Client
	Sequences        *sequences.Client
}

func New(key string) *Client {
	c := new(Client)
	client := client.New(key)
	c.AppointmentLinks = appointmentlinks.New(client)
	c.Availability = availability.New(client)
	c.CodeSnippet = codesnippet.New(client)
	c.Contacts = contacts.New(client)
	c.ContactGroups = contactgroups.New(client)
	c.Events = events.New(client)
	c.FileRequests = filerequests.New(client)
	c.Integrations = integrations.New(client)
	c.Messages = messages.New(client)
	c.Polls = polls.New(client)
	c.QAs = qa.New(client)
	c.Reminders = reminders.New(client)
	c.Rules = rules.New(client)
	c.Send = send.New(client)
	c.Sequences = sequences.New(client)
	return c
}
