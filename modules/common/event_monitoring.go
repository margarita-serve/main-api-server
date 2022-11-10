package common

// Event Interface for Monitoring Domain Event
type MonitoringEvent interface {
	Event
	DeploymentID() string
}

// MonitoringCreated Event
type MonitoringCreated struct {
	deploymentID string
}

func (e MonitoringCreated) Name() string {
	return "event.monitoring.created"
}

func (e MonitoringCreated) DeploymentID() string {
	return e.deploymentID
}

func NewEventMonitoringCreated(deploymentID string) MonitoringCreated {
	return MonitoringCreated{deploymentID: deploymentID}
}

// MonitoringCreatFailed Event
type MonitoringCreateFailed struct {
	deploymentID string
	err          error
}

func (e MonitoringCreateFailed) Name() string {
	return "event.monitoring.createFailed"
}

func (e MonitoringCreateFailed) DeploymentID() string {
	return e.deploymentID
}

func (e MonitoringCreateFailed) Err() error {
	return e.err
}

func NewEventMonitoringCreateFailed(deploymentID string, err error) MonitoringCreateFailed {
	return MonitoringCreateFailed{deploymentID: deploymentID, err: err}
}

// MonitoringDataDriftStatusChangedToFailing Event
type MonitoringDataDriftStatusChangedToFailing struct {
	deploymentID string
}

func (e MonitoringDataDriftStatusChangedToFailing) Name() string {
	return "event.monitoring.dataDriftStatusChangedToFailing"
}

func (e MonitoringDataDriftStatusChangedToFailing) DeploymentID() string {
	return e.deploymentID
}

func NewEventMonitoringDataDriftStatusChangedToFailing(deploymentID string) MonitoringDataDriftStatusChangedToFailing {
	return MonitoringDataDriftStatusChangedToFailing{deploymentID: deploymentID}
}

// MonitoringDataDriftStatusChangedToAtrisk Event
type MonitoringDataDriftStatusChangedToAtrisk struct {
	deploymentID string
}

func (e MonitoringDataDriftStatusChangedToAtrisk) Name() string {
	return "event.monitoring.dataDriftStatusChangedToAtrisk"
}

func (e MonitoringDataDriftStatusChangedToAtrisk) DeploymentID() string {
	return e.deploymentID
}

func NewEventMonitoringDataDriftStatusChangedToAtrisk(deploymentID string) MonitoringDataDriftStatusChangedToAtrisk {
	return MonitoringDataDriftStatusChangedToAtrisk{deploymentID: deploymentID}
}

// MonitoringAccuracyStatusChangedToFailing Event
type MonitoringAccuracyStatusChangedToFailing struct {
	deploymentID string
}

func (e MonitoringAccuracyStatusChangedToFailing) Name() string {
	return "event.monitoring.accuracyStatusChangedToFailing"
}

func (e MonitoringAccuracyStatusChangedToFailing) DeploymentID() string {
	return e.deploymentID
}

func NewEventMonitoringAccuracyStatusChangedToFailing(deploymentID string) MonitoringAccuracyStatusChangedToFailing {
	return MonitoringAccuracyStatusChangedToFailing{deploymentID: deploymentID}
}

// MonitoringAccuracyStatusChangedToAtrisk Event
type MonitoringAccuracyStatusChangedToAtrisk struct {
	deploymentID string
}

func (e MonitoringAccuracyStatusChangedToAtrisk) Name() string {
	return "event.monitoring.accuracyStatusChangedToAtrisk"
}

func (e MonitoringAccuracyStatusChangedToAtrisk) DeploymentID() string {
	return e.deploymentID
}

func NewEventMonitoringAccuracyStatusChangedToAtrisk(deploymentID string) MonitoringAccuracyStatusChangedToAtrisk {
	return MonitoringAccuracyStatusChangedToAtrisk{deploymentID: deploymentID}
}

// MonitoringServiceHealthStatusChangedToFailing Event
type MonitoringServiceHealthStatusChangedToFailing struct {
	deploymentID string
}

