package v1

import (
	"movie-crud/api/status_http"
	"movie-crud/models"
	"movie-crud/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Register user
// @Summary Register user
// @Description Register user
// @Router /v1/register [POST]
// @Tags Auth
// @Accept json
// @Produce json
// @Param movie body models.Register true "Register User"
// @Success 201
// @Response 400 {object} status_http.Response{data=string} "Bad request"
// @Failure 500 {object} status_http.Response{data=string} "Internal server error"
func (h *handlerV1) Register(c *gin.Context) {
	var request models.Register

	if err := c.ShouldBindJSON(&request); err != nil {
		h.handleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	if !utils.IsValidEmail(request.Email) {
		h.handleResponse(c, status_http.BadRequest, "invalid email")
		return
	}

	if !utils.IsValidPhone(request.Phone) {
		h.handleResponse(c, status_http.BadRequest, "invalid phone")
		return
	}

	token, err := h.storage.AuthService().Register(&request)
	if err != nil {
		h.handleResponse(c, status_http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, status_http.Created, token)
}
