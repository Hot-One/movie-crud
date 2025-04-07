package repo

import "movie-crud/models"

type UserRepoInterface interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint) error
	GetAllUsers(request models.UserGetListRequest) (*models.UserGetListResponse, error)
}
