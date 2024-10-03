package query

import (
	"github.com/banggibima/agile-backend/internal/module/user"
	"github.com/banggibima/agile-backend/internal/module/user/domain"
	"github.com/google/uuid"
)

type UserQueryUsecase struct {
	Service user.UserQueryService
}

func NewUserQueryUsecase(
	service user.UserQueryService,
) user.UserQueryUsecase {
	return &UserQueryUsecase{
		Service: service,
	}
}

func (u *UserQueryUsecase) Count() (int, error) {
	total, err := u.Service.Count()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (u *UserQueryUsecase) Find(offset, size int, sort, order string) ([]*domain.User, error) {
	data, err := u.Service.Find(offset, size, sort, order)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserQueryUsecase) FindByID(id uuid.UUID) (*domain.User, error) {
	data, err := u.Service.FindByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
