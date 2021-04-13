package entity

// Career represents a career.
type Career struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	IsActive bool   `json:"is_active" db:"is_active"`
}

// TableName represents the table name
func (u Career) TableName() string {
	return "careers"
}
