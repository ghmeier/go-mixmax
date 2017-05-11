package models

import (
	"time"
)

type YesNos struct {
	Results []*YesNo `json:"results"`
	*ResultList
}

type YesNo struct {
	ID      string    `json:"_id"`
	UserID  string    `json:"userId"`
	Created time.Time `json:"created"`
}
