package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Id        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
}
