package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/entity"

// IEmailRepo interface
type IDeploymentRepo interface {
	Save(req *domEntity.Deployment) error
	GetByID(deploymentID string) (*domEntity.Deployment, error)
	GetGovernanceHistory(deploymentID string) (*domEntity.Deployment, error)
	GetForUpdate(deploymentID string) (*domEntity.Deployment, error)
	GetList(name string, pagination interface{}) ([]*domEntity.Deployment, interface{}, error)
	Delete(deploymentID string) error
}
