package models

import (
	"time"
)

type Polls struct {
	Results []*Poll `json:"results"`
	*ResultList
}

type Poll struct {
	ID      string    `json:"_id"`
	UserID  string    `json:"userId"`
	Created time.Time `json:"created"`
	Live    bool      `json:"livePoll"`
	Options []*Option `json:"options"`
}

type Option struct {
	Text        string        `json:"text"`
	Respondents []*Respondent `json:"respondents"`
}

type Respondent struct {
	*Contact
	RespondedAt time.Time `json:"respondedAt"`
}
