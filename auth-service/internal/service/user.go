package service

import (
	"auth-service/config"
	"auth-service/internal/domain/user"
	"auth-service/internal/lib/logger/sl"
	"auth-service/pkg/apperrors"
	"context"
	"log/slog"
)

type usersService struct {
	log *slog.Logger
	cfg *config.Config
	usersRepository user.UserRepository
}

func NewUsersService(log *slog.Logger, cfg *config.Config, usersRepository user.UserRepository) *usersService {
	return &usersService{
		log: log,
		cfg: cfg,
		usersRepository: usersRepository,
	}
}

func (s *usersService) Create(ctx context.Context, email, hashPassword string, roleID int) (user.UserID, error) {
	const op = "usersService.Create"

	log := s.log.With(
		slog.String("op", op),
	)

	u, err := s.usersRepository.Create(ctx, email, hashPassword, roleID)
	if err != nil {
		log.Error("failed to create user", sl.Err(err))
		return 0, err
	}

	return u, nil
}

func (s *usersService) Find(ctx context.Context, id user.UserID) (user.User, error) {
	const op = "usersService.Find"

	log := s.log.With(
		slog.String("op", op),
	)

	u, err := s.usersRepository.Find(ctx, id)
	if err != nil {
		log.Warn("failed to find user by id", sl.Err(err))
		return user.User{}, apperrors.ErrUserNotFound
	}

	return u, nil
}

func (s *usersService) FindByEmail(ctx context.Context, email string) (user.User, error) {
	const op = "usersService.FindByEmail"

	log := s.log.With(
		slog.String("op", op),
	)

	u, err := s.usersRepository.FindByEmail(ctx, email)
	if err != nil {
		log.Warn("failed to find user by email", sl.Err(err))
		return user.User{}, err
	}

	return u, nil
}