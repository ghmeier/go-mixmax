package models

import (
	"time"
)

type Rules struct {
	Results []*Rules `json:"results"`
	*ResultList
}

type Rule struct {
	ID         string    `json:"_id"`
	UserID     string    `json:"userId"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Name       string    `json:"name"`
	EventName  EventName `json:"eventName"`
	Filter     string    `json:"filter,omitempty"`
	Action     *Action   `json:"action"`
	Paused     bool      `json:"isPaused"`
}

type RuleParams struct {
	Name      string    `json:"name,omitempty"`
	EventName EventName `json:"eventName,omitempty"`
	Filter    string    `json:"filter,omitempty"`
	Action    *Action   `json:"action,omitempty"`
	Paused    bool      `json:"isPaused,omitempty"`
}

type Action struct {
	Type            string `json:"type"`
	WebhookURL      string `json:"webhookUrl,omitempty"`
	AccountSID      string `json:"accountSid,omitempty"`
	AuthToken       string `json:"authToken,omitempty"`
	FromPhoneNumber string `json:"fromPhoneNumber,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
}

type EventName string

const (
	EVENT_UNOPENED          EventName = "unopened"
	EVENT_OPENED                      = "opened"
	EVENT_CLICKED                     = "clicked"
	EVENT_DOWNLOADED                  = "downloaded"
	EVENT_RECEIVED                    = "message:received"
	EVENT_VOTE                        = "poll:vote"
	EVENT_QA_SUBMIT                   = "qa:submit"
	EVENT_MEETING_CONFIRMED           = "meetinginvited:confirmed"
)
