package models

import ()

type ResultList struct {
	Next       bool   `json:"hasNext,omitempty"`
	Previous   bool   `json:"hasPrevious,omitempty"`
	NextID     string `json:"next,omitempty"`
	PreviousID string `json:"previous,omitempty"`
}

type ExpandQuery struct {
	Expand string
}

type SortQuery struct {
	Sort      string
	Ascending bool
}

type User struct {
	ID string `json:"_id"`
}

type SharedUser struct {
	UserID   string     `json:"userId"`
	Email    string     `json:"email"`
	Role     SharedRole `json:"role"`
	Accepted bool       `json:"accepted"`
}

type SharedTeam struct {
	ID   string     `json:"_id"`
	Role SharedRole `json:"role"`
}

type Shared struct {
	Users []*SharedUser `json:"users"`
	Teams []*SharedTeam `json:"teams"`
}

type SharedRole string

const (
	SHARED_READ  SharedRole = "read-only"
	SHARED_WRITE            = "read-write"
)
