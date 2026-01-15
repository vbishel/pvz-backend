package entity

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	RoleModerator   = "moderator"
	RolePupEmployee = "pup_employee"
	RoleClient      = "client"
)

