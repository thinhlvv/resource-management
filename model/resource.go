package model

import "time"

// Resource is data access model.
type Resource struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
