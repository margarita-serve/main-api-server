package organization

type OrganizationID string
type UserID string
type GroupID string

type Organization struct {
	ID     OrganizationID
	Name   string
	Users  []UserID
	Groups []GroupID
}
