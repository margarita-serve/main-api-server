package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/domain/entity"

//interface
type IPredictionEnvRepo interface {
	Save(req *domEntity.PredictionEnv) error
	GetByID(req string) (*domEntity.PredictionEnv, error)
	GetForUpdate(predictionEnvID string) (*domEntity.PredictionEnv, error)
	GetList(name string, pagination interface{}, filter map[string]interface{}) ([]*domEntity.PredictionEnv, interface{}, error)
	GetListByProject(name string, pagination interface{}, filter map[string]interface{}) ([]*domEntity.PredictionEnv, interface{}, error)
	Delete(req string) error
}
