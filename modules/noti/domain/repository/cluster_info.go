package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/domain/entity"

//interface
type IClusterInfoRepo interface {
	Save(req *domEntity.ClusterInfo) error
	GetByID(req string) (*domEntity.ClusterInfo, error)
	GetForUpdate(ClusterInfoID string) (*domEntity.ClusterInfo, error)
	GetList(name string, pagination interface{}, filter map[string]interface{}) ([]*domEntity.ClusterInfo, interface{}, error)
	Delete(req string) error
}
