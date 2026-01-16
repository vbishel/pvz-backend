package auth

import (
	"auth-service/config"
	"auth-service/internal/domain/entity"
	"auth-service/internal/lib/jwt"
	"auth-service/pkg/apperrors"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, u entity.User) (entity.UserID, error)
	Find(ctx context.Context, id entity.UserID) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
}

type authService struct {
	cfg            *config.Config
	userRepository UserRepository
}

func NewAuthService(userRepository UserRepository, cfg *config.Config) *authService {
	return &authService{
		userRepository: userRepository,
		cfg:            cfg,
	}
}

func (s *authService) Register(ctx context.Context, registerDTO RegisterDTO) (entity.UserID, error) {
	u := entity.User{
		Email:    registerDTO.email,
		Password: registerDTO.password,
	}
	u.HashPassword()

	uid, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (s *authService) Login(ctx context.Context, loginDTO LoginDTO) (string, error) {
	u, err := s.userRepository.FindByEmail(ctx, loginDTO.email)
	if err != nil {
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
