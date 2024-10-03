package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/user/domain"
)

type UserWrapper struct{}

func NewUserWrapper() *UserWrapper {
	return &UserWrapper{}
}

func (w *UserWrapper) WrapMeta(page, size, count, total int, sort, order string) *domain.Meta {
	return &domain.Meta{
		Page:  page,
		Size:  size,
		Count: count,
		Total: total,
		Sort:  sort,
		Order: order,
	}
}

func (w *UserWrapper) List(meta *domain.Meta, data []*domain.User) *domain.List {
	return &domain.List{
		Meta: meta,
		Data: data,
	}
}

func (w *UserWrapper) Detail(data *domain.User) *domain.Detail {
	return &domain.Detail{
		Data: data,
	}
}

func (w *UserWrapper) Error(err error) *domain.Error {
	return &domain.Error{
		Error: err.Error(),
	}
}
