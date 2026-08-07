package analytics

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type DashboardService struct {
	repository *DashboardRepository
}

func NewDashboardService(
	repository *DashboardRepository,
) *DashboardService {

	return &DashboardService{
		repository: repository,
	}
}

func (s *DashboardService) GetLinkAnalytics(
	ctx context.Context,
	id uuid.UUID,
) ([]db.GetLinkAnalyticsRow, error) {

	var pgID pgtype.UUID

	copy(
		pgID.Bytes[:],
		id[:],
	)

	pgID.Valid = true

	return s.repository.GetLinkAnalytics(
		ctx,
		pgID,
	)
}

func (s *DashboardService) GetClickTimeline(
	ctx context.Context,
	id uuid.UUID,
) ([]db.GetClickTimelineRow, error) {

	var pgID pgtype.UUID

	copy(
		pgID.Bytes[:],
		id[:],
	)

	pgID.Valid = true

	return s.repository.GetClickTimeline(
		ctx,
		pgID,
	)
}

func (s *DashboardService) GetDashboardSummary(
	ctx context.Context,
) (db.GetDashboardSummaryRow, error) {

	return s.repository.GetDashboardSummary(
		ctx,
	)
}
