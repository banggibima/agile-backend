package command

import (
	"github.com/banggibima/agile-backend/internal/module/profile"
	"github.com/banggibima/agile-backend/internal/module/profile/domain"
)

type ProfileCommandService struct {
	Repository profile.ProfilePostgresRepository
}

func NewProfileCommandService(
	repository profile.ProfilePostgresRepository,
) profile.ProfileCommandService {
	return &ProfileCommandService{
		Repository: repository,
	}
}

func (s *ProfileCommandService) Save(payload *domain.Profile) error {
	return s.Repository.Save(payload)
}

func (s *ProfileCommandService) Edit(payload *domain.Profile) error {
	return s.Repository.Edit(payload)
}

func (s *ProfileCommandService) EditPartial(payload *domain.Profile) error {
	return s.Repository.EditPartial(payload)
}

func (s *ProfileCommandService) Remove(payload *domain.Profile) error {
	return s.Repository.Remove(payload)
}
