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

func (h *HTTP) TodoModule() {
	postgresRepository := todopersistence.NewTodoPostgresRepository(h.Postgres)

	commandService := todocommand.NewTodoCommandService(postgresRepository)
	commandUsecase := todocommand.NewTodoCommandUsecase(commandService)

	queryService := todoquery.NewTodoQueryService(postgresRepository)
	queryUsecase := todoquery.NewTodoQueryUsecase(queryService)

	checker := tododelivery.NewTodoChecker()
	wrapper := tododelivery.NewTodoWrapper()
	handler := tododelivery.NewTodoHandler(commandUsecase, queryUsecase, checker, wrapper)
	router := tododelivery.NewTodoRouter(h.Echo, handler)

	router.Resource()
}

func (h *HTTP) UserModule() {
	postgresRepository := userpersistence.NewUserPostgresRepository(h.Postgres)

	commandService := usercommand.NewUserCommandService(postgresRepository)
	commandUsecase := usercommand.NewUserCommandUsecase(commandService)

	queryService := userquery.NewUserQueryService(postgresRepository)
	queryUsecase := userquery.NewUserQueryUsecase(queryService)

	checker := userdelivery.NewUserChecker()
	wrapper := userdelivery.NewUserWrapper()
	handler := userdelivery.NewUserHandler(commandUsecase, queryUsecase, checker, wrapper)
	router := userdelivery.NewUserRouter(h.Echo, handler)

	router.Resource()
}

func (h *HTTP) Set() error {
	h.TodoModule()
	h.UserModule()

	loggerMiddleware := middleware.NewLoggerMiddleware(h.Config, h.Logger)
	corsMiddleware := middleware.NewCORSMiddleware(h.Config)

	h.Echo.Use(loggerMiddleware.WithConfig())
	h.Echo.Use(corsMiddleware.WithConfig())

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
