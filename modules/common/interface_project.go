package common

type GetProjectInternalResponseDTO struct {
	ProjectID   string
	Name        string
	Description string
	CreatedBy   string
	CreatedAt   string
}

type GetProjectListResponseDTO struct {
	Rows []GetProjectInternalResponseDTO
}

type IProjectService interface {
	GetListInternal(userID string) (*GetProjectListResponseDTO, error)
	GetByIDInternal(projectID string) (*GetProjectInternalResponseDTO, error)
}
