package dto

type CreateProjectRequestDTO struct {
	Name        string `json:"name" example:"Project 01" extensions:"x-order=1"` // 프로젝트 명
	Description string `json:"description" example:"New Project" extensions:"x-order=2"`
}

type CreateProjectResponseDTO struct {
	ProjectID string
}

type GetProjectRequestDTO struct {
	ProjectID string `json:"projectID" validate:"false" swaggerignore:"true"` // 프로젝트 ID
}

type GetProjectResponseDTO struct {
	ProjectID     string
	Name          string
	Description   string
	ModelPackages interface{}
	Deployments   interface{}
}

type GetProjectListRequestDTO struct {
	Name  string `json:"name" extensions:"x-order=1"`                  // 검색조건: 프로젝트 명
	Limit int    `json:"limit" extensions:"x-order=2"`                 // 한번에 조회 할 건수
	Page  int    `json:"page" extensions:"x-order=3"`                  // 조회 할 페이지, 첫 조회후 TotalPages 범위 내에서 선택 후 보낸다
	Sort  string `enums:"CreateAsc,CreateDesc" extensions:"x-order=4"` //정열방식, CreateAsc: 생성시간 내림차순, CraeteDesc: 생성시간 역차순
}

type GetProjectListResponseDTO struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

type DeleteProjectRequestDTO struct {
	ProjectID string
}

type DeleteProjectResponseDTO struct {
	Message string
}

type UpdateProjectRequestDTO struct {
	ProjectID   string `json:"projectID" validate:"false" swaggerignore:"true"`                     // 프로젝트 ID
	Name        string `json:"name" example:"Edited Project 01" extensions:"x-order=1"`             // 프로젝트 명
	Description string `json:"description" example:"this is Edited Project" extensions:"x-order=2"` // 프로젝트 설명
}

type UpdateProjectResponseDTO struct {
	Message string
}
