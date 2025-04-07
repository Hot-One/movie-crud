package postgres

import (
	"movie-crud/models"
	"movie-crud/storage/repo"
)

type MovieResository struct {
	db *Postgres
}

func NewMovieRepository(db *Postgres) repo.MovieRepoInterface {
	return &MovieResository{
		db: db,
	}
}

func (u *MovieResository) CreateMovie(movie *models.Movie) error {
	result := u.db.Db.Table("movies").
		Create(movie)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *MovieResository) GetMovieByID(id uint) (*models.Movie, error) {
	var movie models.Movie

	result := u.db.Db.Table("movies").
		First(&movie, id).
		Where("deleted_at IS NULL")

	if result.Error != nil {
		return nil, result.Error
	}

	return &movie, nil
}

func (u *MovieResository) UpdateMovie(movie *models.Movie) (*models.Movie, error) {

	result := u.db.Db.Table("movies").
		Save(movie)

	if result.Error != nil {
		return nil, result.Error
	}

	return u.GetMovieByID(movie.Id)
}

func (u *MovieResository) DeleteMovie(id uint) error {
	var movie models.Movie
	result := u.db.Db.Table("movies").
		Delete(&movie, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *MovieResository) GetAllMovies(request models.MovieGetListRequest) (*models.MovieGetListResponse, error) {
	var (
		movies []models.Movie
		total  int64
		where  = "deleted_at IS NULL"
	)

	result := u.db.Db.Table("movies").
		Where(where).
		Count(&total)
	if result.Error != nil {
		return nil, result.Error
	}

	result = u.db.Db.Table("movies").
		Where(where).
		Limit(int(request.Limit)).
		Offset(int(request.Offset)).
		Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}

	return &models.MovieGetListResponse{
		Total:    total,
		Response: movies,
	}, nil
}
