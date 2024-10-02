package delivery

import (
	"github.com/banggibima/backend-agile/internal/module/todo"
	"github.com/labstack/echo/v4"
)

type TodoRouter struct {
	Echo        *echo.Echo
	TodoHandler todo.TodoHandler
}

func NewTodoRouter(
	echo *echo.Echo,
	todoHandler todo.TodoHandler,
) *TodoRouter {
	return &TodoRouter{
		Echo:        echo,
		TodoHandler: todoHandler,
	}
}

func (r *TodoRouter) Resource() {
	todos := r.Echo.Group("/api/todos")

	todos.GET("", r.TodoHandler.Find)
	todos.GET("/:id", r.TodoHandler.FindByID)
	todos.POST("", r.TodoHandler.Save)
	todos.PUT("/:id", r.TodoHandler.Edit)
	todos.PATCH("/:id", r.TodoHandler.EditPartial)
	todos.DELETE("/:id", r.TodoHandler.Remove)
}
