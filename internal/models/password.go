package models

import (
	"time"
)

type PasswordReset struct {
	ID        uint      `gorm:"primaryKey"`
	Phone     string    `gorm:"unique"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
