package types

type Predictor struct {
	Modelspec   *Modelspec `json:"modelspec,omitempty" ` // 프로젝트 ID
	Resource    *Resource  `json:"resource,omitempty"`
	Logger      string     `json:"logger,omitempty"`
	MinReplicas int        `json:"min_replicas,omitempty"`
	MaxReplicas int        `json:"max_replicas,omitempty"`
}

type Modelspec struct {
	Modelframwwork string `json:"modelframwwork,omitempty" `
	Storageuri     string `json:"storageuri,omitempty" `
	RuntimeVersion string `json:"runtime_version,omitempty" `
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
