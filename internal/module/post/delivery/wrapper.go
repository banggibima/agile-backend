package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/post/domain"
)

type PostWrapper struct{}

func NewPostWrapper() *PostWrapper {
	return &PostWrapper{}
}

func (w *PostWrapper) WrapMeta(page, size, count, total int, sort, order string) *domain.Meta {
	return &domain.Meta{
		Page:  page,
		Size:  size,
		Count: count,
		Total: total,
		Sort:  sort,
		Order: order,
	}
}

func (w *PostWrapper) List(meta *domain.Meta, data []*domain.Post) *domain.List {
	return &domain.List{
		Meta: meta,
		Data: data,
	}
}

func (w *PostWrapper) Detail(data *domain.Post) *domain.Detail {
	return &domain.Detail{
		Data: data,
	}
}

func (w *PostWrapper) Error(err error) *domain.Error {
	return &domain.Error{
		Error: err.Error(),
	}
}
