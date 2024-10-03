package query

import (
	"github.com/banggibima/agile-backend/internal/module/tag"
	"github.com/banggibima/agile-backend/internal/module/tag/domain"
	"github.com/google/uuid"
)

type TagQueryUsecase struct {
	Service tag.TagQueryService
}

func NewTagQueryUsecase(
	service tag.TagQueryService,
) tag.TagQueryUsecase {
	return &TagQueryUsecase{
		Service: service,
	}
}

func (u *TagQueryUsecase) Count() (int, error) {
	total, err := u.Service.Count()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (u *TagQueryUsecase) Find(offset, size int, sort, order string) ([]*domain.Tag, error) {
	data, err := u.Service.Find(offset, size, sort, order)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *TagQueryUsecase) FindByID(id uuid.UUID) (*domain.Tag, error) {
	data, err := u.Service.FindByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
