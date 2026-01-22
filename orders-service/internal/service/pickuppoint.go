package pickuppoint

import (
	"context"
	"log/slog"
	"orders-service/internal/config"
	"orders-service/internal/domain/city"
	"orders-service/internal/domain/pickuppoint"
	"orders-service/internal/domain/role"
	"orders-service/pkg/apperrors"
)

type pupService struct {
	cfg           *config.Config
	log           *slog.Logger
	pupRepository pickuppoint.PickupPointRepository
}

func NewPupService(cfg *config.Config, log *slog.Logger, pupRepository pickuppoint.PickupPointRepository) *pupService {
	return &pupService{
		cfg:           cfg,
		log:           log,
		pupRepository: pupRepository,
	}
}

func (s *pupService) Create(ctx context.Context, cityID city.CityID, roleID int) (pickuppoint.PickupPointID, error){
	if (roleID != role.RoleModeratorId) {
		return 0, apperrors.ErrInsufficientPermissions
	}

	id, err := s.pupRepository.Create(cityID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *pupService) Find(ctx context.Context, id pickuppoint.PickupPointID) (pickuppoint.PickupPoint, error) {
	p, err := s.pupRepository.Find(id)

	if err != nil {
		return pickuppoint.PickupPoint{}, err
	}

	return p, nil
}

