package model

import "time"

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
	ID       int
	Quota    int
	Email    string
	Password string
	Role     UserRole

	DeletedAt *time.Time
}
