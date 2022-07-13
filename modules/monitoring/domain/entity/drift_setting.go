package entity

type DataDriftSetting struct {
	MonitorRange               string  `gorm:"size:256"`
	DriftMetricType            string  `gorm:"size:256"`
	DriftThreshold             float32 `gorm:"size:256"`
	ImportanceThreshold        float32 `gorm:"size:256"`
	LowImportanceAtRiskCount   int     `gorm:"size:8"`
	LowImportanceFailingCount  int     `gorm:"size:8"`
	HighImportanceAtRiskCount  int     `gorm:"size:8"`
	HighImportanceFailingCount int     `gorm:"size:8"`
}
