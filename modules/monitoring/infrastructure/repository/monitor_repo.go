package repository

import (
	"fmt"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/repository"
	sysError "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/error"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"gorm.io/gorm/clause"
	"net/http"
)

func NewMonitorRepo(h *handler.Handler) (domRepo.IMonitorRepo, error) {
	repo := new(MonitorRepo)
	repo.handler = h
	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	repo.SetDBConnectionName(cfg.Databases.MainDB.ConnectionName)

	return repo, nil
}

type MonitorRepo struct {
	BaseRepo
}

func (r *MonitorRepo) Save(req *domEntity.Monitor) error {
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return err
	}

	//if err := dbCon.Create(&req).Error; err != nil {
	if err := dbCon.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&req).Error; err != nil {
		return &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	return nil
}

func (r *MonitorRepo) Get(ID string) (*domEntity.Monitor, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var domEntity = &domEntity.Monitor{}
	var count int64

	if err := dbCon.Where("id = ?", ID).Preload(clause.Associations).Find(&domEntity).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid monitor id")}
	}

	// return response
	resp := domEntity

	return resp, nil
}

func (r *MonitorRepo) ByName(name string) ([]domEntity.Monitor, error) {
	// select db
	print(name)
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var monitors []domEntity.Monitor
	domEntity := &monitors

	if err := dbCon.Where("name LIKE ?", "%"+name+"%").Find(domEntity).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// return response
	resp := *domEntity

	return resp, nil
}

func (r *MonitorRepo) GetAll(name string) ([]domEntity.Monitor, error) {
	// select db
	print(name)
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var monitors []domEntity.Monitor
	domEntity := &monitors

	if err := dbCon.Where(" = ?", "%"+name+"%").Find(domEntity).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	// return response
	resp := *domEntity

	return resp, nil
}

func (r *MonitorRepo) Delete(id string) error {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return err
	}

	var domEntity = &domEntity.Monitor{ID: id}
	var count int64

	if err := dbCon.Where("id = ?", domEntity.ID).Find(domEntity).Count(&count).Error; err != nil {
		return &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid monitor id")}
	}

	if err := dbCon.Select(clause.Associations).Delete(domEntity).Error; err != nil {
		return &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	return nil
}
