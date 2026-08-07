package auth

import (
	"context"

	db "github.com/abass-codes/redira/internal/database/db"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(
	queries *db.Queries,
) *Repository {
	return &Repository{
		queries: queries,
	}
}

func (r *Repository) CreateUser(
	ctx context.Context,
	email string,
	passwordHash string,
) (db.User, error) {

	return r.queries.CreateUser(
		ctx,
		db.CreateUserParams{
			Email:        email,
			PasswordHash: passwordHash,
		},
	)
}

func (r *Repository) GetUserByEmail(
	ctx context.Context,
	email string,
) (db.User, error) {

	return r.queries.GetUserByEmail(
		ctx,
		email,
	)
}
