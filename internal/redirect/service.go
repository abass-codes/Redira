package redirect

import (
	"context"
	"time"

	db "github.com/abass-codes/redira/internal/database/db"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	queries *db.Queries
	redis   *redis.Client
}

func NewService(
	queries *db.Queries,
	redis *redis.Client,
) *Service {

	return &Service{
		queries: queries,
		redis:   redis,
	}
}

func (s *Service) Redirect(
	ctx context.Context,
	shortCode string,
) (string, error) {

	cacheKey := "redirect:" + shortCode

	cached, err := s.redis.Get(
		ctx,
		cacheKey,
	).Result()

	if err == nil && cached != "" {

		return cached, nil
	}

	link, err := s.queries.GetRedirectLink(
		ctx,
		shortCode,
	)

	if err != nil {

		return "", ErrNotFound
	}

	if !link.Active {

		return "", ErrDisabled
	}

	if link.ExpiresAt.Valid {

		if link.ExpiresAt.Time.Before(time.Now()) {

			return "", ErrExpired
		}
	}

	// Analytics event tracking

	err = s.queries.CreateClickEvent(
		ctx,
		db.CreateClickEventParams{
			LinkID: link.ID,
		},
	)

	if err != nil {

		return link.OriginalUrl, nil
	}

	_ = s.redis.Set(
		ctx,
		cacheKey,
		link.OriginalUrl,
		time.Hour,
	).Err()

	return link.OriginalUrl, nil
}
