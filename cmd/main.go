package main

import (
	"github.com/banggibima/agile-backend/config"
	"github.com/banggibima/agile-backend/internal/transport/http"
	"github.com/banggibima/agile-backend/pkg/echo"
	"github.com/banggibima/agile-backend/pkg/logrus"
	"github.com/banggibima/agile-backend/pkg/postgres"
)

func main() {
	config, err := config.Load()
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

	postgres, err := postgres.Client(config, logrus)
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
