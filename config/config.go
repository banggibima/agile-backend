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
}

type App struct {
	Name    string
	Port    int
	Version string
}

type Postgres struct {
	Driver  string
	Url     string
	SSLMode string
}

type Redis struct {
	Url string
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

func Get(v *viper.Viper) (*Config, error) {
	config := &Config{
		App: App{
			Name:    v.GetString("APP_NAME"),
			Port:    v.GetInt("APP_PORT"),
			Version: v.GetString("APP_VERSION"),
		},
		Postgres: Postgres{
			Driver:  v.GetString("POSTGRES_DRIVER"),
			Url:     v.GetString("POSTGRES_URL"),
			SSLMode: v.GetString("POSTGRES_SSLMODE"),
		},
		Redis: Redis{
			Url: v.GetString("REDIS_URL"),
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
			Endpoint:        viper.GetString("MINIO_ENDPOINT"),
			AccessKeyID:     viper.GetString("MINIO_ACCESS_KEY_ID"),
			SecretAccessKey: viper.GetString("MINIO_SECRET_ACCESS_KEY"),
			UseSSL:          viper.GetBool("MINIO_USE_SSL"),
		},
	}

	return config, nil
}
