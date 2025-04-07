package api

import (
	v1 "movie-crud/api/handler/v1"
	"movie-crud/config"
	"movie-crud/pkg/logger"
	"movie-crud/storage"

	"movie-crud/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf    config.Config
	Logger  logger.Logger
	Storage storage.IStorage
}

// @title GromTemplate API
// @version 1.0
// @description API for Gorm Template
func NewRouter(option Option) *gin.Engine {
	docs.SwaggerInfo.Title = option.Conf.ServiceName
	docs.SwaggerInfo.Schemes = []string{option.Conf.HTTPScheme}

	router := gin.New()
	url := ginSwagger.URL("swagger/doc.json")

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:  option.Logger,
		Cfg:     option.Conf,
		Storage: option.Storage,
	})

	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.Use(customCORSMiddleware())
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/v1")

	api.POST("/register", handlerV1.Register)
	api.POST("/login", handlerV1.Login)

	user := api.Group("/user")
	{
		user.POST("", handlerV1.CreateUser)
		user.PUT("", handlerV1.UpdateUser)
		user.GET(":id", handlerV1.GetSingleUser)
		user.DELETE(":id", handlerV1.DeleteUser)
		user.GET("", handlerV1.GetAllUsers)
	}

	movie := api.Group("/movie")
	{
		movie.POST("", handlerV1.CreateMovie)
		movie.PUT("", handlerV1.UpdateMovie)
		movie.GET(":id", handlerV1.GetSingleMovie)
		movie.DELETE(":id", handlerV1.DeleteMovie)
		movie.GET("", handlerV1.GetAllMovies)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "3600")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
