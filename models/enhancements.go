package models

import (
	"time"
)

type Enhancements struct {
	Results []*Enhancement `json:"results"`
	*ResultList
}

type Enhancement struct {
	ID           string    `json:"_id"`
	UserID       string    `json:"userId"`
	CreatedAt    time.Time `json:"createdAt"`
	Name         string    `json:"name"`
	Icon         string    `json:"icon"`
	IconTooltip  string    `json:"iconTooltip"`
	EditorURL    string    `json:"editorUrl"`
	Commands     []string  `json:"commands"`
	ResolveURL   string    `json:"resolveUrl"`
	ActivateURL  string    `json:"activateUrl"`
	EditorHeight int       `json:"editorHeight"`
	EditorWidth  int       `json:"editorWidth"`
	Category     string    `json:"category"`
	Editable     bool      `json:"isEditable"`
	Shared       Shared    `json:"shared"`
}
