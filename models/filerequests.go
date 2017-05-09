package models

import (
	"time"
)

type FileRequests struct {
	Results []*FileRequest `json:"results"`
}

type FileRequest struct {
	ID        string    `json:"_id"`
	UserID    string    `json:"userId"`
	Timestamp time.Time `json:"timestamp"`
	Uploads   []*Upload `json:"uploads"`
}

type Upload struct {
	Timestamp   time.Time `json:"timestamp"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	IPAddress   string    `json:"ipAddress"`
	DownloadURL string    `json:"downloadUrl"`
}
