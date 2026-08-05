package links

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, originalURL string) (*db.Link, error) {
	shortCode, err := GenerateUniqueCode(ctx)
	if err != nil {
		return nil, err
	}

	return s.repository.Create(ctx, originalURL, shortCode)
}

func (s *Service) Redirect(ctx context.Context, shortCode string) (*db.Link, error) {
	link, err := s.repository.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	if err := s.repository.IncrementClicks(ctx, link.ID); err != nil {
		return nil, err
	}

	return link, nil
}
