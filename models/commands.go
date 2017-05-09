package models

import (
	"time"
)

type Commands struct {
	Results    []*Command `json:"results"`
	PreviousID string     `json:"previous"`
	NextID     string     `json:"next"`
	Previous   bool       `json:"previous"`
	Next       bool       `json:"next"`
}

type Command struct {
	ID                       string      `json:"_id"`
	UserID                   string      `json:"userId"`
	CreatedAt                time.Time   `json:"createdAt"`
	Name                     string      `json:"name"`
	Commands                 []string    `json:"commands"`
	ParameterPlaceholder     string      `json:"parameterPlaceholder"`
	CommandResolveURL        string      `json:"commandResolveUrl"`
	CommandParameterHingsURL string      `json:"commandParameterHintsUrl"`
	Shared                   interface{} `json:"shared"`
}
