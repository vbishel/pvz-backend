package apperrors

import "errors"

var (
	ErrCityNotFound = errors.New("city not found")
	ErrCategoryNotFound = errors.New("category not found")
)