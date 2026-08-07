package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	db "github.com/abass-codes/redira/internal/database/db"
)

type Service struct {
	repository *Repository
}

func NewService(
	repository *Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) Register(
	ctx context.Context,
	email string,
	password string,
) (db.User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return db.User{}, err
	}

	return s.repository.CreateUser(
		ctx,
		email,
		string(passwordHash),
	)
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	user, err := s.repository.GetUserByEmail(
		ctx,
		email,
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return GenerateToken(
		user.ID.String(),
	)
}
