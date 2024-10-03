package query

import (
	"github.com/banggibima/agile-backend/internal/module/post"
	"github.com/banggibima/agile-backend/internal/module/post/domain"
	"github.com/google/uuid"
)

type PostQueryUsecase struct {
	Service post.PostQueryService
}

func NewPostQueryUsecase(
	service post.PostQueryService,
) post.PostQueryUsecase {
	return &PostQueryUsecase{
		Service: service,
	}
}

func (u *PostQueryUsecase) Count() (int, error) {
	total, err := u.Service.Count()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (u *PostQueryUsecase) Find(offset, size int, sort, order string) ([]*domain.Post, error) {
	data, err := u.Service.Find(offset, size, sort, order)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *PostQueryUsecase) FindByID(id uuid.UUID) (*domain.Post, error) {
	data, err := u.Service.FindByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
