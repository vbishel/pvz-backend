package service

import (
	"auth-service/internal/config"
	"auth-service/internal/domain/user"
	"auth-service/pkg/apperrors"
	"context"
	"log/slog"
)

type usersService struct {
	cfg *config.Config
	usersRepository user.UserRepository
}

func NewUsersService(log *slog.Logger, cfg *config.Config, usersRepository user.UserRepository) *usersService {
	return &usersService{
		cfg: cfg,
		usersRepository: usersRepository,
	}
}

func (s *usersService) Create(ctx context.Context, email, hashPassword string, roleID int) (user.UserID, error) {
	u, err := s.usersRepository.Create(ctx, email, hashPassword, roleID)
	if err != nil {
		return 0, err
	}

	return u, nil
}

func (s *usersService) Find(ctx context.Context, id user.UserID) (user.User, error) {
	u, err := s.usersRepository.Find(ctx, id)
	if err != nil {
		return user.User{}, apperrors.ErrUserNotFound
	}

	return u, nil
}

func (s *usersService) FindByEmail(ctx context.Context, email string) (user.User, error) {
	u, err := s.usersRepository.FindByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}