package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/user"
	"github.com/banggibima/agile-backend/internal/module/user/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserCommandUsecase user.UserCommandUsecase
	UserQueryUsecase   user.UserQueryUsecase
	UserChecker        user.UserChecker
	UserWrapper        user.UserWrapper
}

func NewUserHandler(
	userCommandUsecase user.UserCommandUsecase,
	userQueryUsecase user.UserQueryUsecase,
	userChecker user.UserChecker,
	userWrapper user.UserWrapper,
) user.UserHandler {
	return &UserHandler{
		UserCommandUsecase: userCommandUsecase,
		UserQueryUsecase:   userQueryUsecase,
		UserChecker:        userChecker,
		UserWrapper:        userWrapper,
	}
}

func (h *UserHandler) Find(c echo.Context) error {
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

	total, err := h.UserQueryUsecase.Count()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	data, err := h.UserQueryUsecase.Find(page, size, sort, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	meta.Total = total
	meta.Count = len(data)

	meta = h.UserWrapper.WrapMeta(meta.Page, meta.Size, meta.Count, meta.Total, meta.Sort, meta.Order)

	return c.JSON(http.StatusOK, h.UserWrapper.List(meta, data))
}

func (h *UserHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserChecker.FindByID(uuid); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	data, err := h.UserQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	if data == nil {
		return c.JSON(http.StatusNotFound, h.UserWrapper.Error(errors.New("data not found")))
	}

	return c.JSON(http.StatusOK, h.UserWrapper.Detail(data))
}

func (h *UserHandler) Save(c echo.Context) error {
	data := new(domain.User)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserChecker.Save(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserCommandUsecase.Save(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	return c.JSON(http.StatusCreated, h.UserWrapper.Detail(data))
}

func (h *UserHandler) Edit(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.User)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	exist, err := h.UserQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	if exist == nil {
		return c.JSON(http.StatusNotFound, h.UserWrapper.Error(errors.New("data not found")))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserChecker.Edit(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserCommandUsecase.Edit(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.UserWrapper.Detail(data))
}

func (h *UserHandler) EditPartial(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.User)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	exist, err := h.UserQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	if exist == nil {
		return c.JSON(http.StatusNotFound, h.UserWrapper.Error(errors.New("data not found")))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserChecker.EditPartial(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserCommandUsecase.EditPartial(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.UserWrapper.Detail(data))
}

func (h *UserHandler) Remove(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.User)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	exist, err := h.UserQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	if exist == nil {
		return c.JSON(http.StatusNotFound, h.UserWrapper.Error(errors.New("data not found")))
	}

	data.ID = exist.ID

	if err := h.UserChecker.Remove(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.UserWrapper.Error(err))
	}

	if err := h.UserCommandUsecase.Remove(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.UserWrapper.Error(err))
	}

	return c.JSON(http.StatusNoContent, nil)
}
