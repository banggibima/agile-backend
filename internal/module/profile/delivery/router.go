package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/profile"
	"github.com/labstack/echo/v4"
)

type ProfileRouter struct {
	Echo           *echo.Echo
	ProfileHandler profile.ProfileHandler
}

func NewProfileRouter(
	echo *echo.Echo,
	profileHandler profile.ProfileHandler,
) *ProfileRouter {
	return &ProfileRouter{
		Echo:           echo,
		ProfileHandler: profileHandler,
	}
}

func (r *ProfileRouter) Resource() {
	profiles := r.Echo.Group("/api/profiles")

	profiles.GET("", r.ProfileHandler.Find)
	profiles.GET("/:id", r.ProfileHandler.FindByID)
	profiles.POST("", r.ProfileHandler.Save)
	profiles.PUT("/:id", r.ProfileHandler.Edit)
	profiles.PATCH("/:id", r.ProfileHandler.EditPartial)
	profiles.DELETE("/:id", r.ProfileHandler.Remove)
}
