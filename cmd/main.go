package main

import (
	"movie-crud/api"
	"movie-crud/config"
	"movie-crud/pkg/logger"
	"movie-crud/storage"
	"movie-crud/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	var cfg = config.Load()
	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.New(cfg.ServiceName, *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	newPg := postgres.NewPostgres(&cfg)

	strg := storage.NewStorage(newPg)

	server := api.NewRouter(api.Option{
		Conf:    cfg,
		Logger:  log,
		Storage: strg,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
