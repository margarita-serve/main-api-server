package repository

import (
	"fmt"
	"net/http"

	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/service"
	sysError "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/error"
	"gorm.io/gorm/clause"
)

// NewAuthenticationRepo new AuthenticationRepo implement IAuthenticationRepo
func NewWebHookRepo(h *handler.Handler) (domRepo.IWebHookRepo, error) {
	repo := new(WebHookRepo)
	repo.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	repo.SetDBConnectionName(cfg.Databases.MainDB.ConnectionName)

	return repo, nil
}

// AuthenticationRepo type
type WebHookRepo struct {
	BaseRepo
}

func (r *WebHookRepo) Save(req *domEntity.WebHook) error {
	// select db
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

func (r *WebHookRepo) GetByID(id string) (*domEntity.WebHook, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var domEntity = &domEntity.WebHook{}
	var count int64

	if err := dbCon.Where("id = ?", id).Preload(clause.Associations).Find(&domEntity).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid WebHook id")}
	}

	// return response
	resp := domEntity

	return resp, nil

}

func (r *WebHookRepo) GetList(queryName string, pagination interface{}, filterEntity interface{}) ([]*domEntity.WebHook, interface{}, error) {
	p := pagination.(Pagination)

	var mapSort string
	switch p.Sort {
	case "CreateDesc":
		mapSort = "id desc"
	case "CreateASC":
		mapSort = "id asc"
	}

	p.Sort = mapSort

	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, p, err
	}

	var result []*domEntity.WebHook

	filterDomain := filterEntity.(domEntity.WebHook)
	entityModel := domEntity.WebHook{
		DeploymentID: filterDomain.DeploymentID,
	}

	if err := dbCon.Model(&entityModel).Scopes(paginate(&entityModel, &p, dbCon, queryName)).Find(&result).Error; err != nil {
		return nil, p, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	resp := result

	return resp, p, nil
}

func (r *WebHookRepo) GetForUpdate(id string) (*domEntity.WebHook, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var entityModel = &domEntity.WebHook{ID: id}
	var count int64

	if err := dbCon.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).Preload(clause.Associations).Find(&entityModel).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid id")}
	}

	// return response
	resp := entityModel

	return resp, nil

}

func (r *WebHookRepo) Delete(id string) error {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return err
	}

	var domEntity = &domEntity.WebHook{}
	var count int64

	if err := dbCon.Where("id = ?", id).Find(&domEntity).Count(&count).Error; err != nil {
		return &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid id")}
	}

	if err := dbCon.Select(clause.Associations).Delete(domEntity).Error; err != nil {
		return &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	return nil
}

func (r *WebHookRepo) GetListByInternal(filterEntity interface{}) ([]*domEntity.WebHook, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var result []*domEntity.WebHook

	filterDomain := filterEntity.(domEntity.WebHook)
	entityModel := domEntity.WebHook{
		DeploymentID:  filterDomain.DeploymentID,
		TriggerSource: filterDomain.TriggerSource,
	}

	if err := dbCon.Where(entityModel).Find(&result).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	resp := result

	return resp, nil
}
