package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/tag"
	"github.com/labstack/echo/v4"
)

type TagRouter struct {
	Echo       *echo.Echo
	TagHandler tag.TagHandler
}

func NewTagRouter(
	echo *echo.Echo,
	tagHandler tag.TagHandler,
) *TagRouter {
	return &TagRouter{
		Echo:       echo,
		TagHandler: tagHandler,
	}
}

func (r *TagRouter) Resource() {
	tags := r.Echo.Group("/api/tags")

	tags.GET("", r.TagHandler.Find)
	tags.GET("/:id", r.TagHandler.FindByID)
	tags.POST("", r.TagHandler.Save)
	tags.PUT("/:id", r.TagHandler.Edit)
	tags.PATCH("/:id", r.TagHandler.EditPartial)
	tags.DELETE("/:id", r.TagHandler.Remove)
}
