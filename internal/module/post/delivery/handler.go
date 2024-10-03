package delivery

import (
	"net/http"
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/post"
	"github.com/banggibima/agile-backend/internal/module/post/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	PostCommandUsecase post.PostCommandUsecase
	PostQueryUsecase   post.PostQueryUsecase
	PostChecker        post.PostChecker
	PostWrapper        post.PostWrapper
}

func NewPostHandler(
	postCommandUsecase post.PostCommandUsecase,
	postQueryUsecase post.PostQueryUsecase,
	postChecker post.PostChecker,
	postWrapper post.PostWrapper,
) post.PostHandler {
	return &PostHandler{
		PostCommandUsecase: postCommandUsecase,
		PostQueryUsecase:   postQueryUsecase,
		PostChecker:        postChecker,
		PostWrapper:        postWrapper,
	}
}

func (h *PostHandler) Find(c echo.Context) error {
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

	total, err := h.PostQueryUsecase.Count()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	data, err := h.PostQueryUsecase.Find(page, size, sort, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	meta.Total = total
	meta.Count = len(data)

	meta = h.PostWrapper.WrapMeta(meta.Page, meta.Size, meta.Count, meta.Total, meta.Sort, meta.Order)

	return c.JSON(http.StatusOK, h.PostWrapper.List(meta, data))
}

func (h *PostHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostChecker.FindByID(uuid); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	data, err := h.PostQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.PostWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.PostWrapper.Detail(data))
}

func (h *PostHandler) Save(c echo.Context) error {
	data := new(domain.Post)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostChecker.Save(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostCommandUsecase.Save(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	return c.JSON(http.StatusCreated, h.PostWrapper.Detail(data))
}

func (h *PostHandler) Edit(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Post)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	exist, err := h.PostQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.PostWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostChecker.Edit(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostCommandUsecase.Edit(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.PostWrapper.Detail(data))
}

func (h *PostHandler) EditPartial(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Post)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	exist, err := h.PostQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.PostWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostChecker.EditPartial(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostCommandUsecase.EditPartial(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.PostWrapper.Detail(data))
}

func (h *PostHandler) Remove(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Post)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	exist, err := h.PostQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.PostWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := h.PostChecker.Remove(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.PostWrapper.Error(err))
	}

	if err := h.PostCommandUsecase.Remove(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.PostWrapper.Error(err))
	}

	return c.JSON(http.StatusNoContent, nil)
}
