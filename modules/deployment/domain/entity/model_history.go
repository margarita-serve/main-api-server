package entity

import "time"

type ModelHistory struct {
	ID              string `gorm:"size:256"`
	Name            string `gorm:"size:256"`
	Version         string `gorm:"size:256"`
	StartDate       time.Time
	EndDate         time.Time
	ApplyHistoryTag string `gorm:"size:256"` //"current", "previous"
	DeploymentID    string `gorm:"size:256;type:not null"`
	ModelPackageID  string `gorm:"size:256"`
}

func newModelHistory(id string, name string, version string, modelPackageID string) *ModelHistory {
	modelHistory := &ModelHistory{
		ID:              id,
		Name:            name,
		Version:         version,
		StartDate:       time.Now(),
		EndDate:         time.Time{},
		ApplyHistoryTag: "Current",
		ModelPackageID:  modelPackageID,
	}

	return modelHistory
}
