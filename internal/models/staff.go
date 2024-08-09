package models

import (
	"time"
)

type Staff struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `gorm:"unique" json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	Clinic      *string   `json:"clinic"`
	Is_Admin    bool      `json:"is_admin"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
