package command

import (
	"github.com/banggibima/agile-backend/internal/module/post"
	"github.com/banggibima/agile-backend/internal/module/post/domain"
)

type PostCommandService struct {
	Repository post.PostPostgresRepository
}

func NewPostCommandService(
	repository post.PostPostgresRepository,
) post.PostCommandService {
	return &PostCommandService{
		Repository: repository,
	}
}

func (s *PostCommandService) Save(payload *domain.Post) error {
	return s.Repository.Save(payload)
}

func (s *PostCommandService) Edit(payload *domain.Post) error {
	return s.Repository.Edit(payload)
}

func (s *PostCommandService) EditPartial(payload *domain.Post) error {
	return s.Repository.EditPartial(payload)
}

func (s *PostCommandService) Remove(payload *domain.Post) error {
	return s.Repository.Remove(payload)
}
