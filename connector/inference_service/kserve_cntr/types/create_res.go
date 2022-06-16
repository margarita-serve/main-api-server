package types

type CreateInferenceServiceResponse struct {
	Message       string `json:"projectID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"`      // 프로젝트 ID
	Inferencename string `json:"modelPackageID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"` // 모델패키지 ID
}
