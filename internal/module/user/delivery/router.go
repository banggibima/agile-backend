package delivery

import (
	"github.com/banggibima/backend-agile/internal/module/user"
	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	Echo        *echo.Echo
	UserHandler user.UserHandler
}

func NewUserRouter(
	echo *echo.Echo,
	userHandler user.UserHandler,
) *UserRouter {
	return &UserRouter{
		Echo:        echo,
		UserHandler: userHandler,
	}
}

func (r *UserRouter) Resource() {
	users := r.Echo.Group("/api/users")

	users.GET("", r.UserHandler.Find)
	users.GET("/:id", r.UserHandler.FindByID)
	users.POST("", r.UserHandler.Save)
	users.PUT("/:id", r.UserHandler.Edit)
	users.PATCH("/:id", r.UserHandler.EditPartial)
	users.DELETE("/:id", r.UserHandler.Remove)
}
