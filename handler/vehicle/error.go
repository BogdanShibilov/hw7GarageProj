package vehicle

import "errors"

var (
	InvalidInputErr  = errors.New("invalid input")
	AlreadyExistsErr = errors.New("vehicle already exists")
	NotFoundErr      = errors.New("vehicle not found")
)
