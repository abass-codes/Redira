package analytics

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) RecordClick(
	ctx context.Context,
	linkID pgtype.UUID,
	ip string,
	userAgent string,
	referer string,
) error {
	return s.repository.CreateEvent(
		ctx,
		linkID,
		ip,
		userAgent,
		referer,
	)
}

func (s *Service) GetAnalytics(
	ctx context.Context,
	linkID pgtype.UUID,
) ([]db.GetLinkAnalyticsRow, error) {
	return s.repository.GetAnalytics(ctx, linkID)
}
