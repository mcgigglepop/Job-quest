package models

import "time"

// Job represents a job application entry
type Job struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"` // Foreign key to track the owner
	Company     string    `json:"company"`
	Position    string    `json:"position"`
	Location    string    `json:"location"`
	Remote      string    `json:"remote"` // yes, no, hybrid
	Status      string    `json:"status"` // applied, interviewing, offer, rejected
	AppliedDate time.Time `json:"applied_date"`
	Notes       string    `json:"notes"`
	ResumeURL   string    `json:"resume_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"-"`
}
