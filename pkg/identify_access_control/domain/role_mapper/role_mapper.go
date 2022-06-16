package role_mapper

type RoleMapperID string
type RoleName string

type TargetType string

const (
	User  TargetType = "share_role"
	Group TargetType = "user_role"
	Org   TargetType = "user_role"
)

type RoleMapper struct {
	ID         RoleMapperID
	RoleName   RoleName
	TargetType TargetType
}
