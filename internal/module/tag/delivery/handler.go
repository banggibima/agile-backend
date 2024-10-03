package delivery

import (
	"net/http"
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/tag"
	"github.com/banggibima/agile-backend/internal/module/tag/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	TagCommandUsecase tag.TagCommandUsecase
	TagQueryUsecase   tag.TagQueryUsecase
	TagChecker        tag.TagChecker
	TagWrapper        tag.TagWrapper
}

func NewTagHandler(
	tagCommandUsecase tag.TagCommandUsecase,
	tagQueryUsecase tag.TagQueryUsecase,
	tagChecker tag.TagChecker,
	tagWrapper tag.TagWrapper,
) tag.TagHandler {
	return &TagHandler{
		TagCommandUsecase: tagCommandUsecase,
		TagQueryUsecase:   tagQueryUsecase,
		TagChecker:        tagChecker,
		TagWrapper:        tagWrapper,
	}
}

func (h *TagHandler) Find(c echo.Context) error {
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

	total, err := h.TagQueryUsecase.Count()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	data, err := h.TagQueryUsecase.Find(page, size, sort, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	meta.Total = total
	meta.Count = len(data)

	meta = h.TagWrapper.WrapMeta(meta.Page, meta.Size, meta.Count, meta.Total, meta.Sort, meta.Order)

	return c.JSON(http.StatusOK, h.TagWrapper.List(meta, data))
}

func (h *TagHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagChecker.FindByID(uuid); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	data, err := h.TagQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TagWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.TagWrapper.Detail(data))
}

func (h *TagHandler) Save(c echo.Context) error {
	data := new(domain.Tag)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagChecker.Save(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagCommandUsecase.Save(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	return c.JSON(http.StatusCreated, h.TagWrapper.Detail(data))
}

func (h *TagHandler) Edit(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Tag)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	exist, err := h.TagQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TagWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagChecker.Edit(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagCommandUsecase.Edit(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.TagWrapper.Detail(data))
}

func (h *TagHandler) EditPartial(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Tag)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	exist, err := h.TagQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TagWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagChecker.EditPartial(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagCommandUsecase.EditPartial(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.TagWrapper.Detail(data))
}

func (h *TagHandler) Remove(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Tag)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	exist, err := h.TagQueryUsecase.FindByID(uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, h.TagWrapper.Error(err))
		}

		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	data.ID = exist.ID

	if err := h.TagChecker.Remove(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.TagWrapper.Error(err))
	}

	if err := h.TagCommandUsecase.Remove(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.TagWrapper.Error(err))
	}

	return c.JSON(http.StatusNoContent, nil)
}
