package models

import ()

type Send struct {
	To      []*Recipient `json:"to,omitempty"`
	CC      []*Recipient `json:"cc,omitempty"`
	BCC     []*Recipient `json:"bcc,omitempty"`
	Subject string       `json:"subject,omitempty"`
	HTML    string       `json:"html,omitempty"`
	From    string       `json:"from,omitempty"`
}

type Recipient struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
