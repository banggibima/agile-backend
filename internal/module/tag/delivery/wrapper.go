package delivery

import (
	"github.com/banggibima/agile-backend/internal/module/tag/domain"
)

type TagWrapper struct{}

func NewTagWrapper() *TagWrapper {
	return &TagWrapper{}
}

func (w *TagWrapper) WrapMeta(page, size, count, total int, sort, order string) *domain.Meta {
	return &domain.Meta{
		Page:  page,
		Size:  size,
		Count: count,
		Total: total,
		Sort:  sort,
		Order: order,
	}
}

func (w *TagWrapper) List(meta *domain.Meta, data []*domain.Tag) *domain.List {
	return &domain.List{
		Meta: meta,
		Data: data,
	}
}

func (w *TagWrapper) Detail(data *domain.Tag) *domain.Detail {
	return &domain.Detail{
		Data: data,
	}
}

func (w *TagWrapper) Error(err error) *domain.Error {
	return &domain.Error{
		Error: err.Error(),
	}
}
