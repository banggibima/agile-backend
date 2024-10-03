package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/profile/domain"
)

type ProfileWrapper struct{}

func NewProfileWrapper() *ProfileWrapper {
	return &ProfileWrapper{}
}

func (w *ProfileWrapper) WrapMeta(page, size, count, total int, sort, order string) *domain.Meta {
	return &domain.Meta{
		Page:  page,
		Size:  size,
		Count: count,
		Total: total,
		Sort:  sort,
		Order: order,
	}
}

func (w *ProfileWrapper) List(meta *domain.Meta, data []*domain.Profile) *domain.List {
	return &domain.List{
		Meta: meta,
		Data: data,
	}
}

func (w *ProfileWrapper) Detail(data *domain.Profile) *domain.Detail {
	return &domain.Detail{
		Data: data,
	}
}

func (w *ProfileWrapper) Error(err error) *domain.Error {
	return &domain.Error{
		Error: err.Error(),
	}
}
