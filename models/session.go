package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Id        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"not null;unique"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
}
