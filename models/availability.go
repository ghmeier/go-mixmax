package models

import (
	"time"
)

type Availabilities struct {
	Results []*Availability `json:"results,omitempty"`
	Next    bool            `json:"hasNext"`
}
type Availability struct {
	ID             string     `json:"_id"`
	UserId         string     `json:"userId,omitempty"`
	Location       string     `json:"location,omitempty"`
	Description    string     `json:"description,omitempty"`
	Title          string     `json:"title,omitempty"`
	Timezone       string     `json:"timezone,omitempty"`
	CalendarID     string     `json:"calendarId,omitempty"`
	CalendarName   string     `json:"calendarName,omitempty"`
	Modify         bool       `json:"guestsCanModify,omitempty"`
	Timeslots      []Timeslot `json:"timeslots,omitempty"`
	DoubleBookings bool       `json:"allowDoubleBookings,omitempty"`
	Guests         []Guest    `json:"guests,omitempty"`
}

type Timeslot struct {
	Start  time.Time    `json:"start,omitempty"`
	End    time.Time    `json:"end,omitempty"`
	Events []TimedEvent `json:"events,omitempty"`
}

type TimedEvent struct {
	Guest Guest  `json:"guest,omitempty"`
	ID    string `json:"id,omitempty"`
}

type Guest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
