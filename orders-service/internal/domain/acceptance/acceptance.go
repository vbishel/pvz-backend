package acceptance

import (
	"orders-service/internal/domain/pickuppoint"
	"orders-service/internal/domain/status"
)

type AcceptanceID int

type Acceptance struct {
	ID       int                        `json:"id"`
	Date     string                     `json:"date"`
	PupID    pickuppoint.PickupPointID `json:"pup_id"`
	StatusID status.StatusID            `json:"status_id"`
}
