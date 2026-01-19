package status

type StatusID int

type Status struct {
	ID   StatusID `json:"id"`
	Name string   `json:"name"`
}
