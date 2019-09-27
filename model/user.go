package model

import "time"

// Quota ...
type Quota int

// Int converts type to int
func (q Quota) Int() int {
	return int(q)
}

// UnlimitedQuota marks user has unlimited to create resources.
var UnlimitedQuota Quota = -1

// UserRole stricts role of user's type.
type UserRole int

// Int convert type to int.
func (r UserRole) Int() int {
	return int(r)
}

// String returns name of role
func (r UserRole) String() string {
	return Roles[r]
}

var (
	// RoleUser ...
	RoleUser UserRole = 1
	// RoleAdmin ...
	RoleAdmin UserRole = 2
)

// Roles is mapping int and string
var Roles = map[UserRole]string{
	RoleUser:  "user",
	RoleAdmin: "admin",
}

// User is database access model.
type User struct {
	ID             int      `json:"id"`
	Quota          int      `json:"quota"`
	Email          string   `json:"email"`
	HashedPassword string   `json:"-"`
	Role           UserRole `json:"role"`

	DeletedAt *time.Time
}
