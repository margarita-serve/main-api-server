package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/domain/entity"

//interface
type IProjectRepo interface {
	Save(req *domEntity.Project) error
	GetByID(req string, identity interface{}) (*domEntity.Project, error)
	GetByIDInternal(req string) (*domEntity.Project, error)
	GetForUpdate(projectID string, identity interface{}) (*domEntity.Project, error)
	GetList(name string, pagination interface{}, identity interface{}) ([]*domEntity.Project, interface{}, error)
	Delete(req string) error
}
