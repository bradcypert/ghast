package models

import "time"

type Timestamps struct {
	CreatedAt     *time.Time,
	DeletedAt     *time.Time,
	UpdatedAt	  *time.Time
}