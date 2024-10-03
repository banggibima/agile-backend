package tag

import (
	"github.com/banggibima/agile-backend/internal/module/tag/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TagPostgresRepository interface {
	Count() (int, error)
	Find(offset, limit int, sort, order string) ([]*domain.Tag, error)
	FindByID(id uuid.UUID) (*domain.Tag, error)
	Save(payload *domain.Tag) error
	Edit(payload *domain.Tag) error
	EditPartial(payload *domain.Tag) error
	Remove(payload *domain.Tag) error
}

type TagCommandService interface {
	Save(payload *domain.Tag) error
	Edit(payload *domain.Tag) error
	EditPartial(payload *domain.Tag) error
	Remove(payload *domain.Tag) error
}

type TagCommandUsecase interface {
	Save(payload *domain.Tag) error
	Edit(payload *domain.Tag) error
	EditPartial(payload *domain.Tag) error
	Remove(payload *domain.Tag) error
}

type TagQueryService interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Tag, error)
	FindByID(id uuid.UUID) (*domain.Tag, error)
}

type TagQueryUsecase interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Tag, error)
	FindByID(id uuid.UUID) (*domain.Tag, error)
}

type TagChecker interface {
	Find(page int, size int, sort string, order string) error
	FindByID(id uuid.UUID) error
	Save(payload *domain.Tag) error
	Edit(payload *domain.Tag) error
	EditPartial(payload *domain.Tag) error
	Remove(payload *domain.Tag) error
}

type TagWrapper interface {
	WrapMeta(page, size, count, total int, sort, order string) *domain.Meta
	List(meta *domain.Meta, data []*domain.Tag) *domain.List
	Detail(data *domain.Tag) *domain.Detail
	Error(err error) *domain.Error
}

type TagHandler interface {
	Find(c echo.Context) error
	FindByID(c echo.Context) error
	Save(c echo.Context) error
	Edit(c echo.Context) error
	EditPartial(c echo.Context) error
	Remove(c echo.Context) error
}

type TagRouter interface {
	Resource() error
}
