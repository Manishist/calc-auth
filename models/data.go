package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

// JSONStringArray is a custom type to handle []string as JSON
type JSONStringArray []string

// Scan implements the sql.Scanner interface for JSONStringArray
func (a *JSONStringArray) Scan(value interface{}) error {
	if value == nil {
		*a = JSONStringArray{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, a)
}

// Value implements the driver.Valuer interface for JSONStringArray
func (a JSONStringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "[]", nil
	}
	return json.Marshal(a)
}

// Data model with the JSONStringArray custom type for the History field
type Data struct {
	gorm.Model
	UserEmail   string          `json:"user_email" gorm:"index"`
	UserName    string          `json:"user_name"`
	ProjectName string          `json:"project_name"`
	History     JSONStringArray `json:"history" gorm:"type:json"`
}
