package query

import (
	"github.com/banggibima/backend-agile/internal/module/user"
	"github.com/banggibima/backend-agile/internal/module/user/domain"
	"github.com/google/uuid"
)

type UserQueryService struct {
	Repository user.UserPostgresRepository
}

func NewUserQueryService(
	repository user.UserPostgresRepository,
) user.UserQueryService {
	return &UserQueryService{
		Repository: repository,
	}
}

func (s *UserQueryService) Count() (int, error) {
	return s.Repository.Count()
}

func (s *UserQueryService) Find(page, size int, sort, order string) ([]*domain.User, error) {
	return s.Repository.Find(page, size, sort, order)
}

func (s *UserQueryService) FindByID(id uuid.UUID) (*domain.User, error) {
	return s.Repository.FindByID(id)
}
