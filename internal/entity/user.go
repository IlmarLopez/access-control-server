package entity

import "time"

// User represents a user.
type User struct {
	ID                 string     `json:"id" db:"id"`
	Username           string     `json:"username" db:"username"`
	Password           string     `json:"password,omitempty" db:"password"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at" db:"updated_at"`
	RoleID             string     `json:"role_id" db:"role_id"`
	RoleName           string     `json:"role_name" db:"role_name"`
	IsActive           bool       `json:"is_active" db:"is_active"`
	FirstName          string     `json:"first_name" db:"first_name"`
	LastName           string     `json:"last_name" db:"last_name"`
	Email              string     `json:"email" db:"email"`
	RegistrationNumber string     `json:"registration_number" db:"registration_number"`
	CareerID           string     `json:"career_id" db:"career_id"`
	CareerName         string     `json:"career_name" db:"career_name"`
	GroupID            string     `json:"group_id" db:"group_id"`
	GroupName          string     `json:"group_name" db:"group_name"`
}

// TableName represents the table name
func (u User) TableName() string {
	return "users"
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetRole returns the user name role.
func (u User) GetRoleID() string {
	return u.RoleID
}

// GetUsername returns the user username.
func (u User) GetUsername() string {
	return u.Username
}
