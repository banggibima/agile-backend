package http

import (
	"database/sql"
	"fmt"

	"github.com/banggibima/agile-backend/config"
	postcommand "github.com/banggibima/agile-backend/internal/module/post/application/command"
	postquery "github.com/banggibima/agile-backend/internal/module/post/application/query"
	postdelivery "github.com/banggibima/agile-backend/internal/module/post/delivery"
	postpersistence "github.com/banggibima/agile-backend/internal/module/post/infrastructure/persistence"
	profilecommand "github.com/banggibima/agile-backend/internal/module/profile/application/command"
	profilequery "github.com/banggibima/agile-backend/internal/module/profile/application/query"
	profiledelivery "github.com/banggibima/agile-backend/internal/module/profile/delivery"
	profilepersistence "github.com/banggibima/agile-backend/internal/module/profile/infrastructure/persistence"
	tagcommand "github.com/banggibima/agile-backend/internal/module/tag/application/command"
	tagquery "github.com/banggibima/agile-backend/internal/module/tag/application/query"
	tagdelivery "github.com/banggibima/agile-backend/internal/module/tag/delivery"
	tagpersistence "github.com/banggibima/agile-backend/internal/module/tag/infrastructure/persistence"
	todocommand "github.com/banggibima/agile-backend/internal/module/todo/application/command"
	todoquery "github.com/banggibima/agile-backend/internal/module/todo/application/query"
	tododelivery "github.com/banggibima/agile-backend/internal/module/todo/delivery"
	todopersistence "github.com/banggibima/agile-backend/internal/module/todo/infrastructure/persistence"
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
	postPostgresRepository := postpersistence.NewPostPostgresRepository(h.Postgres)
	postCommandService := postcommand.NewPostCommandService(postPostgresRepository)
	postCommandUsecase := postcommand.NewPostCommandUsecase(postCommandService)
	postQueryService := postquery.NewPostQueryService(postPostgresRepository)
	postQueryUsecase := postquery.NewPostQueryUsecase(postQueryService)
	postChecker := postdelivery.NewPostChecker()
	postWrapper := postdelivery.NewPostWrapper()
	postHandler := postdelivery.NewPostHandler(postCommandUsecase, postQueryUsecase, postChecker, postWrapper)
	postRouter := postdelivery.NewPostRouter(h.Echo, postHandler)

	profilePostgresRepository := profilepersistence.NewProfilePostgresRepository(h.Postgres)
	profileCommandService := profilecommand.NewProfileCommandService(profilePostgresRepository)
	profileCommandUsecase := profilecommand.NewProfileCommandUsecase(profileCommandService)
	profileQueryService := profilequery.NewProfileQueryService(profilePostgresRepository)
	profileQueryUsecase := profilequery.NewProfileQueryUsecase(profileQueryService)
	profileChecker := profiledelivery.NewProfileChecker()
	profileWrapper := profiledelivery.NewProfileWrapper()
	profileHandler := profiledelivery.NewProfileHandler(profileCommandUsecase, profileQueryUsecase, profileChecker, profileWrapper)
	profileRouter := profiledelivery.NewProfileRouter(h.Echo, profileHandler)

	tagPostgresRepository := tagpersistence.NewTagPostgresRepository(h.Postgres)
	tagCommandService := tagcommand.NewTagCommandService(tagPostgresRepository)
	tagCommandUsecase := tagcommand.NewTagCommandUsecase(tagCommandService)
	tagQueryService := tagquery.NewTagQueryService(tagPostgresRepository)
	tagQueryUsecase := tagquery.NewTagQueryUsecase(tagQueryService)
	tagChecker := tagdelivery.NewTagChecker()
	tagWrapper := tagdelivery.NewTagWrapper()
	tagHandler := tagdelivery.NewTagHandler(tagCommandUsecase, tagQueryUsecase, tagChecker, tagWrapper)
	tagRouter := tagdelivery.NewTagRouter(h.Echo, tagHandler)

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

	postRouter.Resource()
	profileRouter.Resource()
	tagRouter.Resource()
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
