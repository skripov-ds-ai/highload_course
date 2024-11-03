package user

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"

	models "github.com/skripov-ds-ai/highload_course/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, userID string) (models.User, error)
	ListByPrefixFirstNameSecondName(ctx context.Context, firstName, secondName string) ([]models.User, error)
	Create(ctx context.Context, user models.User) (string, error)
	Update(ctx context.Context, user models.UpdateUserParams) error
	Delete(ctx context.Context, userID string) error
}

type Storage struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Get(ctx context.Context, userID string) (user models.User, err error) {
	args := map[string]any{
		"id": userID,
	}
	q := `
select
	*
from public.users
where id = :id`
	res, err := s.db.NamedQueryContext(ctx, q, args)
	if err != nil {
		return models.User{}, err
	}
	ok := res.Next()
	if !ok {
		return models.User{}, models.ErrNotFound
	}
	defer func() {
		e := res.Close()
		if e != nil {
			err = errors.Join(err, e)
		}
	}()
	err = res.StructScan(&user)
	if err != nil {
		return models.User{}, err
	}
	err = res.Err()
	return user, err
}

func (s *Storage) ListByPrefixFirstNameSecondName(ctx context.Context, firstName, secondName string) (users []models.User, err error) {
	users = make([]models.User, 0)
	args := map[string]any{
		"first_name":  firstName + "%",
		"second_name": secondName + "%",
	}
	q := `
select
	u.id,
	u.first_name,
	u.second_name,
	u.birthdate,
	u.biography,
	u.city,
	u.gender
from public.users u
where first_name like :first_name and second_name like :second_name
order by id
;`
	res, err := s.db.NamedQueryContext(ctx, q, args)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := res.Close()
		if e != nil {
			err = errors.Join(err, e)
		}
	}()
	for res.Next() {
		user := models.User{}
		err = res.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *Storage) Create(ctx context.Context, user models.User) (userID string, err error) {
	args := map[string]any{
		"first_name":    user.FirstName,
		"second_name":   user.SecondName,
		"birthdate":     user.Birthdate,
		"biography":     user.Biography,
		"city":          user.City,
		"password_hash": user.PasswordHash,
	}
	q := `
insert into public.users
    (first_name, second_name, birthdate, biography, city, password_hash)
values
	(:first_name, :second_name, :birthdate, :biography, :city, :password_hash)
returning id`
	res, err := s.db.NamedQueryContext(ctx, q, args)
	if err != nil {
		return "", err
	}
	defer func() {
		e := res.Close()
		if e != nil {
			err = errors.Join(err, e)
		}
	}()
	for res.Next() {
		err = res.Scan(&userID)
		if err != nil {
			return "", err
		}
	}
	err = res.Err()
	return userID, err
}

func (s *Storage) Update(ctx context.Context, user models.UpdateUserParams) error {
	return nil
}

func (s *Storage) Delete(ctx context.Context, userID string) error {
	return nil
}
