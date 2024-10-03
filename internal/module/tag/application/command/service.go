package command

import (
	"github.com/banggibima/agile-backend/internal/module/tag"
	"github.com/banggibima/agile-backend/internal/module/tag/domain"
)

type TagCommandService struct {
	Repository tag.TagPostgresRepository
}

func NewTagCommandService(
	repository tag.TagPostgresRepository,
) tag.TagCommandService {
	return &TagCommandService{
		Repository: repository,
	}
}

func (s *TagCommandService) Save(payload *domain.Tag) error {
	return s.Repository.Save(payload)
}

func (s *TagCommandService) Edit(payload *domain.Tag) error {
	return s.Repository.Edit(payload)
}

func (s *TagCommandService) EditPartial(payload *domain.Tag) error {
	return s.Repository.EditPartial(payload)
}

func (s *TagCommandService) Remove(payload *domain.Tag) error {
	return s.Repository.Remove(payload)
}
