package query

import (
	"github.com/banggibima/agile-backend/internal/module/post"
	"github.com/banggibima/agile-backend/internal/module/post/domain"
	"github.com/google/uuid"
)

type PostQueryService struct {
	Repository post.PostPostgresRepository
}

func NewPostQueryService(
	repository post.PostPostgresRepository,
) post.PostQueryService {
	return &PostQueryService{
		Repository: repository,
	}
}

func (s *PostQueryService) Count() (int, error) {
	return s.Repository.Count()
}

func (s *PostQueryService) Find(page, size int, sort, order string) ([]*domain.Post, error) {
	return s.Repository.Find(page, size, sort, order)
}

func (s *PostQueryService) FindByID(id uuid.UUID) (*domain.Post, error) {
	return s.Repository.FindByID(id)
}
