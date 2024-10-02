package query

import (
	"github.com/banggibima/backend-agile/internal/module/todo"
	"github.com/banggibima/backend-agile/internal/module/todo/domain"
	"github.com/google/uuid"
)

type TodoQueryUsecase struct {
	Service todo.TodoQueryService
}

func NewTodoQueryUsecase(
	service todo.TodoQueryService,
) todo.TodoQueryUsecase {
	return &TodoQueryUsecase{
		Service: service,
	}
}

func (u *TodoQueryUsecase) Count() (int, error) {
	total, err := u.Service.Count()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (u *TodoQueryUsecase) Find(offset, size int, sort, order string) ([]*domain.Todo, error) {
	data, err := u.Service.Find(offset, size, sort, order)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *TodoQueryUsecase) FindByID(id uuid.UUID) (*domain.Todo, error) {
	data, err := u.Service.FindByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
