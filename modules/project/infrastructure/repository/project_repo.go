package repository

import (
	"fmt"
	"math"
	"net/http"

	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/service"
	sysError "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/error"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NewAuthenticationRepo new AuthenticationRepo implement IAuthenticationRepo
func NewProjectRepo(h *handler.Handler) (domRepo.IProjectRepo, error) {
	repo := new(ProjectRepo)
	repo.handler = h

	cfg, err := h.GetConfig()
	if err != nil {
		return nil, err
	}
	repo.SetDBConnectionName(cfg.Databases.MainDB.ConnectionName)

	return repo, nil
}

// AuthenticationRepo type
type ProjectRepo struct {
	BaseRepo
}

func (r *ProjectRepo) Save(req *domEntity.Project) error {
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

func (r *ProjectRepo) GetByID(ID string, i interface{}) (*domEntity.Project, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	identityInfo := i.(identity.Identity)

	var domEntity = &domEntity.Project{}
	var count int64

	if err := dbCon.Where("id = ? AND sys_created_by = ?", ID, identityInfo.Claims.Username).Preload(clause.Associations).Find(&domEntity).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid project id")}
	}

	// return response
	resp := domEntity

	return resp, nil

}

func (r *ProjectRepo) GetByIDInternal(ID string) (*domEntity.Project, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	var domEntity = &domEntity.Project{}
	var count int64

	if err := dbCon.Where("id = ?", ID).Preload(clause.Associations).Find(&domEntity).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid project id")}
	}

	// return response
	resp := domEntity

	return resp, nil

}

//Paging 및 Sorting을 위한 코드
func paginate(value interface{}, pagination *Pagination, dbCon *gorm.DB, queryName string, i *identity.Identity) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	var tmpLimit int

	dbCon.Model(value).Where("name like ? AND sys_created_by = ?", "%"+queryName+"%", i.Claims.Username).Count(&totalRows)
	pagination.TotalRows = totalRows

	if pagination.Limit <= 0 {
		tmpLimit = 1
	} else {
		tmpLimit = pagination.Limit
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(tmpLimit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Where("(name like ?) AND (sys_created_by = ?)", "%"+queryName+"%", i.Claims.Username)
	}
}

func (r *ProjectRepo) GetList(queryName string, pagination interface{}, i interface{}) ([]*domEntity.Project, interface{}, error) {
	p := pagination.(Pagination)
	identityInfo := i.(identity.Identity)

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

	var entityModel []*domEntity.Project

	if err := dbCon.Model(entityModel).Scopes(paginate(&entityModel, &p, dbCon, queryName, &identityInfo)).Find(&entityModel).Error; err != nil {
		return nil, p, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	resp := entityModel

	return resp, p, nil
}

func (r *ProjectRepo) GetForUpdate(modelPackageID string, i interface{}) (*domEntity.Project, error) {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return nil, err
	}

	identityInfo := i.(identity.Identity)

	var entityModel = &domEntity.Project{}
	var count int64

	if err := dbCon.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ? AND sys_created_by = ?", modelPackageID, identityInfo.Claims.Username).Preload(clause.Associations).Find(&entityModel).Count(&count).Error; err != nil {
		return nil, &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return nil, &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid modelpackage id")}
	}

	// return response
	resp := entityModel

	return resp, nil

}

func (r *ProjectRepo) Delete(id string) error {
	// select db
	dbCon, err := r.handler.GetGormDB(r.dbConnectionName)
	if err != nil {
		return err
	}

	var domEntity = &domEntity.Project{ID: id}
	var count int64

	if err := dbCon.Where("id = ?", domEntity).Find(&domEntity).Count(&count).Error; err != nil {
		return &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	if count == 0 {
		return &sysError.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("invalid deployment id")}
	}

	if err := dbCon.Select(clause.Associations).Delete(domEntity).Error; err != nil {
		return &sysError.SystemError{StatusCode: http.StatusInternalServerError, Err: err}
	}

	return nil
}
