package feature

import (
	appDeploymentSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/service"
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	appProjectSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/service"
	appResourceSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewFeature new Feature
func NewFeature(h *handler.Handler) (*Feature, error) {
	var err error

	f := new(Feature)
	f.handler = h

	if f.System, err = NewSystem(h); err != nil {
		return nil, err
	}

	if f.OpenAPI, err = NewOpenAPI(h); err != nil {
		return nil, err
	}

	ClusterInfoSvc, err := appResourceSvc.NewClusterInfoService(h)
	if err != nil {
		return nil, err
	}

	PredictionEnvSvc, err := appResourceSvc.NewPredictionEnvService(h, ClusterInfoSvc)
	if err != nil {
		return nil, err
	}

	if f.Resource, err = NewResource(h, ClusterInfoSvc, PredictionEnvSvc); err != nil {
		return nil, err
	}

	ProjectSvc, err := appProjectSvc.NewProjectService(h)
	if err != nil {
		return nil, err
	}

	if f.Project, err = NewProject(h, ProjectSvc); err != nil {
		return nil, err
	}

	ModelPackageSvc, err := appModelPackageSvc.NewModelPackageService(h, ProjectSvc)
	if err != nil {
		return nil, err
	}
	if f.ModelPackage, err = NewModelPackage(h, ModelPackageSvc); err != nil {
		return nil, err
	}

	DeploymentGetByIDInternalSvc, err := appDeploymentSvc.NewDeploymentGetByIDInternalService(h)
	if err != nil {
		return nil, err
	}
	DeploymentGovernanceHistorySvc, err := appDeploymentSvc.NewDeploymentGovernanceHistoryService(h)
	if err != nil {
		return nil, err
	}

	if f.Auths, err = NewFAuths(h); err != nil {
		return nil, err
	}

	if f.Email, err = NewFEmail(h); err != nil {
		return nil, err
	}

	if f.Noti, err = NewFNoti(h, f.Email.appEmail.EmailSvc, DeploymentGetByIDInternalSvc, ProjectSvc, f.Auths.appAuths.AuthenticationSvc, DeploymentGovernanceHistorySvc); err != nil {
		return nil, err
	}

	if f.Monitor, err = NewMonitor(h, ModelPackageSvc); err != nil {
		return nil, err
	}

	if f.Deployment, err = NewDeployment(h, PredictionEnvSvc, ProjectSvc, ModelPackageSvc, f.Monitor.appMonitor.MonitorSvc); err != nil {
		return nil, err
	}

	return f, nil
}

// Feature represet Feature
type Feature struct {
	BaseFeature

	System       *FSystem
	OpenAPI      *FOpenAPI
	Deployment   *FDeployment
	ModelPackage *FModelPackage
	Monitor      *FMonitor
	Auths        *FAuths
	Email        *FEmail
	Project      *FProject
	Resource     *FResource
	Noti         *FNoti
}
