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

	cfg, err := config.Get(v)
	if err != nil {
		panic(err)
	}

	log, err := logrus.Init(cfg)
	if err != nil {
		panic(err)
	}

	e, err := echo.Init(cfg)
	if err != nil {
		panic(err)
	}

	pq, err := postgres.Client(cfg)
	if err != nil {
		panic(err)
	}

	if err := http.NewHTTP(
		cfg,
		e,
		log,
		pq,
	).Start(); err != nil {
		panic(err)
	}
}
