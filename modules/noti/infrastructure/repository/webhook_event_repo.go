package repository

import (
	"fmt"
	"math"
	"net/http"

	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/service"
	sysError "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/error"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NewAuthenticationRepo new AuthenticationRepo implement IAuthenticationRepo
func NewWebHookEventRepo(h *handler.Handler) (domRepo.IWebHookEventRepo, error) {
	repo := new(WebHookEventRepo)
	repo.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	repo.SetDBConnectionName(cfg.Databases.MainDB.ConnectionName)

	return repo, nil
}

// AuthenticationRepo type
type WebHookEventRepo struct {
	BaseRepo
}

func (r *WebHookEventRepo) Save(req *domEntity.WebHookEvent) error {
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

func (r *WebHookEventRepo) GetByID(ID string) (*domEntity.WebHookEvent, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var domEntity = &domEntity.WebHookEvent{}
	var count int64

	if err := dbCon.Where("id = ?", ID).Preload(clause.Associations).Find(&domEntity).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid webhook event id")}
	}

	// return response
	resp := domEntity

	return resp, nil

}

//Paging 및 Sorting을 위한 코드
func (r *WebHookEventRepo) paginate(value interface{}, pagination *Pagination, dbCon *gorm.DB, queryName string, filter string) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	var tmpLimit int

	dbCon.Model(value).Where("name like ? AND deployment_id = ?", "%"+queryName+"%", filter).Count(&totalRows)
	pagination.TotalRows = totalRows

	if pagination.Limit <= 0 {
		tmpLimit = 1
	} else {
		tmpLimit = pagination.Limit
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(tmpLimit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Where("name like ? AND deployment_id = ?", "%"+queryName+"%", filter)
	}
}

func (r *WebHookEventRepo) GetList(queryName string, pagination interface{}, filterEntity interface{}) ([]*domEntity.WebHookEvent, interface{}, error) {
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

	var result []*domEntity.WebHookEvent

	filterDomain := filterEntity.(domEntity.WebHookEvent)
	entityModel := domEntity.WebHookEvent{}

	if err := dbCon.Model(&entityModel).Scopes(r.paginate(&entityModel, &p, dbCon, queryName, filterDomain.DeploymentID)).Find(&result).Error; err != nil {
		return nil, p, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	resp := result

	return resp, p, nil
}

func (r *WebHookEventRepo) GetForUpdate(ID string) (*domEntity.WebHookEvent, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var entityModel = &domEntity.WebHookEvent{}
	var count int64

	if err := dbCon.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", ID).Preload(clause.Associations).Find(&entityModel).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid id")}
	}

	// return response
	resp := entityModel

	return resp, nil

}

func (r *WebHookEventRepo) Delete(id string) error {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return err
	}

	var domEntity = &domEntity.WebHookEvent{}
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
