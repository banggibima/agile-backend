package user

import (
	"github.com/banggibima/agile-backend/internal/module/user/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserPostgresRepository interface {
	Count() (int, error)
	Find(offset, limit int, sort, order string) ([]*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
	Save(payload *domain.User) error
	Edit(payload *domain.User) error
	EditPartial(payload *domain.User) error
	Remove(payload *domain.User) error
}

type UserCommandService interface {
	Save(payload *domain.User) error
	Edit(payload *domain.User) error
	EditPartial(payload *domain.User) error
	Remove(payload *domain.User) error
}

type UserCommandUsecase interface {
	Save(payload *domain.User) error
	Edit(payload *domain.User) error
	EditPartial(payload *domain.User) error
	Remove(payload *domain.User) error
}

type UserQueryService interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
}

type UserQueryUsecase interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
}

type UserChecker interface {
	Find(page int, size int, sort string, order string) error
	FindByID(id uuid.UUID) error
	Save(payload *domain.User) error
	Edit(payload *domain.User) error
	EditPartial(payload *domain.User) error
	Remove(payload *domain.User) error
}

type UserWrapper interface {
	WrapMeta(page, size, count, total int, sort, order string) *domain.Meta
	List(meta *domain.Meta, data []*domain.User) *domain.List
	Detail(data *domain.User) *domain.Detail
	Error(err error) *domain.Error
}

type UserHandler interface {
	Find(c echo.Context) error
	FindByID(c echo.Context) error
	Save(c echo.Context) error
	Edit(c echo.Context) error
	EditPartial(c echo.Context) error
	Remove(c echo.Context) error
}

type UserRouter interface {
	Resource() error
}
