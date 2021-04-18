package entity

import "time"

// BuildingAccess represents a building access.
type BuildingAccess struct {
	ID          string     `json:"id" db:"id"`
	BuildingID  string     `json:"building_id" db:"building_id"`
	UserID      string     `json:"user_id" db:"user_id"`
	CheckIn     time.Time  `json:"check_in" db:"check_in"`
	CheckOut    *time.Time `json:"check_out" db:"check_out"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"update_at" db:"update_at"`
	User        User       `json:"user" db:"-"`
	Description *string    `json:"description" db:"description"`
}

// TableName represents the table name
func (u BuildingAccess) TableName() string {
	return "building_access"
}
