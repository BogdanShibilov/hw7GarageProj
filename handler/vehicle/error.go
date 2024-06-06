package vehicle

import "errors"

var (
	InvalidIdErr     = errors.New("invalid vehicle id parameter")
	InvalidNameErr   = errors.New("invalid name parameter")
	InvalidBrandErr  = errors.New("invalid brand parameter")
	InvalidModelErr  = errors.New("invalid model parameter")
	AlreadyExistsErr = errors.New("vehicle already exists")
	NotFoundErr      = errors.New("vehicle not found")
)
