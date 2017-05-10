package models

import (
	"time"
)

type LinkResolvers struct {
	Results []*LinkResolver `json:"results"`
	*ResultList
}

type LinkResolver struct {
	ID          string    `json:"_id"`
	UserID      string    `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
	ResolveURL  string    `json:"resolveUrl"`
	Regex       string    `json:"regex"`
	Order       int       `json:"order"`
}
