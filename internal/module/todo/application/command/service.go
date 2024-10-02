package command

import (
	"github.com/banggibima/backend-agile/internal/module/todo"
	"github.com/banggibima/backend-agile/internal/module/todo/domain"
)

type TodoCommandService struct {
	Repository todo.TodoPostgresRepository
}

func NewTodoCommandService(
	repository todo.TodoPostgresRepository,
) todo.TodoCommandService {
	return &TodoCommandService{
		Repository: repository,
	}
}

func (s *TodoCommandService) Save(payload *domain.Todo) error {
	return s.Repository.Save(payload)
}

func (s *TodoCommandService) Edit(payload *domain.Todo) error {
	return s.Repository.Edit(payload)
}

func (s *TodoCommandService) EditPartial(payload *domain.Todo) error {
	return s.Repository.EditPartial(payload)
}

func (s *TodoCommandService) Remove(payload *domain.Todo) error {
	return s.Repository.Remove(payload)
}
