package service

import (
	"auth-service/config"
	"auth-service/internal/domain/role"
	"auth-service/internal/domain/user"
	"auth-service/internal/lib/jwt"
	"auth-service/pkg/apperrors"
	"context"
	"errors"
	"log/slog"
)

type authService struct {
	cfg         *config.Config
	userFinder  UserFinder
	userCreator UserCreator
}

type UserCreator interface {
	Create(
		ctx context.Context,
		email string,
		passwordHash string,
		roleID int,
	) (user.UserID, error)
}

type UserFinder interface {
	FindByEmail(
		ctx context.Context,
		email string,
	) (user.User, error)
}

func NewAuthService(log *slog.Logger, cfg *config.Config, userFinder UserFinder, userCreator UserCreator) *authService {
	return &authService{
		cfg:         cfg,
		userFinder:  userFinder,
		userCreator: userCreator,
	}
}

func (s *authService) Register(ctx context.Context, email, password string) (user.UserID, error) {
	u := user.User{
		Email:    email,
		Password: password,
	}
	u.HashPassword()

	uid, err := s.userCreator.Create(ctx, u.Email, u.Password, role.RoleClientId)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.userFinder.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return "", apperrors.ErrUserIncorrectEmailOrPassword
		}

		return "", err
	}

	if ok := u.CheckPassword(u.Password); !ok {
		return "", apperrors.ErrUserIncorrectEmailOrPassword
	}

	token, err := jwt.NewToken(&u, s.cfg.AccessToken.TTL, s.cfg.AccessToken.SigningKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
