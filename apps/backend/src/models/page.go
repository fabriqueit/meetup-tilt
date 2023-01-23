package models

import "time"

// Page represents data about an Page.
type Page struct {
	ID        uint `json:"id" gorm:"primary_key"` // Primary key
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `json:"title"` // Page title
	Body      string `json:"body"`  // Page body
}
