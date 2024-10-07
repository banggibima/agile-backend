package main

import (
	"github.com/banggibima/agile-backend/config"
	"github.com/banggibima/agile-backend/internal/transport/http"
	"github.com/banggibima/agile-backend/pkg/echo"
	"github.com/banggibima/agile-backend/pkg/logrus"
	"github.com/banggibima/agile-backend/pkg/postgres"
	"github.com/banggibima/agile-backend/pkg/viper"
)

func main() {
	v, err := viper.Init()
	if err != nil {
		panic(err)
	}

	config, err := config.Get(v)
	if err != nil {
		panic(err)
	}

	logrus, err := logrus.Init(config)
	if err != nil {
		panic(err)
	}

	echo, err := echo.Init(config)
	if err != nil {
		panic(err)
	}

	postgres, err := postgres.Client(config)
	if err != nil {
		panic(err)
	}

	if err := http.NewHTTP(
		config,
		echo,
		logrus,
		postgres,
	).Start(); err != nil {
		panic(err)
	}
}
