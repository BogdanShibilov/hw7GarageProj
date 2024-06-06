package user

import "errors"

var (
	InvalidInputErr  = errors.New("invalid input")
	AlreadyExistsErr = errors.New("user already exists")
	NotFoundErr      = errors.New("user not found")
)
