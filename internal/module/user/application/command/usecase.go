package command

import (
	"github.com/banggibima/agile-backend/internal/module/user"
	"github.com/banggibima/agile-backend/internal/module/user/domain"
)

type UserCommandUsecase struct {
	Service user.UserCommandService
}

func NewUserCommandUsecase(
	service user.UserCommandService,
) user.UserCommandUsecase {
	return &UserCommandUsecase{
		Service: service,
	}
}

func (u *UserCommandUsecase) Save(payload *domain.User) error {
	err := u.Service.Save(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserCommandUsecase) Edit(payload *domain.User) error {
	err := u.Service.Edit(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserCommandUsecase) EditPartial(payload *domain.User) error {
	err := u.Service.EditPartial(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserCommandUsecase) Remove(payload *domain.User) error {
	err := u.Service.Remove(payload)
	if err != nil {
		return err
	}

	return nil
}
