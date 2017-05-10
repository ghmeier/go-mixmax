package models

import (
	"time"
)

type Snippets struct {
	Results []*Snippet `json:"results"`
	*ResultList
}

type Snippet struct {
	ID        string      `json:"_id"`
	UserID    string      `json:"userId"`
	CreatedAt time.Time   `json:"createdAt"`
	SavedAt   time.Time   `json:"savedAt"`
	DeletedAt time.Time   `json:"deletedAt"`
	SavedBy   string      `json:"savedBy"`
	Name      string      `json:"name"`
	Title     string      `json:"title"`
	Source    string      `json:"source"`
	Shared    interface{} `json:"shared"`
}

type SnippetParams struct {
	Title  string `json:"title,omitempty"`
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}

type SnippetFilterParams struct {
	IncludeDeleted bool
	Expand         string
	Fields         []string
}

type SnippetSendParams struct {
	To        []*Recipient      `json:"to"`
	CC        []*Recipient      `json:"cc"`
	BCC       []*Recipient      `json:"bcc"`
	Subject   string            `json:"subject"`
	Tracking  bool              `json:"trackingEnabled"`
	Variables map[string]string `json:"variables"`
}
