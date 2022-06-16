package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/entity"

// IEmailRepo interface
type IDeploymentRepo interface {
	Save(req *domEntity.Deployment) error
	Get(req string) (*domEntity.Deployment, error)
	ByName(req string) ([]domEntity.Deployment, error)
	Delete(req string) error
}
