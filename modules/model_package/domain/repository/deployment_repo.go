package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/entity"

//interface
type IModelPackageRepo interface {
	Save(req *domEntity.ModelPackage) error
	GetByID(req string) (*domEntity.ModelPackage, error)
	GetForUpdate(modelPackageID string) (*domEntity.ModelPackage, error)
	GetList(name string, pagination interface{}) ([]*domEntity.ModelPackage, interface{}, error)
	Delete(req string) error
}
