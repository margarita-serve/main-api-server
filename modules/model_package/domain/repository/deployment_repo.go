package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/entity"

//interface
type IModelPackageRepo interface {
	Save(req *domEntity.ModelPackage) error
	Get(req string) (*domEntity.ModelPackage, error)
	ByName(req string) ([]domEntity.ModelPackage, error)
	Delete(req string) error
}
