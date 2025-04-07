package v1

import (
	"movie-crud/api/status_http"
	"movie-crud/models"
	"movie-crud/pkg/logger"
	"movie-crud/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Create User
// @Summary Create User
// @Description Create User
// @Security BearerAuth
// @Router /v1/user [POST]
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserCreate true "Create User"
// @Success 201
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		h.handleResponse(c, status_http.BadRequest, err)
		return
	}

	if len(user.Username) < 6 {
		h.handleResponse(c, status_http.BadRequest, "username must be at least 6 characters")
		return
	}

	if !utils.IsValidPhone(user.Phone) {
		h.handleResponse(c, status_http.BadRequest, "invalid phone")
		return
	}

	if !utils.IsValidEmail(user.Email) {
		h.handleResponse(c, status_http.BadRequest, "invalid email")
		return
	}

	err := h.storage.UserService().CreateUser(&user)
	if err != nil {
		h.log.Error("failed while creating user", logger.Error(err))
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.Created, nil)
}

// Update User
// @Summary Update User
// @Description Update User
// @Router /v1/user [PUT]
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserUpdate true "Update User"
// @Success 200 {object} status_http.Response{data=string} "User data"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var request models.User

	if err := c.ShouldBindJSON(&request); err != nil {
		h.handleResponse(c, status_http.BadRequest, err)
		return
	}

	response, err := h.storage.UserService().UpdateUser(&request)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.OK, response)
}

// Get Single User By Id
// @Summary Get Single User
// @Description Get Single User By Id
// @Router /v1/user/{id} [GET]
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param id path uint true "user_id"
// @Success 200 {object} status_http.Response{data=string} "User data"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) GetSingleUser(c *gin.Context) {
	id := c.Param("id")
	userID := utils.StringToUint(id)

	response, err := h.storage.UserService().GetUserByID(userID)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.OK, response)
}

// Delete User
// @Summary Delete User
// @Description Delete User
// @Router /v1/user/{id} [DELETE]
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param id path uint true "user_id"
// @Success 204 {object} status_http.Response{data=string} "No Content"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID := utils.StringToUint(id)

	err := h.storage.UserService().DeleteUser(userID)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.NoContent, nil)
}

// Get All Users
// @Summary Get All Users
// @Description Get All Users
// @Router /v1/user [GET]
// @Security BearerAuth
// @Tags User
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} status_http.Response{data=string} "User data"
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) GetAllUsers(c *gin.Context) {
	var request models.UserGetListRequest

	request.Limit = h.getDefaultLimit(cast.ToInt32(c.Query("limit")))
	request.Offset = h.getDefaultOffset(cast.ToInt32(c.Query("offset")))

	response, err := h.storage.UserService().GetAllUsers(request)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.OK, response)
}
