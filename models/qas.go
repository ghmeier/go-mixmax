package models

import (
	"time"
)

type QAs struct {
	Results []*QA `json:"results"`
	*ResultList
}

type QA struct {
	ID           string      `json:"_id"`
	CreatorID    string      `json:"creatorId"`
	CreationDate time.Time   `json:"creationDate"`
	Title        string      `json:"title"`
	Questions    []*Question `json:"questions"`
}

type Question struct {
	Type    string    `json:"type"`
	Content string    `json:"content"`
	Answers []*Answer `json:"answers"`
}

type Answer struct {
	RecipientEmail string `json:"recipientEmail"`
	RecipientName  string `json:"recipientName"`
	Answer         string `json:"answer"`
}
