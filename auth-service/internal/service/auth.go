package service

import (
	"auth-service/config"
	"auth-service/internal/domain/user"
	"auth-service/internal/lib/jwt"
	"auth-service/internal/lib/logger/sl"
	"auth-service/pkg/apperrors"
	"context"
	"errors"
	"log/slog"
)

type authService struct {
	log         *slog.Logger
	cfg         *config.Config
	userFinder  UserFinder
	userCreator UserCreator
}

type UserCreator interface {
	Create(
		ctx context.Context,
		email string,
		passwordHash string,
	) (user.UserID, error)
}

type UserFinder interface {
	FindByEmail(
		ctx context.Context,
		email string,
	) (user.User, error)
}

func New(log *slog.Logger, cfg *config.Config, userFinder UserFinder, userCreator UserCreator) *authService {
	return &authService{
		cfg:         cfg,
		userFinder:  userFinder,
		userCreator: userCreator,
	}
}

func (s *authService) Register(ctx context.Context, email, password string) (user.UserID, error) {
	const op = "AuthService.Register"

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("attemtping to register user")

	u := user.User{
		Email:    email,
		Password: password,
	}
	u.HashPassword()

	uid, err := s.userCreator.Create(ctx, u.Email, u.Password)
	if err != nil {
		log.Error("failed to create user", sl.Err(err))
		return 0, err
	}

	log.Info("registered new user", "UserID", uid)

	return uid, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	const op = "AuthService.Login"

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("attempting to login user")

	u, err := s.userFinder.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			log.Warn("user not found", sl.Err(err))
			return "", apperrors.ErrUserIncorrectEmailOrPassword
		}

		log.Error("failed to get user", sl.Err(err))
		return "", err
	}

	if ok := u.CheckPassword(u.Password); !ok {
		return "", apperrors.ErrUserIncorrectEmailOrPassword
	}

	token, err := jwt.NewToken(&u, s.cfg.AccessToken.TTL, s.cfg.AccessToken.SigningKey)
	if err != nil {
		return "", err
	}

	log.Info("user logged in")

	return token, nil
}
