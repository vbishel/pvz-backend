package acceptance

import (
	"orders-service/internal/domain/pickuppoint"
	"orders-service/internal/domain/status"
)

type AcceptanceRepository interface {
	Create(pupID pickuppoint.PickupPointID) (AcceptanceID, error)
	SetStatus(acceptanceID AcceptanceID, status status.StatusID) (error)
}
