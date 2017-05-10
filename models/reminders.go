package models

import (
	"time"
)

type Reminders struct {
	Results []*Reminder `json:"results"`
	*ResultList
}

type Reminder struct {
	ID               string    `json:"_id,omitempty"`
	UserID           string    `json:"userId,omitempty"`
	Date             time.Time `json:"date"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	Dismissed        bool      `json:"dismissed,omitempty"`
	Trigger          string    `json:"trigger"`
	ThreadSubject    string    `json:"threadSubject,omitempty"`
	Description      string    `json:"description,omitempty"`
	Service          string    `json:"service"`
	ServiceMessageID string    `json:"serviceMessageId,omitempty"`
	ServiceThreadID  string    `json:"serviceThreadId,omitempty"`
	MixmaxMessageID  string    `json:"mixmaxMessageId,omitempty"`
}

type ReminderParams struct {
	Date             time.Time  `json:"date,omitempty"`
	Trigger          string     `json:"trigger,omitempty"`
	ThreadSubject    string     `json:"threadSubject,omitempty"`
	Description      string     `json:"description,omitempty"`
	Service          string     `json:"service,omitempty"`
	ServiceMessageID string     `json:"serviceMessageId,omitempty"`
	ServiceThreadID  string     `json:"serviceThreadId,omitempty"`
	MixmaxMessageID  string     `json:"mixmaxMessageId,omitempty"`
	Recipients       []*Contact `json:"recipients,omitempty"`
	Dismissed        bool       `json:"dismissed,omitempty"`
	ArchiveThread    bool
}

type ReminderFilterParams struct {
	Service          string
	ServiceMessageID string
	ServiceThreadID  string
	MixmaxMessageID  string
	Dismissed        bool
	Limit            int
}
