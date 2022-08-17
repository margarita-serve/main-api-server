package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rs/xid"
)

type EventHistory struct {
	ID           string `gorm:"size:20"`
	EventType    string `gorm:"size:20"`
	EventDate    time.Time
	LogMessage   string
	UserID       string `gorm:"size:20"`
	DeploymentID string `gorm:"size:20;type:not null"`
}

// Validate
func (r *EventHistory) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.EventType, validation.Required, validation.In("Create", "Delete", "ReplaceModel", "Update", "Active", "InActive", "DataDriftAlert", "AccuracyAlert", "ServiceAlert")),
		validation.Field(&r.UserID, validation.Required),
	)
}

func newEventHistory(eventType string, logMessage string, userId string) (*EventHistory, error) {
	eventHistory := &EventHistory{
		ID:         xid.New().String(),
		EventType:  eventType,
		LogMessage: logMessage,
		EventDate:  time.Now(),
		UserID:     userId,
	}

	err := eventHistory.Validate()
	if err != nil {
		return nil, err
	}

	return eventHistory, err

}
