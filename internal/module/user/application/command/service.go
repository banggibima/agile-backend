package command

import (
	"github.com/banggibima/agile-backend/internal/module/user"
	"github.com/banggibima/agile-backend/internal/module/user/domain"
)

type UserCommandService struct {
	Repository user.UserPostgresRepository
}

func NewUserCommandService(
	repository user.UserPostgresRepository,
) user.UserCommandService {
	return &UserCommandService{
		Repository: repository,
	}
}

func (s *UserCommandService) Save(payload *domain.User) error {
	return s.Repository.Save(payload)
}

func (s *UserCommandService) Edit(payload *domain.User) error {
	return s.Repository.Edit(payload)
}

func (s *UserCommandService) EditPartial(payload *domain.User) error {
	return s.Repository.EditPartial(payload)
}

func (s *UserCommandService) Remove(payload *domain.User) error {
	return s.Repository.Remove(payload)
}
