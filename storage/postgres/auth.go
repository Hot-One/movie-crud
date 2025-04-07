package postgres

import (
	"movie-crud/config"
	"movie-crud/models"
	"movie-crud/pkg/security"
	"movie-crud/storage/repo"
	"time"
)

type AuthResository struct {
	cfg config.Config
	db  *Postgres
}

func NewAuthRepository(db *Postgres, cfg config.Config) repo.AuthRepoInteface {
	return &AuthResository{
		cfg: cfg,
		db:  db,
	}
}

func (r *AuthResository) Register(request *models.Register) (string, error) {
	user := models.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Phone:    request.Phone,
	}

	// Hash password
	hashedPassword, err := security.HashPasswordBcrypt(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	tx := r.db.Db.Begin()

	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		}
	}()

	// Create user
	if err := tx.Table("users").Create(&user).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	session := models.Session{
		UserId:    user.Id,
		ExpiresAt: time.Now().Add(config.AccessTokenExpiresInTime),
	}

	if err := tx.Table("sessions").Create(&session).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	tokenData := map[string]any{
		"id":         session.Id,
		"user_id":    session.UserId,
		"expires_at": session.ExpiresAt,
	}

	token, err := security.GenerateJWT(tokenData, config.AccessTokenExpiresInTime, r.cfg.SignKey)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return token, nil
}
