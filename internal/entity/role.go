package entity

// Role represents an role record.
type Role struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

// TableName represents the table name
func (u Role) TableName() string {
	return "roles"
}
