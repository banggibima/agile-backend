package command

import (
	"github.com/banggibima/agile-backend/internal/module/profile"
	"github.com/banggibima/agile-backend/internal/module/profile/domain"
)

type ProfileCommandUsecase struct {
	Service profile.ProfileCommandService
}

func NewProfileCommandUsecase(
	service profile.ProfileCommandService,
) profile.ProfileCommandUsecase {
	return &ProfileCommandUsecase{
		Service: service,
	}
}

func (u *ProfileCommandUsecase) Save(payload *domain.Profile) error {
	err := u.Service.Save(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProfileCommandUsecase) Edit(payload *domain.Profile) error {
	err := u.Service.Edit(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProfileCommandUsecase) EditPartial(payload *domain.Profile) error {
	err := u.Service.EditPartial(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *ProfileCommandUsecase) Remove(payload *domain.Profile) error {
	err := u.Service.Remove(payload)
	if err != nil {
		return err
	}

	return nil
}
