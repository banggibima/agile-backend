package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/post"
	"github.com/labstack/echo/v4"
)

type PostRouter struct {
	Echo        *echo.Echo
	PostHandler post.PostHandler
}

func NewPostRouter(
	echo *echo.Echo,
	postHandler post.PostHandler,
) *PostRouter {
	return &PostRouter{
		Echo:        echo,
		PostHandler: postHandler,
	}
}

func (r *PostRouter) Resource() {
	posts := r.Echo.Group("/api/posts")

	posts.GET("", r.PostHandler.Find)
	posts.GET("/:id", r.PostHandler.FindByID)
	posts.POST("", r.PostHandler.Save)
	posts.PUT("/:id", r.PostHandler.Edit)
	posts.PATCH("/:id", r.PostHandler.EditPartial)
	posts.DELETE("/:id", r.PostHandler.Remove)
}
