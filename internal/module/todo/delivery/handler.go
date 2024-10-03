package delivery

import (
	"net/http"
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/todo"
	"github.com/banggibima/agile-backend/internal/module/todo/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoCommandUsecase todo.TodoCommandUsecase
	TodoQueryUsecase   todo.TodoQueryUsecase
	TodoChecker        todo.TodoChecker
	TodoWrapper        todo.TodoWrapper
}

func NewTodoHandler(
	todoCommandUsecase todo.TodoCommandUsecase,
	todoQueryUsecase todo.TodoQueryUsecase,
	todoChecker todo.TodoChecker,
	todoWrapper todo.TodoWrapper,
) todo.TodoHandler {
	return &TodoHandler{
		TodoCommandUsecase: todoCommandUsecase,
		TodoQueryUsecase:   todoQueryUsecase,
		TodoChecker:        todoChecker,
		TodoWrapper:        todoWrapper,
	}
}

func (h *TodoHandler) Find(c echo.Context) error {
	meta := new(domain.Meta)

	meta.Page, _ = strconv.Atoi(c.QueryParam("page"))
	meta.Size, _ = strconv.Atoi(c.QueryParam("size"))

	if meta.Page <= 0 && meta.Size == 0 {
		meta.Page = 0
		meta.Size = 0
	} else {
		if meta.Page <= 0 {
			meta.Page = 1
		}
		if meta.Size <= 0 {
			meta.Size = 10
		}
	}

	meta.Sort = c.QueryParam("sort")
	meta.Order = c.QueryParam("order")

	if meta.Sort == "" {
		meta.Sort = "created_at"
	}
	if meta.Order == "" {
		meta.Order = "desc"
	}

	page := 0
	size := 0

	if meta.Page == 0 && meta.Size == 0 {
		page = 0
		size = 0
	} else {
		page = (meta.Page - 1) * meta.Size
		size = meta.Size
	}

	sort := meta.Sort
	order := meta.Order

	total, err := h.TodoQueryUsecase.Count()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	data, err := h.TodoQueryUsecase.Find(page, size, sort, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	meta.Total = total
	meta.Count = len(data)

	meta = h.TodoWrapper.WrapMeta(meta.Page, meta.Size, meta.Count, meta.Total, meta.Sort, meta.Order)

	return c.JSON(http.StatusOK, h.TodoWrapper.List(meta, data))
}

func (h *TodoHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoChecker.FindByID(uuid); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	data, err := h.TodoQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TodoWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.TodoWrapper.Detail(data))
}

func (h *TodoHandler) Save(c echo.Context) error {
	data := new(domain.Todo)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoChecker.Save(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoCommandUsecase.Save(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	return c.JSON(http.StatusCreated, h.TodoWrapper.Detail(data))
}

func (h *TodoHandler) Edit(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Todo)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	exist, err := h.TodoQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TodoWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoChecker.Edit(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoCommandUsecase.Edit(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.TodoWrapper.Detail(data))
}

func (h *TodoHandler) EditPartial(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Todo)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	exist, err := h.TodoQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TodoWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoChecker.EditPartial(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoCommandUsecase.EditPartial(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.TodoWrapper.Detail(data))
}

func (h *TodoHandler) Remove(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Todo)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	exist, err := h.TodoQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TodoWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := h.TodoChecker.Remove(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TodoWrapper.Error(err))
	}

	if err := h.TodoCommandUsecase.Remove(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TodoWrapper.Error(err))
	}

	return c.JSON(http.StatusNoContent, nil)
}
