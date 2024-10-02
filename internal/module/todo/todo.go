package todo

import (
	"github.com/banggibima/backend-agile/internal/module/todo/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoPostgresRepository interface {
	Count() (int, error)
	Find(offset, limit int, sort, order string) ([]*domain.Todo, error)
	FindByID(id uuid.UUID) (*domain.Todo, error)
	Save(payload *domain.Todo) error
	Edit(payload *domain.Todo) error
	EditPartial(payload *domain.Todo) error
	Remove(payload *domain.Todo) error
}

type TodoCommandService interface {
	Save(payload *domain.Todo) error
	Edit(payload *domain.Todo) error
	EditPartial(payload *domain.Todo) error
	Remove(payload *domain.Todo) error
}

type TodoCommandUsecase interface {
	Save(payload *domain.Todo) error
	Edit(payload *domain.Todo) error
	EditPartial(payload *domain.Todo) error
	Remove(payload *domain.Todo) error
}

type TodoQueryService interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Todo, error)
	FindByID(id uuid.UUID) (*domain.Todo, error)
}

type TodoQueryUsecase interface {
	Count() (int, error)
	Find(page, size int, sort, order string) ([]*domain.Todo, error)
	FindByID(id uuid.UUID) (*domain.Todo, error)
}

type TodoChecker interface {
	Find(page int, size int, sort string, order string) error
	FindByID(id uuid.UUID) error
	Save(payload *domain.Todo) error
	Edit(payload *domain.Todo) error
	EditPartial(payload *domain.Todo) error
	Remove(payload *domain.Todo) error
}

type TodoWrapper interface {
	WrapMeta(page, size, total int, sort, order string) *domain.Meta
	WrapList(meta *domain.Meta, data []*domain.Todo) *domain.List
	WrapDetail(data *domain.Todo) *domain.Detail
	WrapError(err error) *domain.Error
}

type TodoHandler interface {
	Find(c echo.Context) error
	FindByID(c echo.Context) error
	Save(c echo.Context) error
	Edit(c echo.Context) error
	EditPartial(c echo.Context) error
	Remove(c echo.Context) error
}

type TodoRouter interface {
	Resource() error
}
