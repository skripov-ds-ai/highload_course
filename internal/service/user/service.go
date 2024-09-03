package user

import (
	"context"
	"github.com/skripov-ds-ai/highload_course/internal/db/repository/user"
	"golang.org/x/crypto/bcrypt"

	models "github.com/skripov-ds-ai/highload_course/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, userID string) (models.User, error)
	ListByPrefixFirstNameSecondName(ctx context.Context, firstName, secondName string) ([]models.User, error)
	Create(ctx context.Context, user models.User) (string, error)
	Update(ctx context.Context, user models.UpdateUserParams) error
	Delete(ctx context.Context, userID string) error
}

type UserService interface {
	Get(ctx context.Context, userID string) (models.User, error)
	Register(ctx context.Context, params models.CreateUserParams) (string, error)
	ListByPrefixFirstNameSecondName(ctx context.Context, firstName, secondName string) (models.Users, error)
}

type Service struct {
	repository user.Repository
}

func NewUserService(repository user.Repository) UserService {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Get(ctx context.Context, userID string) (models.User, error) {
	return s.repository.Get(ctx, userID)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *Service) Register(ctx context.Context, params models.CreateUserParams) (string, error) {
	passwordHash, err := hashPassword(params.Password)
	if err != nil {
		return "", err
	}
	return s.repository.Create(ctx, models.User{
		PasswordHash: passwordHash,
		FirstName:    params.FirstName,
		SecondName:   params.SecondName,
		Birthdate:    params.Birthdate,
		Biography:    params.Biography,
		City:         params.City,
	})
}

func (s *Service) ListByPrefixFirstNameSecondName(ctx context.Context, firstName, secondName string) (models.Users, error) {
	return s.repository.ListByPrefixFirstNameSecondName(ctx, firstName, secondName)
}
