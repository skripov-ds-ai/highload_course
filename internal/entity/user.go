package entity

import (
	"github.com/google/uuid"
	"github.com/skripov-ds-ai/highload_course/internal/generated"
	"time"
)

type RegisterUserResponse struct {
	UserID string `json:"user_id"`
}

func NewCreateUserParams(b generated.PostUserRegisterJSONBody) CreateUserParams {
	return CreateUserParams{
		Biography:  b.Biography,
		City:       b.City,
		FirstName:  b.FirstName,
		SecondName: b.SecondName,
		Gender:     b.Gender,
		Password:   b.Password,
	}
}

type CreateUserParams struct {
	Biography  *string
	Birthdate  time.Time
	City       *string
	FirstName  string
	SecondName string
	Gender     *string
	Password   string
}

type UpdateUserParams struct {
	ID           string
	PasswordHash *string
	FirstName    *string
	SecondName   *string
	Birthdate    *time.Time
	Biography    *string
	City         *string
}

type User struct {
	ID           uuid.UUID `db:"id"`
	PasswordHash string    `db:"password_hash"`
	FirstName    string    `db:"first_name"`
	SecondName   string    `db:"second_name"`
	Birthdate    time.Time `db:"birthdate"`
	Biography    *string   `db:"biography"`
	City         *string   `db:"city"`
}

// TODO: TO User from generated
type UserJson struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Birthdate  time.Time `json:"birthdate"`
	Biography  *string   `json:"biography"`
	City       *string   `json:"city"`
}

func (u *User) ToModel() UserJson {
	return UserJson{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Birthdate:  u.Birthdate,
		Biography:  u.Biography,
		City:       u.City,
	}
}

type Users []User

func (u Users) ToModel() []UserJson {
	users := make([]UserJson, len(u))
	for i, user := range u {
		users[i] = user.ToModel()
	}
	return users
}

type UserToken struct {
	Token string `json:"token"`
}
