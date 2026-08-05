package links

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

func (r *Repository) Create(ctx context.Context, originalURL, shortCode string) (*db.Link, error) {
	link, err := r.queries.CreateLink(ctx, db.CreateLinkParams{
		OriginalUrl: originalURL,
		ShortCode:   shortCode,
	})
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *Repository) GetByShortCode(ctx context.Context, shortCode string) (*db.Link, error) {
	link, err := r.queries.GetLinkByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *Repository) IncrementClicks(ctx context.Context, id pgtype.UUID) error {
	return r.queries.IncrementClickCount(ctx, id)
}
