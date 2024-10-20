package http

import (
	"database/sql"
	"fmt"

	"github.com/banggibima/agile-backend/config"
	profilecommand "github.com/banggibima/agile-backend/internal/module/profile/application/command"
	profilequery "github.com/banggibima/agile-backend/internal/module/profile/application/query"
	profiledelivery "github.com/banggibima/agile-backend/internal/module/profile/delivery"
	profilepersistence "github.com/banggibima/agile-backend/internal/module/profile/infrastructure/persistence"
	usercommand "github.com/banggibima/agile-backend/internal/module/user/application/command"
	userquery "github.com/banggibima/agile-backend/internal/module/user/application/query"
	userdelivery "github.com/banggibima/agile-backend/internal/module/user/delivery"
	userpersistence "github.com/banggibima/agile-backend/internal/module/user/infrastructure/persistence"
	"github.com/banggibima/agile-backend/internal/transport/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type HTTP struct {
	Config   *config.Config
	Echo     *echo.Echo
	Logger   *logrus.Logger
	Postgres *sql.DB
}

func NewHTTP(
	config *config.Config,
	echo *echo.Echo,
	logger *logrus.Logger,
	postgres *sql.DB,
) *HTTP {
	return &HTTP{
		Config:   config,
		Echo:     echo,
		Logger:   logger,
		Postgres: postgres,
	}
}

func (h *HTTP) Set() error {
	profilePostgresRepository := profilepersistence.NewProfilePostgresRepository(h.Postgres)
	profileCommandService := profilecommand.NewProfileCommandService(profilePostgresRepository)
	profileCommandUsecase := profilecommand.NewProfileCommandUsecase(profileCommandService)
	profileQueryService := profilequery.NewProfileQueryService(profilePostgresRepository)
	profileQueryUsecase := profilequery.NewProfileQueryUsecase(profileQueryService)
	profileChecker := profiledelivery.NewProfileChecker()
	profileWrapper := profiledelivery.NewProfileWrapper()
	profileHandler := profiledelivery.NewProfileHandler(profileCommandUsecase, profileQueryUsecase, profileChecker, profileWrapper)
	profileRouter := profiledelivery.NewProfileRouter(h.Echo, profileHandler)

	userPostgresRepository := userpersistence.NewUserPostgresRepository(h.Postgres)
	userCommandService := usercommand.NewUserCommandService(userPostgresRepository)
	userCommandUsecase := usercommand.NewUserCommandUsecase(userCommandService)
	userQueryService := userquery.NewUserQueryService(userPostgresRepository)
	userQueryUsecase := userquery.NewUserQueryUsecase(userQueryService)
	userChecker := userdelivery.NewUserChecker()
	userWrapper := userdelivery.NewUserWrapper()
	userHandler := userdelivery.NewUserHandler(userCommandUsecase, userQueryUsecase, userChecker, userWrapper)
	userRouter := userdelivery.NewUserRouter(h.Echo, userHandler)

	loggerMiddleware := middleware.NewLoggerMiddleware(h.Config, h.Logger)
	corsMiddleware := middleware.NewCORSMiddleware(h.Config)

	h.Echo.Use(loggerMiddleware.WithConfig())
	h.Echo.Use(corsMiddleware.WithConfig())

	profileRouter.Resource()
	userRouter.Resource()

	return nil
}

func (h *HTTP) Start() error {
	if err := h.Set(); err != nil {
		return err
	}

	port := fmt.Sprintf(":%s", h.Config.App.Port)
	starter := fmt.Sprintf("http server started on [::]:%s", port)

	h.Logger.Infof(starter)

	if err := h.Echo.Start(port); err != nil {
		return err
	}

	return nil
}
