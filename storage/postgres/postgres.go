package postgres

import (
	"fmt"
	"log"
	"movie-crud/config"
	"movie-crud/models"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func NewPostgres(cfg *config.Config) *Postgres {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
		cfg.PostgresPort,
	)

	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.Session{})

	return &Postgres{Db: db}
}
