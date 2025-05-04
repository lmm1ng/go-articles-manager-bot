package user

import "errors"

var (
	ErrAlreadyExists = errors.New("Record already exists")
	ErrNotFound      = errors.New("Record not found")
)
