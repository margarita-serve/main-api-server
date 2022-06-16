package types

type GetInferenceServiceRequest struct {
	InferenceServer string
	Namespace       string `json:"namespace" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"`     //
	Inferencename   string `json:"inferencename" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"` //
}
