package apperrors

import "errors"

var (
	ErrInsufficientPermissions = errors.New("insufficient permissions")
)
