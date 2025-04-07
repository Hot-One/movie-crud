package v1

import (
	"movie-crud/api/status_http"
	"movie-crud/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken = c.GetHeader("Authorization")

		if accessToken == "" {
			h.handleResponse(c, status_http.Unauthorized, "Access token is required")
			c.Abort()
			return
		}

		// Split the Bearer token
		tokenParts := strings.Split(accessToken, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			h.handleResponse(c, status_http.Unauthorized, "Invalid access token format")
			c.Abort()
			return
		}
		accessToken = tokenParts[1]

		response, err := h.storage.AuthService().HasAccess(&models.HasAccessUserRequest{
			Token: accessToken,
		})
		if err != nil {
			h.handleResponse(c, status_http.Unauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("user_id", response.UserId)
		c.Next()
	}
}
