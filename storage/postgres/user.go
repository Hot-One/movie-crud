package postgres

import (
	"movie-crud/models"
	"movie-crud/pkg/security"
	"movie-crud/storage/repo"
)

type UserResository struct {
	db *Postgres
}

func NewUserRepository(db *Postgres) repo.UserRepoInterface {
	return &UserResository{
		db: db,
	}
}

func (u *UserResository) CreateUser(user *models.User) error {
	hashedPassword, err := security.HashPasswordBcrypt(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	result := u.db.Db.Table("users").
		Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *UserResository) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	result := u.db.Db.Table("users").
		First(&user, id).
		Where("deleted_at IS NULL")

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserResository) UpdateUser(user *models.User) (*models.User, error) {
	if len(user.Password) != 60 {
		hashedPassword, err := security.HashPasswordBcrypt(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	result := u.db.Db.Table("users").
		Save(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return u.GetUserByID(user.Id)
}

func (u *UserResository) DeleteUser(id uint) error {
	var user models.User
	result := u.db.Db.Table("users").
		Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserResository) GetAllUsers(request models.UserGetListRequest) (*models.UserGetListResponse, error) {
	var (
		users []models.User
		total int64
		where = "deleted_at IS NULL"
	)

	result := u.db.Db.Table("users").
		Where(where).
		Count(&total)
	if result.Error != nil {
		return nil, result.Error
	}

	result = u.db.Db.Table("users").
		Where(where).
		Limit(int(request.Limit)).
		Offset(int(request.Offset)).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.UserGetListResponse{
		Total:    total,
		Response: users,
	}, nil
}
