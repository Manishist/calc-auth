package models

import (
	"time"
)

type User struct {
	Id                    uint       `json:"id"`
	Name                  string     `json:"name"`
	Email                 string     `json:"email" gorm:"unique"`
	Password              []byte     `json:"-"`
	IsAdmin               bool       `json:"is_admin"`
	TotalTimeConsumed     uint64     `json:"total_time_consumed"` // in minutes
	TotalTimeToday        uint64     `json:"total_time_today"`    // in minutes
	LoggedInDaysLast7Days uint64     `json:"logged_in_days_last_7_days"`
	LastLogin             *time.Time `json:"-"`
	LastLoginDate         *time.Time `json:"-"`
}
