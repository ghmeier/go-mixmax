package models

import (
	"time"
)

type Contacts struct {
	Results []*Contact `json:"results"`
	*ResultList
}

type Contact struct {
	ID        string            `json:"_id,omitempty"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	UserID    string            `json:"userId,omitempty"`
	ContactID string            `json:"contactID,omitempty"`
	Analytics *ContactAnalytics `json:"analytics,omitempty"`
}

type ContactParams struct {
	Email        string      `json:"email"`
	Name         string      `json:"name"`
	Groups       []string    `json:"groups"`
	Meta         interface{} `json:"meta"`
	ContactID    string      `json:"contactId"`
	SalseforceID string      `json:"salesforceId"`
	MarkAsUsed   bool        `json:"markAsUsed"`
	Enrich       bool        `json:"enrich"`
}

type ContactAnalytics struct {
	Sent                      int                 `json:"sent"`
	Opened                    int                 `json:"opened"`
	Clicked                   int                 `json:"clicked"`
	Downloaded                int                 `json:"downloaded"`
	Replied                   int                 `json:"replied"`
	Delivered                 int                 `json:"delivered"`
	DeliveredWithTrackedOpens int                 `json:"deliveredWithTrackedOpens"`
	DeliveredWithTrackedLinks int                 `json:"deliveredWithTrackedLinks"`
	DeliveredWithTrackedFiles int                 `json:"deliveredWithTrackedFiles"`
	Percentages               *ContactPercentages `json:"percentages"`
}

type ContactPercentages struct {
	Opened     float64 `json:"opened"`
	Clicked    float64 `json:"clicked"`
	Downloaded float64 `json:"downloaded"`
	Replied    float64 `json:"replied"`
}

type Notes struct {
	Results []*Note `json:"results"`
	*ResultList
}

type Note struct {
	ID         string    `json:"_id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	UserID     string    `json:"userId"`
	ContactID  string    `json:"contactId"`
	Text       string    `json:"text"`
}
