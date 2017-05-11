package models

import (
	"time"
)

type SnippetTags struct {
	Results []*SnippetTag `json:"results"`
	*ResultList
}

type SnippetTag struct {
	ID        string      `json:"_id"`
	UserID    string      `json:"userId"`
	CreatedAt time.Time   `json:"createdAt"`
	SavedAt   time.Time   `json:"savedAt"`
	Name      string      `json:"name"`
	Count     int         `json:"count,omitempty"`
	Shared    interface{} `json:"shared"`
}

type SnippetTagFilterParams struct {
	Search string
	*ExpandQuery
	*SortQuery
}
