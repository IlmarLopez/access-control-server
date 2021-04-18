package entity

import "time"

// Building represents a building.
type Building struct {
	ID          string     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	UserLimit   int        `json:"user_limit" db:"user_limit"`
	ActiveUsers int        `json:"active_users" db:"active_users"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	IsActive    bool       `json:"is_active" db:"is_active"`
}

// TableName represents the table name
func (u Building) TableName() string {
	return "buildings"
}
