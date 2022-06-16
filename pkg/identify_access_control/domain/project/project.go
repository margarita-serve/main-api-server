package project

type ProjectID string
type UserID string
type GroupID string
type ComputeInstanceID string

type Project struct {
	ID                ProjectID
	Name              string
	Description       string
	ComputeInstanceID []ComputeInstanceID
}
