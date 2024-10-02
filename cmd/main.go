package main

import (
	"github.com/banggibima/backend-agile/config"
	"github.com/banggibima/backend-agile/internal/transport/http"
	"github.com/banggibima/backend-agile/pkg/echo"
	"github.com/banggibima/backend-agile/pkg/logrus"
	"github.com/banggibima/backend-agile/pkg/postgres"
	"github.com/banggibima/backend-agile/pkg/viper"
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
