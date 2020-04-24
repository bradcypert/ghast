package models

import "time"

// Timestamps is an embeddable struct to help add timestamps to Gorm models
type Timestamps struct {
	CreatedAt *time.Time
	DeletedAt *time.Time
	UpdatedAt *time.Tim
}
