package command

import (
	"github.com/banggibima/agile-backend/internal/module/todo"
	"github.com/banggibima/agile-backend/internal/module/todo/domain"
)

type TodoCommandUsecase struct {
	Service todo.TodoCommandService
}

func NewTodoCommandUsecase(
	service todo.TodoCommandService,
) todo.TodoCommandUsecase {
	return &TodoCommandUsecase{
		Service: service,
	}
}

func (u *TodoCommandUsecase) Save(payload *domain.Todo) error {
	err := u.Service.Save(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *TodoCommandUsecase) Edit(payload *domain.Todo) error {
	err := u.Service.Edit(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *TodoCommandUsecase) EditPartial(payload *domain.Todo) error {
	err := u.Service.EditPartial(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *TodoCommandUsecase) Remove(payload *domain.Todo) error {
	err := u.Service.Remove(payload)
	if err != nil {
		return err
	}

	return nil
}
