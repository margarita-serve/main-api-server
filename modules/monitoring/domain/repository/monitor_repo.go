package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/entity"

type IMonitorRepo interface {
	Save(req *domEntity.Monitor) error
	Get(req string) (*domEntity.Monitor, error)
	ByName(req string) ([]domEntity.Monitor, error)
	Delete(req string) error
}
