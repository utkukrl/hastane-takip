package models

import (
	"time"
)

type Province struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"unique;not null" json:"name"`
	Districts []District `gorm:"foreignKey:ProvinceID" json:"districts"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type District struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	ProvinceID uint      `json:"province_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
