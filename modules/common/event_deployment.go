// for describing Deployment relevant Domain Event
package common

// Event Interface for Deployment Domain Event
type DeploymentEvent interface {
	Event
	DeploymentID() string
}

// DeploymentCreated Event
type DeploymentCreated struct {
	deploymentID           string
	modelPackageID         string
	featureDriftTracking   bool
	accuracyMonitoring     bool
	associationID          string
	associationIDInFeature bool
	modelHistoryID         string
}

func (e DeploymentCreated) Name() string {
	return "event.deployment.created"
}

func (e DeploymentCreated) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentCreated) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentCreated) FeatureDriftTracking() bool {
	return e.featureDriftTracking
}

func (e DeploymentCreated) AccuracyMonitoring() bool {
	return e.accuracyMonitoring
}

func (e DeploymentCreated) AssociationID() string {
	return e.modelPackageID
}

func (e DeploymentCreated) AssociationIDInFeature() bool {
	return e.associationIDInFeature
}

func (e DeploymentCreated) ModelHistoryID() string {
	return e.modelHistoryID
}

func NewEventDeploymentCreated(deploymentID string, modelPackageID string, featureDriftTracking bool, accuracyMonitoring bool, associationID string, associationIDInFeature bool, modelHistoryID string) DeploymentCreated {
	return DeploymentCreated{deploymentID: deploymentID,
		modelPackageID:         modelPackageID,
		featureDriftTracking:   featureDriftTracking,
		accuracyMonitoring:     accuracyMonitoring,
		associationID:          associationID,
		associationIDInFeature: associationIDInFeature,
		modelHistoryID:         modelHistoryID,
	}
}

type DeploymentInferenceServiceCreated struct {
	deploymentID           string
	modelPackageID         string
	featureDriftTracking   bool
	accuracyMonitoring     bool
	associationID          string
	associationIDInFeature bool
	modelHistoryID         string
}

func (e DeploymentInferenceServiceCreated) Name() string {
	return "event.deployment.inferenceServiceCreated"
}

func (e DeploymentInferenceServiceCreated) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentInferenceServiceCreated) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentInferenceServiceCreated) FeatureDriftTracking() bool {
	return e.featureDriftTracking
}

func (e DeploymentInferenceServiceCreated) AccuracyMonitoring() bool {
	return e.accuracyMonitoring
}

func (e DeploymentInferenceServiceCreated) AssociationID() string {
	return e.associationID
}

func (e DeploymentInferenceServiceCreated) AssociationIDInFeature() bool {
	return e.associationIDInFeature
}

func (e DeploymentInferenceServiceCreated) ModelHistoryID() string {
	return e.modelHistoryID
}

func NewEventDeploymentInferenceServiceCreated(deploymentID string, modelPackageID string, featureDriftTracking bool, accuracyMonitoring bool, associationID string, associationIDInFeature bool, modelHistoryID string) DeploymentInferenceServiceCreated {
	return DeploymentInferenceServiceCreated{deploymentID: deploymentID,
		modelPackageID:         modelPackageID,
		featureDriftTracking:   featureDriftTracking,
		accuracyMonitoring:     accuracyMonitoring,
		associationID:          associationID,
		associationIDInFeature: associationIDInFeature,
		modelHistoryID:         modelHistoryID,
	}
}

// DeploymentModelReplaced Event
type DeploymentModelReplaced struct {
	deploymentID   string
	modelPackageID string
	modelHistoryID string
}

func (e DeploymentModelReplaced) Name() string {
	return "event.deployment.modelReplaced"
}

func (e DeploymentModelReplaced) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentModelReplaced) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentModelReplaced) ModelHistoryID() string {
	return e.modelHistoryID
}

func NewEventDeploymentModelReplaced(deploymentID string, modelPackageID string, modelHistoryID string) DeploymentModelReplaced {
	return DeploymentModelReplaced{deploymentID: deploymentID, modelPackageID: modelPackageID, modelHistoryID: modelHistoryID}
}

// DeploymentUpdated Event
type DeploymentAssociationIDUpdated struct {
	deploymentID   string
	modelPackageID string
	associationID  string
	currentModelID string
}

func (e DeploymentAssociationIDUpdated) Name() string {
	return "event.deployment.deploymentAssociationIDUpdated"
}

func (e DeploymentAssociationIDUpdated) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentAssociationIDUpdated) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentAssociationIDUpdated) AssociationID() string {
	return e.associationID
}

func (e DeploymentAssociationIDUpdated) CurrentModelID() string {
	return e.currentModelID
}

func NewEventDeploymentAssociationIDUpdated(deploymentID string, modelPackageID string, currentModelID string, associationID string) DeploymentAssociationIDUpdated {
	return DeploymentAssociationIDUpdated{deploymentID: deploymentID, modelPackageID: modelPackageID, currentModelID: currentModelID, associationID: associationID}
}

