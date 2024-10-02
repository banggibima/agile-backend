package query

import (
	"github.com/banggibima/backend-agile/internal/module/todo"
	"github.com/banggibima/backend-agile/internal/module/todo/domain"
	"github.com/google/uuid"
)

type TodoQueryService struct {
	Repository todo.TodoPostgresRepository
}

func NewTodoQueryService(
	repository todo.TodoPostgresRepository,
) todo.TodoQueryService {
	return &TodoQueryService{
		Repository: repository,
	}
}

func (s *TodoQueryService) Count() (int, error) {
	return s.Repository.Count()
}

func (s *TodoQueryService) Find(page, size int, sort, order string) ([]*domain.Todo, error) {
	return s.Repository.Find(page, size, sort, order)
}

func (s *TodoQueryService) FindByID(id uuid.UUID) (*domain.Todo, error) {
	return s.Repository.FindByID(id)
}
