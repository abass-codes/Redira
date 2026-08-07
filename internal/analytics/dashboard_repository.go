package analytics

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type DashboardRepository struct {
	queries *db.Queries
}

func NewDashboardRepository(
	queries *db.Queries,
) *DashboardRepository {

	return &DashboardRepository{
		queries: queries,
	}
}

func (r *DashboardRepository) GetLinkAnalytics(
	ctx context.Context,
	id pgtype.UUID,
) ([]db.GetLinkAnalyticsRow, error) {

	return r.queries.GetLinkAnalytics(
		ctx,
		id,
	)
}

func (r *DashboardRepository) GetClickTimeline(
	ctx context.Context,
	id pgtype.UUID,
) ([]db.GetClickTimelineRow, error) {

	return r.queries.GetClickTimeline(
		ctx,
		id,
	)
}

func (r *DashboardRepository) GetDashboardSummary(
	ctx context.Context,
) (db.GetDashboardSummaryRow, error) {

	return r.queries.GetDashboardSummary(
		ctx,
	)
}
