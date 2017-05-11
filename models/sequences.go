package models

import (
	"time"
)

type Sequences struct {
	Results []*Sequence `json:"results"`
	*ResultList
}

type Sequence struct {
	ID                  string               `json:"_id"`
	UserID              string               `json:"userId"`
	CreatedAt           time.Time            `json:"createdAt"`
	UpdatedAt           time.Time            `json:"updatedAt"`
	Name                string               `json:"name"`
	SendStagesAsReplies bool                 `json:"sendStagesAsReplies"`
	Stages              []string             `json:"stages"`
	CC                  []*SequenceRecipient `json:"cc"`
	BCC                 []*SequenceRecipient `json:"bcc"`
	LinkTracking        bool                 `json:"linkTrackingEnabled"`
	Notifications       bool                 `json:"notificationsEnabled"`
	TeamIDs             []string             `json:"teamIds"`
	Variables           []string             `json:"variables"`
	CRMsConnnected      []CRMConnect         `json:"crmConnected"`
	Timezone            string               `json:"timezone"`
	Shared              interface{}          `json:"shared"`
}

type SequenceFilterParams struct {
	Name   string
	Folder string
	*ExpandQuery
}

type SequenceCancel struct {
	Recipients []string `json:"recipients"`
}

type SequenceRecipients struct {
	Results []*SequenceRecipientDetails `json:"results"`
	*ResultList
}

type SequenceRecipientDetails struct {
	ID         string            `json:"_id"`
	UserID     string            `json:"userId"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
	SequenceID string            `json:"sequenceId"`
	Variables  map[string]string `json:"variables"`
	State      SequenceState     `json:"state"`
	*SequenceRecipient
}

type SequenceRecipientParams struct {
	Email     string            `json:"email"`
	Variables map[string]string `json:"variables"`
}

type SequenceRecipient struct {
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Status RecipientStatus `json:"status"`
	Errors []string        `json:"errors"`
}

type CRMConnect struct {
	Name  string `json:"name"`
	LogTo bool   `json:"logTo"`
}

type Folder struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Sent struct {
	Sent int `json:"sent"`
}

type RecipientStatus string

const (
	RECIPIENT_SUCCESS      = "success"
	RECIPIENT_ERROR        = "error"
	RECIPIENT_DUPLICATED   = "duplicated"
	RECIPIENT_UNSUBSCRIBED = "unsubscribed"
)

type SequenceState string

const (
	SEQUENCE_ACTIVE    SequenceState = "active"
	SEQUENCE_FAILED    SequenceState = "failed"
	SEQUENCE_COMPLETED SequenceState = "completed"
)
