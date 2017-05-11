package models

import (
	"time"
)

type Teams struct {
	Results []*Team `json:"results"`
	*ResultList
}

type Team struct {
	ID         string    `json:"_id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Name       string    `json:"name"`
	Members    []*Member `json:"members"`
}

type TeamParams struct {
	Name    string          `json:"name"`
	Members []*MemberParams `json:"members"`
}

type Members struct {
	Results []*Member `json:"members"`
	*ResultList
}

type Member struct {
	UserID     string    `json:"userId"`
	Accepted   bool      `json:"accepted,omitempty"`
	Email      string    `json:"email"`
	MemberID   string    `json:"memberId"`
	Role       Role      `json:"role"`
	CreatedAt  time.Time `json:"createdAt"`
	InvitedAt  time.Time `json:"invitedAt"`
	AcceptedAt time.Time `json:"acceptedAt"`
}

type MemberParams struct {
	Email string `json:"email"`
	Role  Role   `json:"role"`
}

type Role string

const (
	ROLE_ADMIN  Role = "admin"
	ROLE_MEMBER      = "member"
)
