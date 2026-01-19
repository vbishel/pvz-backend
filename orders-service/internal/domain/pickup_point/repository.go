package pickup_point

import "orders-service/internal/domain/city"

type PickupPointRepository interface {
	Create(cityID city.CityID) (PickupPointID, error)
	Find(id PickupPointID) (PickupPoint, error)
	Delete(id PickupPointID) error
}
