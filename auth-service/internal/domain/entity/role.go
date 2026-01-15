package entity

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	RoleModeratorId = 1
	RolePupEmployeeId = 2
	RoleClientId = 3
)
