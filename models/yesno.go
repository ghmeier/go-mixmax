package models

import (
	"time"
)

type YesNos struct {
	Results []*YesNo `json:"results"`
	*ResultList
}

type YesNo struct {
	ID      string      `json:"_id"`
	UserID  string      `json:"userId"`
	Created time.Time   `json:"created"`
	Text    string      `json:"text"`
	Answers []*Response `json:"answers"`
}

type Response struct {
	Text        string   `json:"text"`
	Respondents []string `json:"respondents"`
}
