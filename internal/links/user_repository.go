package links

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(
	queries *db.Queries,
) *UserRepository {

	return &UserRepository{
		queries: queries,
	}
}

func toPgUUID(
	id uuid.UUID,
) pgtype.UUID {

	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}

func (r *UserRepository) CreateUserLink(
	ctx context.Context,
	userID uuid.UUID,
	url string,
	shortCode string,
) (db.Link, error) {

	return r.queries.CreateUserLink(
		ctx,
		db.CreateUserLinkParams{
			UserID:      toPgUUID(userID),
			OriginalUrl: url,
			ShortCode:   shortCode,
		},
	)
}

func (r *UserRepository) GetUserLinks(
	ctx context.Context,
	userID uuid.UUID,
) ([]db.Link, error) {

	return r.queries.GetUserLinks(
		ctx,
		toPgUUID(userID),
	)
}

func (r *UserRepository) DeleteUserLink(
	ctx context.Context,
	id uuid.UUID,
	userID uuid.UUID,
) error {

	return r.queries.DeleteUserLink(
		ctx,
		db.DeleteUserLinkParams{
			ID:     toPgUUID(id),
			UserID: toPgUUID(userID),
		},
	)
}
