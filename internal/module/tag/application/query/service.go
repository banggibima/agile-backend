package query

import (
	"github.com/banggibima/agile-backend/internal/module/tag"
	"github.com/banggibima/agile-backend/internal/module/tag/domain"
	"github.com/google/uuid"
)

type TagQueryService struct {
	Repository tag.TagPostgresRepository
}

func NewTagQueryService(
	repository tag.TagPostgresRepository,
) tag.TagQueryService {
	return &TagQueryService{
		Repository: repository,
	}
}

func (s *TagQueryService) Count() (int, error) {
	return s.Repository.Count()
}

func (s *TagQueryService) Find(page, size int, sort, order string) ([]*domain.Tag, error) {
	return s.Repository.Find(page, size, sort, order)
}

func (s *TagQueryService) FindByID(id uuid.UUID) (*domain.Tag, error) {
	return s.Repository.FindByID(id)
}
