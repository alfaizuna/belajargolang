package models

import "database/sql/driver"

type Role string

const (
	AdminRole   Role = "admin"
	StudentRole Role = "student"
)

// Scan implements the Scanner interface.
func (r *Role) Scan(value interface{}) error {
	*r = Role(value.(string))
	return nil
}

// Value implements the driver Valuer interface.
func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}
