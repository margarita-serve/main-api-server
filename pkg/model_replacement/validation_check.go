package model_replacement

type StatusType string

const (
	PASSING StatusType = "passing"
	FAILING StatusType = "failing"
	WARNING StatusType = "warning"
)

type ValidationCheck struct {
	TargetStatus           StatusType
	TargetMessage          string
	TargetTypeStatus       StatusType
	TargetTypeMessage      string
	TargetClassesStatus    StatusType
	TargetClassesMessage   string
	FeaturesStatus         StatusType
	FeaturesMessage        string
	NotCurrentModelStatus  StatusType
	NotCurrentModelMessage string
	PermissionStatus       StatusType
	PermissionMessage      string
}
