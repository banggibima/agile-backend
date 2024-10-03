package query

import (
	"github.com/banggibima/agile-backend/internal/module/profile"
	"github.com/banggibima/agile-backend/internal/module/profile/domain"
	"github.com/google/uuid"
)

type ProfileQueryUsecase struct {
	Service profile.ProfileQueryService
}

func NewProfileQueryUsecase(
	service profile.ProfileQueryService,
) profile.ProfileQueryUsecase {
	return &ProfileQueryUsecase{
		Service: service,
	}
}

func (u *ProfileQueryUsecase) Count() (int, error) {
	total, err := u.Service.Count()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (u *ProfileQueryUsecase) Find(offset, size int, sort, order string) ([]*domain.Profile, error) {
	data, err := u.Service.Find(offset, size, sort, order)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *ProfileQueryUsecase) FindByID(id uuid.UUID) (*domain.Profile, error) {
	data, err := u.Service.FindByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
