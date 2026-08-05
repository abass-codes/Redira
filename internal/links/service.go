package links

import (
	"context"

	"github.com/abass-codes/redira/internal/analytics"
	"github.com/abass-codes/redira/internal/cache"
	db "github.com/abass-codes/redira/internal/database/db"
)

type Service struct {
	repository *Repository
	cache      *cache.Cache
	analytics  *analytics.Service
}

func NewService(
	repository *Repository,
	cache *cache.Cache,
	analyticsService *analytics.Service,
) *Service {
	return &Service{
		repository: repository,
		cache:      cache,
		analytics:  analyticsService,
	}
}

func (s *Service) Create(
	ctx context.Context,
	originalURL string,
) (*db.Link, error) {

	shortCode, err := GenerateUniqueCode(ctx)
	if err != nil {
		return nil, err
	}

	link, err := s.repository.Create(
		ctx,
		originalURL,
		shortCode,
	)

	if err != nil {
		return nil, err
	}

	_ = s.cache.StoreLink(ctx, link)

	return link, nil
}

func (s *Service) Redirect(
	ctx context.Context,
	shortCode string,
	ip string,
	userAgent string,
	referer string,
) (*db.Link, error) {

	var link *db.Link

	// Redis first
	cachedLink, err := s.cache.GetLink(ctx, shortCode)

	if err == nil {
		link = cachedLink
	} else {
		// PostgreSQL fallback
		link, err = s.repository.GetByShortCode(
			ctx,
			shortCode,
		)

		if err != nil {
			return nil, err
		}

		// Store in Redis
		_ = s.cache.StoreLink(ctx, link)
	}

	// Update click counter
	_ = s.repository.IncrementClicks(
		ctx,
		link.ID,
	)

	// Store analytics event
	if s.analytics != nil {
		_ = s.analytics.RecordClick(
			ctx,
			link.ID,
			ip,
			userAgent,
			referer,
		)
	}

	return link, nil
}
