package auth

import (
	"context"
	"github.com/google/uuid"
	models "github.com/skripov-ds-ai/highload_course/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Get(ctx context.Context, userID string) (models.User, error)
}

type AuthService interface {
	Login(ctx context.Context, userID, password string) (string, error)
}

type Service struct {
	repository Repository
}

func NewAuthService(repository Repository) AuthService {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Login(ctx context.Context, userID, password string) (string, error) {
	u, err := s.repository.Get(ctx, userID)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", models.ErrWrongPassword
	}

	// TODO: store session in Redis/DB
	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}
