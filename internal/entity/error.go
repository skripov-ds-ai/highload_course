package entity

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrWrongPassword = errors.New("wrong password")
)
