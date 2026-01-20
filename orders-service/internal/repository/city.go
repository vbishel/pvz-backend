package repository

import (
	"orders-service/internal/config"
	"orders-service/internal/domain/city"
)

type cityRepository struct {
	cfg *config.Config
}

func NewCityRepository(cfg *config.Config) city.CityRepository {
	return &cityRepository{cfg: cfg}
}

func (r *cityRepository) Find(id city.CityID) (city.City, error) {
	panic("not implemented")
}

func (r *cityRepository) FindByName(name string) (city.City, error) {
	panic("not implemented")
}

func (r *cityRepository) Create(name string) (city.CityID, error) {
	panic("not implemented")
}

func (r *cityRepository) Delete(id city.CityID) error {
	panic("not implemented")
}
