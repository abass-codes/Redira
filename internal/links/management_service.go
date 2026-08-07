package links

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ManagementService struct {
	queries *db.Queries
}

func NewManagementService(
	queries *db.Queries,
) *ManagementService {

	return &ManagementService{
		queries: queries,
	}
}

func convertUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}

func (s *ManagementService) Get(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
) (db.Link, error) {

	return s.queries.GetUserLinkByID(
		ctx,
		db.GetUserLinkByIDParams{
			ID:     convertUUID(id),
			UserID: convertUUID(userID),
		},
	)
}

func (s *ManagementService) Update(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
	url string,
) (db.Link, error) {

	return s.queries.UpdateLinkURL(
		ctx,
		db.UpdateLinkURLParams{
			OriginalUrl: url,
			ID:          convertUUID(id),
			UserID:      convertUUID(userID),
		},
	)
}

func (s *ManagementService) Disable(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
) (db.Link, error) {

	return s.queries.DisableLink(
		ctx,
		db.DisableLinkParams{
			ID:     convertUUID(id),
			UserID: convertUUID(userID),
		},
	)
}

func (s *ManagementService) Enable(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
) (db.Link, error) {

	return s.queries.EnableLink(
		ctx,
		db.EnableLinkParams{
			ID:     convertUUID(id),
			UserID: convertUUID(userID),
		},
	)
}
