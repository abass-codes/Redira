package redirect

import (
	"context"
	"log"
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

	// Redis lookup

	cached, err := s.redis.Get(
		ctx,
		cacheKey,
	).Result()

	if err == nil && cached != "" {

		log.Println("REDIS HIT:", shortCode)

		return cached, nil
	}

	log.Println("REDIS MISS:", shortCode)

	// Database lookup

	link, err := s.queries.GetRedirectLink(
		ctx,
		shortCode,
	)

	if err != nil {

		log.Println("DATABASE ERROR:", err)

		return "", ErrNotFound
	}

	log.Println(
		"DATABASE FOUND:",
		link.ShortCode,
		link.OriginalUrl,
	)

	// Check active status

	if !link.Active {

		return "", ErrDisabled
	}

	// Check expiration

	if link.ExpiresAt.Valid {

		if link.ExpiresAt.Time.Before(
			time.Now(),
		) {

			return "", ErrExpired
		}
	}

	// Store in Redis

	err = s.redis.Set(
		ctx,
		cacheKey,
		link.OriginalUrl,
		time.Hour,
	).Err()

	if err != nil {

		log.Println("REDIS SET ERROR:", err)
	}

	return link.OriginalUrl, nil
}
