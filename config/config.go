package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServiceName string
	Environment string
	HTTPPort    string
	HTTPScheme  string

	SignKey string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

func Load() Config {
	if err := godotenv.Load("app/.env"); err != nil {
		if err := godotenv.Load(".env"); err != nil {
			fmt.Println("No .env file found")
		}
		fmt.Println("No /app/.env file found")
	}

	c := Config{}
	c.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "movie-crud"))
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.HTTPScheme = cast.ToString(getOrReturnDefault("HTTP_SCHEME", "http"))

	c.SignKey = cast.ToString(getOrReturnDefault("SIGN_KEY", ""))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "movie-crud"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "movie-crud"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "movie-crud"))

	return c
}

func getOrReturnDefault(key string, defaultValue any) any {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
