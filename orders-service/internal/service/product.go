package service

import (
	"context"
	"log/slog"
	"orders-service/internal/config"
	"orders-service/internal/domain/product"
	"orders-service/internal/domain/role"
	"orders-service/pkg/apperrors"
)

type productService struct {
	cfg               *config.Config
	log               *slog.Logger
	productRepository product.ProductRepository
}

func NewProductService(cfg *config.Config, log *slog.Logger, productRepository product.ProductRepository) *productService {
	return &productService{
		cfg:               cfg,
		log:               log,
		productRepository: productRepository,
	}
}

func (s *productService) Create(ctx context.Context, name string, roleID int) (product.ProductID, error) {
	if roleID != role.RoleModeratorId {
		return 0, apperrors.ErrInsufficientPermissions
	}

	pid, err := s.productRepository.Create(name)
	if err != nil {
		return 0, err
	}

	return pid, nil
}
