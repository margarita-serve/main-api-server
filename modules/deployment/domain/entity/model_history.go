package entity

import "time"

type ModelHistory struct {
	ID              string `gorm:"size:256"`
	Name            string `gorm:"size:256"`
	StartDate       time.Time
	EndDate         time.Time
	ApplyHistoryTag string `gorm:"size:256"` //"current", "previous"
	DeploymentID    string `gorm:"size:256;type:not null"`
}
