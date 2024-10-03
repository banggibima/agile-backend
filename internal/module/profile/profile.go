package profile

import (
	"github.com/banggibima/agile-backend/internal/module/profile/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProfilePostgresRepository interface {
	Count() (int, error)
	Find(offset, limit int, sort, order string) ([]*domain.Profile, error)
	FindByID(id uuid.UUID) (*domain.Profile, error)
	Save(payload *domain.Profile) error
	Edit(payload *domain.Profile) error
	EditPartial(payload *domain.Profile) error
	Remove(payload *domain.Profile) error
}

type ProfileCommandService interface {
	Save(payload *domain.Profile) error
	Edit(payload *domain.Profile) error
	EditPartial(payload *domain.Profile) error
	Remove(payload *domain.Profile) error
}

type ProfileCommandUsecase interface {
	Save(payload *domain.Profile) error
	Edit(payload *domain.Profile) error
	EditPartial(payload *domain.Profile) error
	Remove(payload *domain.Profile) error
}

type ProfileQueryService interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Profile, error)
	FindByID(id uuid.UUID) (*domain.Profile, error)
}

type ProfileQueryUsecase interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Profile, error)
	FindByID(id uuid.UUID) (*domain.Profile, error)
}

type ProfileChecker interface {
	Find(page int, size int, sort string, order string) error
	FindByID(id uuid.UUID) error
	Save(payload *domain.Profile) error
	Edit(payload *domain.Profile) error
	EditPartial(payload *domain.Profile) error
	Remove(payload *domain.Profile) error
}

type ProfileWrapper interface {
	WrapMeta(page, size, count, total int, sort, order string) *domain.Meta
	List(meta *domain.Meta, data []*domain.Profile) *domain.List
	Detail(data *domain.Profile) *domain.Detail
	Error(err error) *domain.Error
}

type ProfileHandler interface {
	Find(c echo.Context) error
	FindByID(c echo.Context) error
	Save(c echo.Context) error
	Edit(c echo.Context) error
	EditPartial(c echo.Context) error
	Remove(c echo.Context) error
}

type ProfileRouter interface {
	Resource() error
}
