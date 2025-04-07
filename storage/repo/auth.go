package repo

import "movie-crud/models"

type AuthRepoInteface interface {
	Register(request *models.Register) (string, error)
	Login(request *models.Login) (string, error)
	HasAccess(request *models.HasAccessUserRequest) (*models.HasAccessUserResponse, error)
}
