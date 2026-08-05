package links

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"
	"github.com/abass-codes/redira/internal/utils"
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
	shortCode, err := utils.GenerateShortCode(6)
	if err != nil {
		return nil, err
	}

	return s.repository.Create(ctx, originalURL, shortCode)
}
