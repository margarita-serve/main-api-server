package types

import "encoding/json"

type CreateInferenceServiceRequest struct {
	InferenceServer        string
	Namespace              string       `json:"namespace" validate:"required"`     // 프로젝트 ID
	Inferencename          string       `json:"inferencename" validate:"required"` // 모델패키지 ID
	Predictor              *Predictor   `json:"predictor" `                        // 배포 명
	Transformer            *Transformer `json:"transformer,omitempty"`
	AutoscalingTargetCount string       `json:"autoscaling_target_count,omitempty"`
}

type Predictor struct {
	Modelspec *Modelspec `json:"modelspec,omitempty" ` // 프로젝트 ID
	Resource  *Resource  `json:"resource,omitempty"`
	Logger    string     `json:"logger,omitempty"`
}

type Modelspec struct {
	Modelframwwork string `json:"modelframwwork,omitempty" `
	Storageuri     string `json:"storageuri,omitempty" `
	RuntimeVersion string `json:"runtimeVersion,omitempty" `
}

type Resource struct {
	Requests *ResourceType `json:"requests,omitempty"` // 프로젝트 ID
	Limits   *ResourceType `json:"limits,omitempty"`   // 모델패키지 ID
}

type ResourceType struct {
	Cpu    string `json:"cpu,omitempty"`    // 프로젝트 ID
	Memory string `json:"memory,omitempty"` // 모델패키지 ID
	Gpu    string `json:"gpu,omitempty"`
}

type Transformer struct {
	Image    string        `json:"image,omitempty" ` // 프로젝트 ID
	Resource *ResourceType `json:"resource,omitempty"`
	Logger   string        `json:"logger,omitempty"`
}

// ToJSON covert to JSON
func (r *CreateInferenceServiceRequest) ToJSON() []byte {
	json, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return nil
	}
	return json
}
