package links

import (
	"context"

	"github.com/abass-codes/redira/internal/analytics"
	"github.com/abass-codes/redira/internal/cache"
	db "github.com/abass-codes/redira/internal/database/db"
	"github.com/google/uuid"
)

type Service struct {
	repository     *Repository
	userRepository *UserRepository
	cache          *cache.Cache
	analytics      *analytics.Service
}

func NewService(
	repository *Repository,
	userRepository *UserRepository,
	cache *cache.Cache,
	analytics *analytics.Service,
) *Service {

	return &Service{
		repository:     repository,
		userRepository: userRepository,
		cache:          cache,
		analytics:      analytics,
	}
}

func (s *Service) Create(
	ctx context.Context,
	url string,
	shortCode string,
	userID *uuid.UUID,
) (*db.Link, error) {

	var link *db.Link
	var err error

	if userID != nil {

		linkValue, err := s.userRepository.CreateUserLink(
			ctx,
			*userID,
			url,
			shortCode,
		)

		if err != nil {
			return nil, err
		}

		link = &linkValue

	} else {

		link, err = s.repository.Create(
			ctx,
			url,
			shortCode,
		)

		if err != nil {
			return nil, err
		}
	}

	if s.cache != nil {
		_ = s.cache.StoreLink(
			ctx,
			link,
		)
	}

	return link, nil
}

func (s *Service) Redirect(
	ctx context.Context,
	shortCode string,
) (*db.Link, error) {

	if link, err := s.cache.GetLink(
		ctx,
		shortCode,
	); err == nil {

		return link, nil
	}

	link, err := s.repository.GetByShortCode(
		ctx,
		shortCode,
	)

	if err != nil {
		return nil, err
	}

	_ = s.cache.StoreLink(
		ctx,
		link,
	)

	return link, nil
}
