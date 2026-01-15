package apperrors

import "errors"

var (
	ErrUserAlreadyExists            = errors.New("user already exists")
	ErrUserNotFound                 = errors.New("user not found")
	ErrUserNotArchived              = errors.New("user cannot be archived")
	ErrUserIncorrectEmailOrPassword = errors.New("incorrect email or password")
	ErrUserPasswordNotGenerated     = errors.New("password generation error")
	ErrUserContextNotFound          = errors.New("user not found in context")
)
