package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Login     string    `json:"login"`
	FullName  string    `json:"full_name"`
	Gender    string    `json:"gender"`
	Age       int       `json:"age"`
	Contacts  string    `json:"contacts"`
	AvatarURL string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}