func (e MonitoringServiceHealthStatusChangedToFailing) Name() string {
	return "event.monitoring.serviceHealthStatusChangedToFailing"
}

func (e MonitoringServiceHealthStatusChangedToFailing) DeploymentID() string {
	return e.deploymentID
}

func NewEventMonitoringServiceHealthStatusChangedToFailing(deploymentID string) MonitoringServiceHealthStatusChangedToFailing {
	return MonitoringServiceHealthStatusChangedToFailing{deploymentID: deploymentID}
}

// MonitoringServiceHealthStatusChangedToAtrisk Event
type MonitoringServiceHealthStatusChangedToAtrisk struct {
	deploymentID string
}

func (e MonitoringServiceHealthStatusChangedToAtrisk) Name() string {
	return "event.monitoring.serviceHealthStatusChangedToAtrisk"
}

func (e MonitoringServiceHealthStatusChangedToAtrisk) DeploymentID() string {
	return e.deploymentID
}

func NewEventMonitoringServiceHealthStatusChangedToAtrisk(deploymentID string) MonitoringServiceHealthStatusChangedToAtrisk {
	return MonitoringServiceHealthStatusChangedToAtrisk{deploymentID: deploymentID}
}

type MonitoringDataDriftMonitorEnabled struct {
	deploymentID string
	userID       string
}

func (e MonitoringDataDriftMonitorEnabled) Name() string {
	return "event.monitoring.dataDriftMonitorEnabled"
}

func (e MonitoringDataDriftMonitorEnabled) DeploymentID() string {
	return e.deploymentID
}

func (e MonitoringDataDriftMonitorEnabled) UserID() string {
	return e.userID
}

func NewEventMonitoringDataDriftMonitorEnabled(deploymentID string, userID string) MonitoringDataDriftMonitorEnabled {
	return MonitoringDataDriftMonitorEnabled{deploymentID: deploymentID, userID: userID}
}

type MonitoringDataDriftMonitorDisabled struct {
	deploymentID string
	userID       string
}

func (e MonitoringDataDriftMonitorDisabled) Name() string {
	return "event.monitoring.dataDriftMonitorDisabled"
}

func (e MonitoringDataDriftMonitorDisabled) DeploymentID() string {
	return e.deploymentID
}

func (e MonitoringDataDriftMonitorDisabled) UserID() string {
	return e.userID
}

func NewEventMonitoringDataDriftMonitorDisabled(deploymentID string, userID string) MonitoringDataDriftMonitorDisabled {
	return MonitoringDataDriftMonitorDisabled{deploymentID: deploymentID, userID: userID}
}

type MonitoringAccuracyMonitorEnabled struct {
	deploymentID string
	userID       string
}

func (e MonitoringAccuracyMonitorEnabled) Name() string {
	return "event.monitoring.accuracyMonitorEnabled"
}

func (e MonitoringAccuracyMonitorEnabled) DeploymentID() string {
	return e.deploymentID
}

func (e MonitoringAccuracyMonitorEnabled) UserID() string {
	return e.userID
}

func NewEventMonitoringAccuracyMonitorEnabled(deploymentID string, userID string) MonitoringAccuracyMonitorEnabled {
	return MonitoringAccuracyMonitorEnabled{deploymentID: deploymentID, userID: userID}
}

type MonitoringAccuracyMonitorDisabled struct {
	deploymentID string
	userID       string
}

func (e MonitoringAccuracyMonitorDisabled) Name() string {
	return "event.monitoring.accuracyMonitorDisabled"
}

func (e MonitoringAccuracyMonitorDisabled) DeploymentID() string {
	return e.deploymentID
}

func (e MonitoringAccuracyMonitorDisabled) UserID() string {
	return e.userID
}

func NewEventMonitoringAccuracyMonitorDisabled(deploymentID string, userID string) MonitoringAccuracyMonitorDisabled {
	return MonitoringAccuracyMonitorDisabled{deploymentID: deploymentID, userID: userID}
}
