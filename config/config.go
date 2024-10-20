package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Postgres Postgres
	Redis    Redis
	JWT      JWT
}

type App struct {
	Name  string
	Env   string
	Key   string
	Debug string
	Host  string
	Port  string
}

type Postgres struct {
	Driver   string
	URL      string
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

type Redis struct {
	URL      string
	Host     string
	Password string
	Port     string
}

type JWT struct {
	AccessSecret  string
	AccessExpire  string
	RefreshSecret string
	RefreshExpire string
	Audience      string
	Issuer        string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil && os.Getenv("APP_ENV") == "development" {
		return nil, err
	}

	config := &Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Env:   os.Getenv("APP_ENV"),
			Key:   os.Getenv("APP_KEY"),
			Debug: os.Getenv("APP_DEBUG"),
			Host:  os.Getenv("APP_HOST"),
			Port:  os.Getenv("APP_PORT"),
		},
		Postgres: Postgres{
			Driver:   os.Getenv("POSTGRES_DRIVER"),
			URL:      os.Getenv("POSTGRES_URL"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Username: os.Getenv("POSTGRES_USERNAME"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DATABASE"),
			SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		},
		Redis: Redis{
			URL:      os.Getenv("REDIS_URL"),
			Host:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Port:     os.Getenv("REDIS_PORT"),
		},
		JWT: JWT{
			AccessSecret:  os.Getenv("JWT_ACCESS_SECRET"),
			AccessExpire:  os.Getenv("JWT_ACCESS_EXPIRE"),
			RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
			RefreshExpire: os.Getenv("JWT_REFRESH_EXPIRE"),
			Audience:      os.Getenv("JWT_AUDIENCE"),
			Issuer:        os.Getenv("JWT_ISSUER"),
		},
	}

	return config, nil
}
