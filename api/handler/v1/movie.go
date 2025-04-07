package v1

import (
	"movie-crud/api/status_http"
	"movie-crud/models"
	"movie-crud/pkg/logger"
	"movie-crud/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Create Movie
// @Summary Create Movie
// @Description Create Movie
// @Router /v1/movie [POST]
// @Security BearerAuth
// @Tags Movie
// @Accept json
// @Produce json
// @Param movie body models.MovieCreate true "Movie User"
// @Success 201
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) CreateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		h.handleResponse(c, status_http.BadRequest, err)
		return
	}

	err := h.storage.MovieService().CreateMovie(&movie)
	if err != nil {
		h.log.Error("failed while creating user", logger.Error(err))
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.Created, nil)
}

// Update Movie
// @Summary Update Movie
// @Description Update Movie
// @Router /v1/movie [PUT]
// @Security BearerAuth
// @Tags Movie
// @Accept json
// @Produce json
// @Param movie body models.MovieUpdate true "Update Movie"
// @Success 200 {object} status_http.Response{data=string} "Movie data"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) UpdateMovie(c *gin.Context) {
	var request models.Movie

	if err := c.ShouldBindJSON(&request); err != nil {
		h.handleResponse(c, status_http.BadRequest, err)
		return
	}

	response, err := h.storage.MovieService().UpdateMovie(&request)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.OK, response)
}

// Get Single Movie By Id
// @Summary Get Single Movie
// @Description Get Single Movie By Id
// @Router /v1/movie/{id} [GET]
// @Security BearerAuth
// @Tags Movie
// @Accept json
// @Produce json
// @Param id path uint true "movie_id"
// @Success 200 {object} status_http.Response{data=string} "Movie data"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) GetSingleMovie(c *gin.Context) {
	id := c.Param("id")
	movieID := utils.StringToUint(id)

	response, err := h.storage.MovieService().GetMovieByID(movieID)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.OK, response)
}

// Delete Movie
// @Summary Delete Movie
// @Description Delete Movie
// @Router /v1/movie/{id} [DELETE]
// @Security BearerAuth
// @Tags Movie
// @Accept json
// @Produce json
// @Param id path uint true "movie_id"
// @Success 204 {object} status_http.Response{data=string} "No Content"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	movieID := utils.StringToUint(id)

	err := h.storage.MovieService().DeleteMovie(movieID)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.NoContent, nil)
}

// Get All Movies
// @Summary Get All Movies
// @Description Get All Movies
// @Router /v1/movie [GET]
// @Security BearerAuth
// @Tags Movie
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} status_http.Response{data=string} "Movie data"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) GetAllMovies(c *gin.Context) {
	var request models.MovieGetListRequest

	request.Limit = h.getDefaultLimit(cast.ToInt32(c.Query("limit")))
	request.Offset = h.getDefaultOffset(cast.ToInt32(c.Query("offset")))

	response, err := h.storage.MovieService().GetAllMovies(request)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.OK, response)
}
