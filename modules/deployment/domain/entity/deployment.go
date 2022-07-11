package entity

import (
	"errors"
	"fmt"
	"time"

	domSvcInferenceSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service"
	domSvcInferenceDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rs/xid"
)

type Deployment struct {
	ID              string  `gorm:"size:256"`
	ProjectID       string  `gorm:"size:256"`
	ModelPackageID  string  `gorm:"size:256"`
	PredictionEnvID string  `gorm:"size:256"`
	Name            string  `gorm:"size:256"`
	Description     string  `gorm:"size:256"`
	Importance      string  `gorm:"size:256"`
	DeployType      string  `gorm:"size:256"`
	RequestCPU      float32 `gorm:"size:8"`
	RequestMEM      float32 `gorm:"size:8"`
	LimitCPU        float32 `gorm:"size:8"`
	LimitMEM        float32 `gorm:"size:8"`
	ActiveStatus    string  `gorm:"size:256"`
	ServiceStatus   string  `gorm:"size:256"`
	ChangeRequested bool
	ModelHistory    []*ModelHistory `gorm:"foreignKey:DeploymentID;references:ID"`
	EventHistory    []*EventHistory `gorm:"foreignKey:DeploymentID;references:ID"`
	BaseEntity
}

// Validate
func (r *Deployment) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.ProjectID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.ModelPackageID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.Name, validation.Length(0, 255)),
		validation.Field(&r.Description, validation.Length(0, 255)),
		validation.Field(&r.PredictionEnvID, validation.Required, validation.NotNil, validation.Length(20, 20)),
		validation.Field(&r.Importance, validation.In("Low", "Moderate", "High", "Critical")),
		validation.Field(&r.RequestCPU, validation.Min(0.1), validation.Max(2.0)),
		validation.Field(&r.RequestMEM, validation.Min(0.1), validation.Max(2.0)),
		validation.Field(&r.ID, validation.Required, validation.NotNil, validation.Length(20, 20)),
	)
}

func NewDeployment(id string, projectID string, modelPackageID string, predictionEnvID string,
	name string, description string, importance string, deployType string, requestCPU float32, requestMEM float32, limitCPU float32, limitMEM float32, ownerID string) (*Deployment, error) {

	var modelHistory []*ModelHistory
	var eventHistory []*EventHistory

	var baseEntity BaseEntity
	baseEntity.CreatedBy = "testuser"

	// Validate
	err := validation.Validate(deployType, validation.Required, validation.In("Normal", "Test", "Challenger"))
	if err != nil {
		return nil, err
	}

	// Default CPU, MEM Resource Setting
	if requestCPU == 0 {
		requestCPU = 1
	}
	if requestMEM == 0 {
		requestMEM = 2
	}

	if limitCPU == 0 {
		limitCPU = requestCPU
	}
	if limitMEM == 0 {
		limitMEM = requestMEM
	}

	deployment := &Deployment{
		id,
		projectID,
		modelPackageID,
		predictionEnvID,
		name,
		description,
		importance,
		deployType,
		requestCPU,
		requestMEM,
		limitCPU,
		limitMEM,
		"",
		"",
		false,
		modelHistory,
		eventHistory,
		baseEntity}

	deployment.SetServiceStatusCreateRequested()
	deployment.SetActiveStatusInActive()

	return deployment, nil
}

func (d *Deployment) AddEventHistory(eventType string, logMessage string, userId string) error {
	eventHistory, err := newEventHistory(eventType, logMessage, userId)
	if err != nil {
		return err
	}
	d.EventHistory = append(d.EventHistory, eventHistory)
	return nil
}

func (d *Deployment) AddModelHistory(name string, version string) {
	for i, mh := range d.ModelHistory {
		if mh.ApplyHistoryTag == "Current" {
			d.ModelHistory[i].ApplyHistoryTag = "Previous"
			d.ModelHistory[i].EndDate = time.Now()
		}
	}

	guid := xid.New().String()

	modelHistory := newModelHistory(guid, name, version)
	d.ModelHistory = append(d.ModelHistory, modelHistory)

}

func (d *Deployment) SetDeploymentActive(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceActiveRequest) error {
	if d.ActiveStatus == "Active" {
		return errors.New("already active")
	}

	res, err := domSvc.InferenceServiceActive(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		d.SetServiceStatusError()
		return err
	}

	d.SetActiveStatusActive()
	return err
}

func (d *Deployment) SetDeploymentInActive(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceInActiveRequest) error {
	if d.ActiveStatus == "InActive" {
		return errors.New("already inactive")
	}

	res, err := domSvc.InferenceServiceInActive(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		d.SetServiceStatusError()
		return err
	}

	d.SetActiveStatusInActive()
	return err
}

func (d *Deployment) RequestCreateInferenceService(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceCreateRequest) error {
	res, err := domSvc.InferenceServiceCreate(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		d.SetServiceStatusError()
		return err
	}

	d.SetActiveStatusActive()
	d.SetServiceStatusReady()

	return err
}

func (d *Deployment) RequestReplaceModelInferenceService(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceReplaceModelRequest) error {
	res, err := domSvc.InferenceServiceReplaceModel(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		d.SetServiceStatusError()
		return err
	}

	d.SetServiceStatusReady()

	return err
}

func (d *Deployment) RequestDeleteInferenceService(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceDeleteRequest) error {
	res, err := domSvc.InferenceServiceDelete(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		return err
	}

	return err
}

func (d *Deployment) UpdateDeploymentName(req string) {
	d.Name = req
}

func (d *Deployment) UpdateDeploymentDescription(req string) {
	d.Description = req
}

func (d *Deployment) UpdateDeploymentImportance(req string) {
	d.Importance = req
}

func (d *Deployment) ChangeModelPackage(req string) {
	d.ModelPackageID = req
}

func (d *Deployment) SetServiceStatusError() {
	d.ServiceStatus = "Error"
}

func (d *Deployment) SetServiceStatusReady() {
	d.ServiceStatus = "Ready"
}

func (d *Deployment) SetServiceStatusReplacingModel() {
	d.ServiceStatus = "ReplacingModel"
}

func (d *Deployment) SetServiceStatusCreateRequested() {
	d.ServiceStatus = "CreateRequested"
}

func (d *Deployment) SetActiveStatusActive() {
	d.ActiveStatus = "Active"
}

func (d *Deployment) SetActiveStatusInActive() {
	d.ActiveStatus = "InActive"
}
