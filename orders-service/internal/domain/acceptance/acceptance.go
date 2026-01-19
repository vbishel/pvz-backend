package acceptance

import (
	"orders-service/internal/domain/pickup_point"
	"orders-service/internal/domain/status"
)

type AcceptanceID int

type Acceptance struct {
	ID       int                        `json:"id"`
	Date     string                     `json:"date"`
	PupID    pickup_point.PickupPointID `json:"pup_id"`
	StatusID status.StatusID            `json:"status_id"`
}
