package product

type ProductID int

type Product struct {
	ID             ProductID `json:"product_id"`
	Name           string    `json:"name"`
	AcceptanceDate string    `json:"acceptance_date"`
	CategoryID     string    `json:"category_id"`
}
