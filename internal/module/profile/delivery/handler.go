package delivery

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/banggibima/agile-backend/internal/module/profile"
	"github.com/banggibima/agile-backend/internal/module/profile/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	ProfileCommandUsecase profile.ProfileCommandUsecase
	ProfileQueryUsecase   profile.ProfileQueryUsecase
	ProfileChecker        profile.ProfileChecker
	ProfileWrapper        profile.ProfileWrapper
}

func NewProfileHandler(
	profileCommandUsecase profile.ProfileCommandUsecase,
	profileQueryUsecase profile.ProfileQueryUsecase,
	profileChecker profile.ProfileChecker,
	profileWrapper profile.ProfileWrapper,
) profile.ProfileHandler {
	return &ProfileHandler{
		ProfileCommandUsecase: profileCommandUsecase,
		ProfileQueryUsecase:   profileQueryUsecase,
		ProfileChecker:        profileChecker,
		ProfileWrapper:        profileWrapper,
	}
}

func (h *ProfileHandler) Find(c echo.Context) error {
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

	total, err := h.ProfileQueryUsecase.Count()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	data, err := h.ProfileQueryUsecase.Find(page, size, sort, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	meta.Total = total
	meta.Count = len(data)

	meta = h.ProfileWrapper.WrapMeta(meta.Page, meta.Size, meta.Count, meta.Total, meta.Sort, meta.Order)

	return c.JSON(http.StatusOK, h.ProfileWrapper.List(meta, data))
}

func (h *ProfileHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileChecker.FindByID(uuid); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	data, err := h.ProfileQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	if data == nil {
		return c.JSON(http.StatusNotFound, h.ProfileWrapper.Error(errors.New("data not found")))
	}

	return c.JSON(http.StatusOK, h.ProfileWrapper.Detail(data))
}

func (h *ProfileHandler) Save(c echo.Context) error {
	data := new(domain.Profile)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileChecker.Save(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileCommandUsecase.Save(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	return c.JSON(http.StatusCreated, h.ProfileWrapper.Detail(data))
}

func (h *ProfileHandler) Edit(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Profile)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	exist, err := h.ProfileQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	if exist == nil {
		return c.JSON(http.StatusNotFound, h.ProfileWrapper.Error(errors.New("data not found")))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileChecker.Edit(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileCommandUsecase.Edit(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.ProfileWrapper.Detail(data))
}

func (h *ProfileHandler) EditPartial(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Profile)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	exist, err := h.ProfileQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	if exist == nil {
		return c.JSON(http.StatusNotFound, h.ProfileWrapper.Error(errors.New("data not found")))
	}

	data.ID = exist.ID

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileChecker.EditPartial(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileCommandUsecase.EditPartial(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	return c.JSON(http.StatusOK, h.ProfileWrapper.Detail(data))
}

func (h *ProfileHandler) Remove(c echo.Context) error {
	id := c.Param("id")
	data := new(domain.Profile)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	exist, err := h.ProfileQueryUsecase.FindByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	if exist == nil {
		return c.JSON(http.StatusNotFound, h.ProfileWrapper.Error(errors.New("data not found")))
	}

	data.ID = exist.ID

	if err := h.ProfileChecker.Remove(data); err != nil {
		return c.JSON(http.StatusBadRequest, h.ProfileWrapper.Error(err))
	}

	if err := h.ProfileCommandUsecase.Remove(data); err != nil {
		return c.JSON(http.StatusInternalServerError, h.ProfileWrapper.Error(err))
	}

	return c.JSON(http.StatusNoContent, nil)
}
