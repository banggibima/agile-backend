package delivery

import (
	"github.com/banggibima/backend-agile/internal/module/todo/domain"
)

type TodoWrapper struct{}

func NewTodoWrapper() *TodoWrapper {
	return &TodoWrapper{}
}

func (w *TodoWrapper) WrapMeta(page, size, total int, sort, order string) *domain.Meta {
	return &domain.Meta{
		Page:  page,
		Size:  size,
		Total: total,
		Sort:  sort,
		Order: order,
	}
}

func (w *TodoWrapper) WrapList(meta *domain.Meta, data []*domain.Todo) *domain.List {
	return &domain.List{
		Meta: meta,
		Data: data,
	}
}

func (w *TodoWrapper) WrapDetail(data *domain.Todo) *domain.Detail {
	return &domain.Detail{
		Data: data,
	}
}

func (w *TodoWrapper) WrapError(err error) *domain.Error {
	return &domain.Error{
		Meta:  nil,
		Error: err.Error(),
	}
}
