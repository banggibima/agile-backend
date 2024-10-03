package command

import (
	"github.com/banggibima/agile-backend/internal/module/tag"
	"github.com/banggibima/agile-backend/internal/module/tag/domain"
)

type TagCommandUsecase struct {
	Service tag.TagCommandService
}

func NewTagCommandUsecase(
	service tag.TagCommandService,
) tag.TagCommandUsecase {
	return &TagCommandUsecase{
		Service: service,
	}
}

func (u *TagCommandUsecase) Save(payload *domain.Tag) error {
	err := u.Service.Save(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *TagCommandUsecase) Edit(payload *domain.Tag) error {
	err := u.Service.Edit(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *TagCommandUsecase) EditPartial(payload *domain.Tag) error {
	err := u.Service.EditPartial(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *TagCommandUsecase) Remove(payload *domain.Tag) error {
	err := u.Service.Remove(payload)
	if err != nil {
		return err
	}

	return nil
}
