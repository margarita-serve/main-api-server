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
