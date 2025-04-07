package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone    string `json:"phone" gorm:"type:varchar(100);unique;not null"`
	Age      int    `json:"age" gorm:"type:int"`
}

type UserCreate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Age      int    `json:"age"`
}

type UserUpdate struct {
	Id       uint   `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Age      int    `json:"age"`
}

type UserGetListRequest struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type UserGetListResponse struct {
	Response []User `json:"response"`
	Total    int64  `json:"total"`
}