// DeploymentFeatureDriftTrackingEnabled Event
type DeploymentFeatureDriftTrackingEnabled struct {
	deploymentID   string
	modelPackageID string
	currentModelID string
}

func (e DeploymentFeatureDriftTrackingEnabled) Name() string {
	return "event.deployment.featureDriftTrackingEnabled"
}

func (e DeploymentFeatureDriftTrackingEnabled) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentFeatureDriftTrackingEnabled) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentFeatureDriftTrackingEnabled) CurrentModelID() string {
	return e.currentModelID
}

func NewEventDeploymentFeatureDriftTrackingEnabled(deploymentID string, modelPackageID string, currentModelID string) DeploymentFeatureDriftTrackingEnabled {
	return DeploymentFeatureDriftTrackingEnabled{deploymentID: deploymentID, modelPackageID: modelPackageID, currentModelID: currentModelID}
}

// DeploymentFeatureDriftTrackingDisabled Event
type DeploymentFeatureDriftTrackingDisabled struct {
	deploymentID   string
	modelPackageID string
	currentModelID string
}

func (e DeploymentFeatureDriftTrackingDisabled) Name() string {
	return "event.deployment.featureDriftTrackingDisabled"
}

func (e DeploymentFeatureDriftTrackingDisabled) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentFeatureDriftTrackingDisabled) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentFeatureDriftTrackingDisabled) CurrentModelID() string {
	return e.currentModelID
}

func NewEventDeploymentFeatureDriftTrackingDisabled(deploymentID string, modelPackageID string, currentModelID string) DeploymentFeatureDriftTrackingDisabled {
	return DeploymentFeatureDriftTrackingDisabled{deploymentID: deploymentID, modelPackageID: modelPackageID, currentModelID: currentModelID}
}

// DeploymentAccuracyAnalyzeEnabled Event
type DeploymentAccuracyAnalyzeEnabled struct {
	deploymentID   string
	modelPackageID string
	currentModelID string
	associationID  string
}

func (e DeploymentAccuracyAnalyzeEnabled) Name() string {
	return "event.deployment.accuracyAnalyzeEnabled"
}

func (e DeploymentAccuracyAnalyzeEnabled) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentAccuracyAnalyzeEnabled) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentAccuracyAnalyzeEnabled) CurrentModelID() string {
	return e.currentModelID
}

func (e DeploymentAccuracyAnalyzeEnabled) AssociationID() string {
	return e.associationID
}

func NewEventDeploymentAccuracyAnalyzeEnabled(deploymentID string, modelPackageID string, currentModelID string, associationID string) DeploymentAccuracyAnalyzeEnabled {
	return DeploymentAccuracyAnalyzeEnabled{deploymentID: deploymentID, modelPackageID: modelPackageID, currentModelID: currentModelID, associationID: associationID}
}

// DeploymentAccuracyAnalyzeDisabled Event
type DeploymentAccuracyAnalyzeDisabled struct {
	deploymentID   string
	modelPackageID string
	currentModelID string
}

func (e DeploymentAccuracyAnalyzeDisabled) Name() string {
	return "event.deployment.accuracyAnalyzeDisabled"
}

func (e DeploymentAccuracyAnalyzeDisabled) DeploymentID() string {
	return e.deploymentID
}

func (e DeploymentAccuracyAnalyzeDisabled) ModelPackageID() string {
	return e.modelPackageID
}

func (e DeploymentAccuracyAnalyzeDisabled) CurrentModelID() string {
	return e.currentModelID
}

func NewEventDeploymentAccuracyAnalyzeDisabled(deploymentID string, modelPackageID string, currentModelID string) DeploymentAccuracyAnalyzeDisabled {
	return DeploymentAccuracyAnalyzeDisabled{deploymentID: deploymentID, modelPackageID: modelPackageID, currentModelID: currentModelID}
}

// DeploymentDeleted Event
type DeploymentDeleted struct {
	deploymentID string
}

func (e DeploymentDeleted) Name() string {
	return "event.deployment.deleted"
}

func (e DeploymentDeleted) DeploymentID() string {
	return e.deploymentID
}

func NewEventDeploymentDeleted(deploymentID string) DeploymentDeleted {
	return DeploymentDeleted{deploymentID: deploymentID}
}

// DeploymentActived Event
type DeploymentActived struct {
	deploymentID string
}

func (e DeploymentActived) Name() string {
	return "event.deployment.actived"
}

func (e DeploymentActived) DeploymentID() string {
	return e.deploymentID
}

func NewEventDeploymentActived(deploymentID string) DeploymentActived {
	return DeploymentActived{deploymentID: deploymentID}
}

// DeploymentInActived Event
type DeploymentInActived struct {
	deploymentID string
}

func (e DeploymentInActived) Name() string {
	return "event.deployment.inActived"
}

func (e DeploymentInActived) DeploymentID() string {
	return e.deploymentID
}

func NewEventDeploymentInActived(deploymentID string) DeploymentInActived {
	return DeploymentInActived{deploymentID: deploymentID}
}
