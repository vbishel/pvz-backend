package category

type CategoryID int

type Category struct {
	ID   CategoryID `json:"category_id"`
	Name string     `json:"name"`
}
