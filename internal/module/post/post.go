package post

import (
	"github.com/banggibima/agile-backend/internal/module/post/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostPostgresRepository interface {
	Count() (int, error)
	Find(offset, limit int, sort, order string) ([]*domain.Post, error)
	FindByID(id uuid.UUID) (*domain.Post, error)
	Save(payload *domain.Post) error
	Edit(payload *domain.Post) error
	EditPartial(payload *domain.Post) error
	Remove(payload *domain.Post) error
}

type PostCommandService interface {
	Save(payload *domain.Post) error
	Edit(payload *domain.Post) error
	EditPartial(payload *domain.Post) error
	Remove(payload *domain.Post) error
}

type PostCommandUsecase interface {
	Save(payload *domain.Post) error
	Edit(payload *domain.Post) error
	EditPartial(payload *domain.Post) error
	Remove(payload *domain.Post) error
}

type PostQueryService interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Post, error)
	FindByID(id uuid.UUID) (*domain.Post, error)
}

type PostQueryUsecase interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Post, error)
	FindByID(id uuid.UUID) (*domain.Post, error)
}

type PostChecker interface {
	Find(page int, size int, sort string, order string) error
	FindByID(id uuid.UUID) error
	Save(payload *domain.Post) error
	Edit(payload *domain.Post) error
	EditPartial(payload *domain.Post) error
	Remove(payload *domain.Post) error
}

type PostWrapper interface {
	WrapMeta(page, size, count, total int, sort, order string) *domain.Meta
	List(meta *domain.Meta, data []*domain.Post) *domain.List
	Detail(data *domain.Post) *domain.Detail
	Error(err error) *domain.Error
}

type PostHandler interface {
	Find(c echo.Context) error
	FindByID(c echo.Context) error
	Save(c echo.Context) error
	Edit(c echo.Context) error
	EditPartial(c echo.Context) error
	Remove(c echo.Context) error
}

type PostRouter interface {
	Resource() error
}
