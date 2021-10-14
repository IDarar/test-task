package domain

import "errors"

var (
	ErrEmptyName         = errors.New("name cannot be empty")
	ErrIdNotSet          = errors.New("err id not set")
	ErrCreatingUser      = errors.New("err creating user")
	ErrUserAlreadyExists = errors.New("err user with such name already exists")
	ErrUserDoesNotExist  = errors.New("err user does not exist")
)
