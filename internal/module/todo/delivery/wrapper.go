package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/todo/domain"
)

type TodoWrapper struct{}

func NewTodoWrapper() *TodoWrapper {
	return &TodoWrapper{}
}

func (w *TodoWrapper) WrapMeta(page, size, count, total int, sort, order string) *domain.Meta {
	return &domain.Meta{
		Page:  page,
		Size:  size,
		Count: count,
		Total: total,
		Sort:  sort,
		Order: order,
	}
}

func (w *TodoWrapper) List(meta *domain.Meta, data []*domain.Todo) *domain.List {
	return &domain.List{
		Meta: meta,
		Data: data,
	}
}

func (w *TodoWrapper) Detail(data *domain.Todo) *domain.Detail {
	return &domain.Detail{
		Data: data,
	}
}

func (w *TodoWrapper) Error(err error) *domain.Error {
	return &domain.Error{
		Error: err.Error(),
	}
}
