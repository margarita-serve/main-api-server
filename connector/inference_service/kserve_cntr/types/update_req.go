package types

import "encoding/json"

type UpdateInferenceServiceRequest struct {
	InferenceServer        string
	Namespace              string       `json:"namespace" validate:"required"`     // 프로젝트 ID
	Inferencename          string       `json:"inferencename" validate:"required"` // 모델패키지 ID
	Predictor              *Predictor   `json:"predictor" `                        // 배포 명
	Transformer            *Transformer `json:"transformer,omitempty"`
	AutoscalingTargetCount string       `json:"autoscaling_target_count,omitempty"`
}

// ToJSON covert to JSON
func (r *UpdateInferenceServiceRequest) ToJSON() []byte {
	json, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return json
}
