package group

type GroupID string
type UserID string

type Group struct {
	ID    GroupID
	Name  string
	Users []UserID
}
