package role

type RoleType string

const (
	ShareRole RoleType = "share_role"
	UserRole  RoleType = "user_role"
)

type RoleName string

type Role struct {
	RoleName RoleName
	RoleType RoleType
}
