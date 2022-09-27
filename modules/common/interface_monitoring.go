package common

type IMonitorService interface {
	GetByIDInternal(monitoringID string) (*MonitorGetByIDInternalResponseDTO, error)
}

type MonitorGetByIDInternalResponseDTO struct {
	// drift, accuracy 생성 상태값 추가
	ID                     string `gorm:"size:256"`
	ModelPackageID         string `gorm:"size:256"`
	FeatureDriftTracking   bool
	AccuracyMonitoring     bool
	AssociationID          string `gorm:"size:256"`
	AssociationIDInFeature bool   `gorm:"size:256"`
	DriftStatus            string `gorm:"size:256"`
	AccuracyStatus         string `gorm:"size:256"`
	ServiceHealthStatus    string `gorm:"size:256"`
	DriftCreated           bool
	AccuracyCreated        bool
	ServiceHealthCreated   bool
}
