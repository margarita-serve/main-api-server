package entity

type AccuracySetting struct {
	MetricType   string  `gorm:"size:256"`
	Measurement  string  `gorm:"size:256"`
	AtRiskValue  float32 `gorm:"size:8"`
	FailingValue float32 `gorm:"size:8"`
}
