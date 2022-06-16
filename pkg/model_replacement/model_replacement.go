package model_replacement

type ReplacementReason string

const (
	ACCURANCY         ReplacementReason = "accurancy"
	DATA_DRIFT        ReplacementReason = "data_drift"
	ERRORS            ReplacementReason = "errors"
	SCHEDULED_REFRESH ReplacementReason = "scheduled_refresh"
	SCORING_SPEED     ReplacementReason = "scoring_speed"
	OTHER             ReplacementReason = "other"
)

type ModelReplacement struct {
	ID                  string `gorm:"primarykey"`
	TargetDeploymentID  string
	ApplyModelPackageID string
	ManualApply         bool
	ReplacementReason   ReplacementReason
	ValidationCheck     ValidationCheck `gorm:"embedded;embeddedPrefix:validaion_check_"`
}
