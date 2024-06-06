package user

import "errors"

var (
	InvalidIdErr     = errors.New("invalid user id parameter")
	InvalidNameErr   = errors.New("invalid name parameter")
	AlreadyExistsErr = errors.New("user already exists")
	NotFoundErr      = errors.New("user not found")
)
