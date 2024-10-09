package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Postgres Postgres
	Redis    Redis
	JWT      JWT
	Minio    Minio
	Mongo    Mongo
	RabbitMQ RabbitMQ
}

type App struct {
	Name  string
	Env   string
	Key   string
	Debug bool
	Host  string
	Port  int
}

type Postgres struct {
	Driver   string
	URL      string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
}

type Redis struct {
	URL      string
	Host     string
	Password string
	Port     int
}

type JWT struct {
	AccessSecret  string
	AccessExpire  int
	RefreshSecret string
	RefreshExpire int
	Audience      string
	Issuer        string
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type Mongo struct {
	URL string
}

type RabbitMQ struct {
	URL string
}

func Load(v *viper.Viper) (*Config, error) {
	config := &Config{
		App: App{
			Name:  v.GetString("APP_NAME"),
			Env:   v.GetString("APP_ENV"),
			Key:   v.GetString("APP_KEY"),
			Debug: v.GetBool("APP_DEBUG"),
			Host:  v.GetString("APP_HOST"),
			Port:  v.GetInt("APP_PORT"),
		},
		Postgres: Postgres{
			Driver:   v.GetString("POSTGRES_DRIVER"),
			URL:      v.GetString("POSTGRES_URL"),
			Host:     v.GetString("POSTGRES_HOST"),
			Port:     v.GetInt("POSTGRES_PORT"),
			Username: v.GetString("POSTGRES_USERNAME"),
			Password: v.GetString("POSTGRES_PASSWORD"),
			Database: v.GetString("POSTGRES_DATABASE"),
			SSLMode:  v.GetString("POSTGRES_SSLMODE"),
		},
		Redis: Redis{
			URL:      v.GetString("REDIS_URL"),
			Host:     v.GetString("REDIS_HOST"),
			Password: v.GetString("REDIS_PASSWORD"),
			Port:     v.GetInt("REDIS_PORT"),
		},
		JWT: JWT{
			AccessSecret:  v.GetString("JWT_ACCESS_SECRET"),
			AccessExpire:  v.GetInt("JWT_ACCESS_EXPIRE"),
			RefreshSecret: v.GetString("JWT_REFRESH_SECRET"),
			RefreshExpire: v.GetInt("JWT_REFRESH_EXPIRE"),
			Audience:      v.GetString("JWT_AUDIENCE"),
			Issuer:        v.GetString("JWT_ISSUER"),
		},
		Minio: Minio{
			Endpoint:        v.GetString("MINIO_ENDPOINT"),
			AccessKeyID:     v.GetString("MINIO_ACCESS_KEY_ID"),
			SecretAccessKey: v.GetString("MINIO_SECRET_ACCESS_KEY"),
			UseSSL:          v.GetBool("MINIO_USE_SSL"),
		},
		Mongo: Mongo{
			URL: v.GetString("MONGO_URL"),
		},
		RabbitMQ: RabbitMQ{
			URL: v.GetString("RABBITMQ_URL"),
		},
	}

	return config, nil
}
