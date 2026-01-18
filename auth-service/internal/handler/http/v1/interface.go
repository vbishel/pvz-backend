package v1

import (
	"auth-service/internal/domain/user"
	"context"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (user.UserID, error)
	Login(ctx context.Context, email, password string) (string, error)
}

type UsersService interface {
	Create(ctx context.Context, email, password string) (user.UserID, error)
	Find(ctx context.Context, id user.UserID) (user.User, error)
	FindByEmail(ctx context.Context, email string) (user.User, error)
}
