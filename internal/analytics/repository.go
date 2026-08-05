package analytics

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (r *Repository) CreateEvent(
	ctx context.Context,
	linkID pgtype.UUID,
	ip string,
	userAgent string,
	referer string,
) error {
	return r.queries.CreateClickEvent(ctx, db.CreateClickEventParams{
		LinkID:    linkID,
		IpAddress: stringPtr(ip),
		UserAgent: stringPtr(userAgent),
		Referer:   stringPtr(referer),
	})
}

func (r *Repository) GetAnalytics(
	ctx context.Context,
	linkID pgtype.UUID,
) ([]db.GetLinkAnalyticsRow, error) {
	return r.queries.GetLinkAnalytics(ctx, linkID)
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
