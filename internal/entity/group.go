package entity

// Group represents a group.
type Group struct {
	ID       string `json:"id" db:"id"`
	CareerID string `json:"career_id" db:"career_id"`
	Name     string `json:"name" db:"name"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

// TableName represents the table name
func (u Group) TableName() string {
	return "groups"
}
