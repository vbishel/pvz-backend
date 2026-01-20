package pickuppoint

type PickupPointID int

type PickupPoint struct {
	ID        PickupPointID `json:"id"`
	CreatedAt string        `json:"created_at"`
	City      string        `json:"city"`
}
