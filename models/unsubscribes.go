package models

type Unsubscribes struct {
	Results []*Unsubscribe `json:"results"`
	*ResultList
}

type Unsubscribe struct {
	TeamID    string `json:"teamId,omitempty"`
	MessageID string `json:"messageId,omitempty"`
	*Contact
}
