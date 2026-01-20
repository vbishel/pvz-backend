package repository

import (
	"orders-service/internal/config"
	"orders-service/internal/domain/acceptance"
	"orders-service/internal/domain/pickuppoint"
	"orders-service/internal/domain/status"
)

type acceptanceRepository struct {
	cfg *config.Config
}

func NewAcceptanceRepository(cfg *config.Config) acceptance.AcceptanceRepository {
	return &acceptanceRepository{cfg: cfg}
}

func (r *acceptanceRepository) Create(pupID pickuppoint.PickupPointID) (acceptance.AcceptanceID, error) {
	panic("not implemented")
}

func (r *acceptanceRepository) SetStatus(id acceptance.AcceptanceID, statusId status.StatusID) (error) {
	panic("not implemented")
}
