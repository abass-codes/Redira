package links

import (
	"context"

	"github.com/abass-codes/redira/internal/cache"
	db "github.com/abass-codes/redira/internal/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *Repository
	cache      *cache.Cache
}

func NewService(repository *Repository, cache *cache.Cache) *Service {
	return &Service{
		repository: repository,
		cache:      cache,
	}
}

func (s *Service) Create(ctx context.Context, originalURL string) (*db.Link, error) {
	shortCode, err := GenerateUniqueCode(ctx)
	if err != nil {
		return nil, err
	}

	link, err := s.repository.Create(ctx, originalURL, shortCode)
	if err != nil {
		return nil, err
	}

	// Cache newly created link
	_ = s.cache.StoreLink(ctx, link)

	return link, nil
}

func (s *Service) Redirect(ctx context.Context, shortCode string) (*db.Link, error) {

	// 1. Redis lookup
	if link, err := s.cache.GetLink(ctx, shortCode); err == nil {
		_ = s.repository.IncrementClicks(ctx, link.ID)
		return link, nil
	}

	// 2. PostgreSQL fallback
	link, err := s.repository.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	// 3. Store in Redis
	_ = s.cache.StoreLink(ctx, link)

	// 4. Increment clicks
	_ = s.repository.IncrementClicks(ctx, link.ID)

	return link, nil
}

func (s *Service) IncrementClicks(ctx context.Context, id pgtype.UUID) error {
	return s.repository.IncrementClicks(ctx, id)
}
