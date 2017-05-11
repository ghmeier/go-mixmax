package models

import (
	"time"
)

type Commands struct {
	Results []*Command `json:"results"`
	*ResultList
}

type Command struct {
	ID                       string    `json:"_id"`
	UserID                   string    `json:"userId"`
	CreatedAt                time.Time `json:"createdAt"`
	Name                     string    `json:"name"`
	Commands                 []string  `json:"commands"`
	ParameterPlaceholder     string    `json:"parameterPlaceholder"`
	CommandResolveURL        string    `json:"commandResolveUrl"`
	CommandParameterHingsURL string    `json:"commandParameterHintsUrl"`
	Shared                   Shared    `json:"shared"`
}
