package models

import ()

type CodeSnippets struct {
	Results []*CodeSnippet `json:"results,omitempty"`
	Next    bool           `json:"hasNext"`
}
type CodeSnippet struct {
	ID         string `json:"_id"`
	UserId     string `json:"userId,omitempty"`
	HTML       string `json:"html,omitempty"`
	Title      string `json:"title,omitempty"`
	Background string `json:"background,omitempty"`
	Code       string `json:"code,omitempty"`
	Theme      string `json:"theme,omitempty"`
	Language   string `json:"language,omitempty"`
}
