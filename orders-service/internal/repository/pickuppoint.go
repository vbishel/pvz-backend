package repository

import (
	"orders-service/internal/config"
	"orders-service/internal/domain/city"
	"orders-service/internal/domain/pickuppoint"
)

type pickuppointRepository struct {
	cfg *config.Config
}

func NewPickupPointRepository(cfg *config.Config) pickuppoint.PickupPointRepository {
	return &pickuppointRepository{cfg: cfg}
}

func (r *pickuppointRepository) Create(cityID city.CityID) (pickuppoint.PickupPointID, error) {
	panic("not implemented")
}

func (r *pickuppointRepository) Find(id pickuppoint.PickupPointID) (pickuppoint.PickupPoint, error) {
	panic("not implemented")
}

func (r *pickuppointRepository) Delete(id pickuppoint.PickupPointID) error {
	panic("not implemented")
}
