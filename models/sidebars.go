package models

import (
	"time"
)

type Sidebars struct {
	Results []*Sidebar `json:"results"`
	*ResultList
}

type Sidebar struct {
	ID        string    `json:"_id"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	IconHTML  string    `json:"iconHTML"`
}
