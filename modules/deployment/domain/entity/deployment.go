package entity

import (
	"errors"
	"fmt"
	"time"

	domSvcInferenceSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service"
	domSvcInferenceDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/service/inference_service/dto"
)

var ErrDeploymentID = errors.New("empty deployment id")

type Deployment struct {
	ID              string `gorm:"size:256"`
	ProjectID       string `gorm:"size:256"`
	ModelPackageID  string `gorm:"size:256"`
	PredictionEnvID string `gorm:"size:256"`
	Name            string `gorm:"size:256"`
	Description     string `gorm:"size:256"`
	Importance      string `gorm:"size:256"`
	DeployType      string `gorm:"size:256"`
	RequestCPU      string `gorm:"size:8"`
	RequestMEM      string `gorm:"size:8"`
	RequestGPU      string `gorm:"size:8"`
	ActiveStatus    string `gorm:"size:256"`
	ServiceStatus   string `gorm:"size:256"`
	ChangeRequested bool
	ModelHistory    []ModelHistory `gorm:"foreignKey:DeploymentID;references:ID"`
	BaseEntity
}

func NewDeployment(id string, projectID string, modelPackageID string, predictionEnvID string,
	name string, description string, importance string, requestCPU string, requestMEM string, requestGPU string, ownerID string) (*Deployment, error) {

	var mh []ModelHistory

	var be BaseEntity
	be.CreatedBy = "testuser"

	deployment := &Deployment{
		id,
		projectID,
		modelPackageID,
		predictionEnvID,
		name,
		description,
		importance,
		"normal",
		requestCPU,
		requestMEM,
		requestGPU,
		"inactive",
		"createRequested",
		false,
		mh,
		be}

	return deployment, nil
}

func (d *Deployment) AddModelHistory(id string, name string, startDate time.Time, endDate time.Time, applyHistoryTag string) {
	var mh ModelHistory
	mh.ID = id
	mh.Name = name
	mh.StartDate = startDate
	mh.EndDate = endDate
	mh.ApplyHistoryTag = applyHistoryTag

	d.ModelHistory = append(d.ModelHistory, mh)
}

func (d *Deployment) SetDeploymentActive(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceCreateRequest) error {
	err := domSvc.InferenceServiceActive(d.ID)

	if err != nil {
		d.ServiceStatus = "error"
		return err
	}

	d.ActiveStatus = "active"
	return err
}

func (d *Deployment) SetDeploymentInactive(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceCreateRequest) error {
	err := domSvc.InferenceServiceInActive(d.ID)

	if err != nil {
		d.ServiceStatus = "error"
		return err
	}

	d.ActiveStatus = "inactive"
	return err
}

func (d *Deployment) RequestCreateInferenceService(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceCreateRequest) error {
	d.ServiceStatus = "createRequested"

	res, err := domSvc.InferenceServiceCreate(&reqDom)
	fmt.Printf("res: %v\n", res)

	if err != nil {
		d.ServiceStatus = "error"
		return err
	}

	d.AddModelHistory("000001", reqDom.ModelName, time.Now(), time.Time{}, "current")

	d.ServiceStatus = "ready"

	return err
}

func (d *Deployment) RequestDeleteInferenceService(domSvc domSvcInferenceSvc.IInferenceServiceAdapter, reqDom domSvcInferenceDTO.InferenceServiceDeleteRequest) error {

	err := domSvc.InferenceServiceDelete(&reqDom)

	if err != nil {
		return err
	}

	return err
}

// func (d *Deployment) InferenceServiceModelReplaceRequest(domSvc domSvcInference.IInferenceServiceAdapter, reqDom domSvcInference.InferenceServiceModelReplaceRequest, modelPackageID string, now time.Time) (*domSvcInference.InferenceServiceModelReplaceResponse, error) {

// 	res, err := domSvc.Update(&reqDom)

// 	if err != nil {
// 		d.ServiceStatus = "error"
// 		return res, err
// 	}

// 	d.ModelPackageID = modelPackageID
// 	d.ModelHistory = append(d.ModelHistory, &ModelHistory{ID: res.ModelHistoryID, Name: reqDom.ModelName, StartDate: now, EndDate: time.Time{}, ApplyHistoryTag: "current"})
// 	d.ServiceStatus = "ready"

// 	return res, err
// }

func (d *Deployment) SetServiceStatusReplacingModel() {
	d.ServiceStatus = "replacingModel"
}

func (d *Deployment) SetImportanceLow() {
	d.Importance = "low"
}

func (d *Deployment) SetImportanceModerate() {
	d.Importance = "moderate"
}

func (d *Deployment) SetImportanceHigh() {
	d.Importance = "high"
}

func (d *Deployment) SetImportanceCritical() {
	d.Importance = "critical"
}

func (d *Deployment) SetDeployTypeNormal() {
	d.Importance = "normal"
}

func (d *Deployment) SetDeployTypeTest() {
	d.Importance = "test"
}

func (d *Deployment) SetDeployTypeChallenger() {
	d.Importance = "challenger"
}
