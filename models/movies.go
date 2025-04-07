package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title    string `json:"title" gorm:"type:varchar(100);not null"`
	Director string `json:"director" gorm:"type:varchar(100);not null"`
	Year     string `json:"year" gorm:"type:varchar(4);not null"`
	Plot     string `json:"plot" gorm:"type:text;not null"`
}

type MovieCreate struct {
	Title    string `json:"title" binding:"required"`
	Director string `json:"director" binding:"required"`
	Year     string `json:"year" binding:"required"`
	Plot     string `json:"plot" binding:"required"`
}

type MovieUpdate struct {
	Id       uint   `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Director string `json:"director" binding:"required"`
	Year     string `json:"year" binding:"required"`
	Plot     string `json:"plot" binding:"required"`
}

type MovieGetListRequest struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type MovieGetListResponse struct {
	Response []Movie `json:"response"`
	Total    int64   `json:"total"`
}
