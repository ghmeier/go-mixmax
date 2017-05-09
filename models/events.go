package models

import (
	"time"
)

type Events struct {
	Results []*Event `json:"results"`
	NextID  string   `json:"next"`
	Next    bool     `json:"hasNext"`
}

type Event struct {
	ID             string     `json:"_id"`
	ThreadID       string     `json:"threadId"`
	Subject        string     `json:"subject"`
	Timestamp      time.Time  `json:"timestamp"`
	RecipientEmail string     `json:"recipientEmail"`
	RecipientName  string     `json:"recipientName"`
	UserID         string     `json:"userId"`
	Device         *Device    `json:"device"`
	Location       *Location  `json:"location"`
	Object         *Object    `json:"object"`
	Action         string     `json:"action"`
	Recipients     []*Contact `json:"recipients"`
}

type Device struct {
	OS      string `json:"os"`
	Program string `json:"program"`
	Mobile  bool   `json:"isMobile"`
}

type Object struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type Location struct {
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
}
