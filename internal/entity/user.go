package entity

import "time"

// User represents a user.
type User struct {
	ID        string     `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password,omitempty" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	RoleID    string     `json:"role_id" db:"role_id"`
	IsActive  bool       `json:"is_active" db:"is_active"`
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	// Permissions []Permission `json:"permissions" db:"-"`
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
