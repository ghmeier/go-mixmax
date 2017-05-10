package models

import ()

type ResultList struct {
	Next       bool   `json:"hasNext,omitempty"`
	Previous   bool   `json:"hasPrevious,omitempty"`
	NextID     string `json:"next,omitempty"`
	PreviousID string `json:"previous,omitempty"`
}
