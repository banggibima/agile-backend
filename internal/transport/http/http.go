package http

import (
	"database/sql"
	"fmt"

	"github.com/banggibima/backend-agile/config"
	todocommand "github.com/banggibima/backend-agile/internal/module/todo/application/command"
	todoquery "github.com/banggibima/backend-agile/internal/module/todo/application/query"
	tododelivery "github.com/banggibima/backend-agile/internal/module/todo/delivery"
	todopersistence "github.com/banggibima/backend-agile/internal/module/todo/infrastructure/persistence"
	usercommand "github.com/banggibima/backend-agile/internal/module/user/application/command"
	userquery "github.com/banggibima/backend-agile/internal/module/user/application/query"
	userdelivery "github.com/banggibima/backend-agile/internal/module/user/delivery"
	userpersistence "github.com/banggibima/backend-agile/internal/module/user/infrastructure/persistence"
	"github.com/banggibima/backend-agile/internal/transport/middleware"
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
	todoPostgresRepository := todopersistence.NewTodoPostgresRepository(h.Postgres)
	todoCommandService := todocommand.NewTodoCommandService(todoPostgresRepository)
	todoCommandUsecase := todocommand.NewTodoCommandUsecase(todoCommandService)
	todoQueryService := todoquery.NewTodoQueryService(todoPostgresRepository)
	todoQueryUsecase := todoquery.NewTodoQueryUsecase(todoQueryService)
	todoChecker := tododelivery.NewTodoChecker()
	todoWrapper := tododelivery.NewTodoWrapper()
	todoHandler := tododelivery.NewTodoHandler(todoCommandUsecase, todoQueryUsecase, todoChecker, todoWrapper)
	todoRouter := tododelivery.NewTodoRouter(h.Echo, todoHandler)

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

	todoRouter.Resource()
	userRouter.Resource()

	return nil
}

func (h *HTTP) Start() error {
	if err := h.Set(); err != nil {
		return err
	}

	port := fmt.Sprintf(":%d", h.Config.App.Port)
	starter := fmt.Sprintf("http server started on [::]:%s", port)

	h.Logger.Infof(starter)

	if err := h.Echo.Start(port); err != nil {
		return err
	}

	return nil
}
