package query

import (
	"github.com/banggibima/agile-backend/internal/module/profile"
	"github.com/banggibima/agile-backend/internal/module/profile/domain"
	"github.com/google/uuid"
)

type ProfileQueryService struct {
	Repository profile.ProfilePostgresRepository
}

func NewProfileQueryService(
	repository profile.ProfilePostgresRepository,
) profile.ProfileQueryService {
	return &ProfileQueryService{
		Repository: repository,
	}
}

func (s *ProfileQueryService) Count() (int, error) {
	return s.Repository.Count()
}

func (s *ProfileQueryService) Find(page, size int, sort, order string) ([]*domain.Profile, error) {
	return s.Repository.Find(page, size, sort, order)
}

func (s *ProfileQueryService) FindByID(id uuid.UUID) (*domain.Profile, error) {
	return s.Repository.FindByID(id)
}
