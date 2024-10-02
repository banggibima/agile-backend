package delivery

import (
	"github.com/banggibima/backend-agile/internal/module/user/domain"
)

type UserWrapper struct{}

func NewUserWrapper() *UserWrapper {
	return &UserWrapper{}
}

func (w *UserWrapper) WrapMeta(page, size, total int, sort, order string) *domain.Meta {
	return &domain.Meta{
		Page:  page,
		Size:  size,
		Total: total,
		Sort:  sort,
		Order: order,
	}
}

func (w *UserWrapper) WrapList(meta *domain.Meta, data []*domain.User) *domain.List {
	return &domain.List{
		Meta: meta,
		Data: data,
	}
}

func (w *UserWrapper) WrapDetail(data *domain.User) *domain.Detail {
	return &domain.Detail{
		Data: data,
	}
}

func (w *UserWrapper) WrapError(err error) *domain.Error {
	return &domain.Error{
		Meta:  nil,
		Error: err.Error(),
	}
}
