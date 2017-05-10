package models

import (
	"time"
)

type Messages struct {
	Results []*Message `json:"results"`
	*ResultList
}

type Message struct {
	ID            string        `json:"_id"`
	UserID        string        `json:"userId"`
	Created       time.Time     `json:"created"`
	Updated       time.Time     `json:"updated"`
	From          *Contact      `json:"from"`
	To            *Contact      `json:"to"`
	CC            []*Contact    `json:"cc"`
	BCC           []*Contact    `json:"bcc"`
	Subject       string        `json:"subject"`
	Tracking      bool          `json:"trackingEnabled"`
	LinkTracking  bool          `json:"linkTrackingEnabled"`
	Notifications bool          `json:"notificationsEnabled"`
	TeamIDs       []string      `json:"teamIds"`
	FollowUp      *FollowUp     `json:"followUp"`
	Attachments   []*Attachment `json:"attachments"`
	SnippetIDs    []string      `json:"snippetIds"`
	Sequence      *Sequence     `json:"sequence"`
	Body          string        `json:"body"`
	Scheduled     time.Time     `json:"scheduled"`
	Sent          time.Time     `json:"sent"`
	SendingError  string        `json:"sendingError"`
	ErrorResponse string        `json:"errorResponse"`
}

type FollowUp struct {
	UnparsedDate string `json:"unparsedDate"`
	Trigger      string `json:"trigger"`
	Description  string `json:"description"`
}

type Attachment struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`
}

type Sequence struct {
	ID          string `json:"id"`
	StageID     string `json:"stageId"`
	RecipiendID string `json:"recipientId"`
}
