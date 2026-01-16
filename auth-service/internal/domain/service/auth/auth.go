package auth

import (
	"auth-service/config"
	"auth-service/internal/domain/entity"
	"auth-service/internal/lib/jwt"
	"auth-service/internal/lib/logger/sl"
	"auth-service/pkg/apperrors"
	"context"
	"errors"
	"log/slog"
)

type UserRepository interface {
	Create(ctx context.Context, u entity.User) (entity.UserID, error)
	Find(ctx context.Context, id entity.UserID) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
}

type authService struct {
	log            *slog.Logger
	cfg            *config.Config
	userRepository UserRepository
}

func New(userRepository UserRepository, cfg *config.Config) *authService {
	return &authService{
		userRepository: userRepository,
		cfg:            cfg,
	}
}

func (s *authService) Register(ctx context.Context, registerDTO RegisterDTO) (entity.UserID, error) {
	const op = "AuthService.Register"

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", registerDTO.email),
	)

	log.Info("attemtping to register user")

	u := entity.User{
		Email:    registerDTO.email,
		Password: registerDTO.password,
	}
	u.HashPassword()

	uid, err := s.userRepository.Create(ctx, u)
	if err != nil {
		log.Error("failed to create user", sl.Err(err))
		return 0, err
	}

	log.Info("registered new user", "UserID", uid)

	return uid, nil
}

func (s *authService) Login(ctx context.Context, loginDTO LoginDTO) (string, error) {
	const op = "AuthService.Login"

	log := s.log.With(
		slog.String("op", op),
		slog.String("email", loginDTO.email),
	)

	log.Info("attempting to login user")

	u, err := s.userRepository.FindByEmail(ctx, loginDTO.email)
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
