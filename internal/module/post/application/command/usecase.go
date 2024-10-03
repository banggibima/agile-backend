package command

import (
	"github.com/banggibima/agile-backend/internal/module/post"
	"github.com/banggibima/agile-backend/internal/module/post/domain"
)

type PostCommandUsecase struct {
	Service post.PostCommandService
}

func NewPostCommandUsecase(
	service post.PostCommandService,
) post.PostCommandUsecase {
	return &PostCommandUsecase{
		Service: service,
	}
}

func (u *PostCommandUsecase) Save(payload *domain.Post) error {
	err := u.Service.Save(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *PostCommandUsecase) Edit(payload *domain.Post) error {
	err := u.Service.Edit(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *PostCommandUsecase) EditPartial(payload *domain.Post) error {
	err := u.Service.EditPartial(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *PostCommandUsecase) Remove(payload *domain.Post) error {
	err := u.Service.Remove(payload)
	if err != nil {
		return err
	}

	return nil
}
