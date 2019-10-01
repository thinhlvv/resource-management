package model

import "time"

// Resource is data access model.
type Resource struct {
	ID        int        `json:"id"`
	UserID    int        `json:"-"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
