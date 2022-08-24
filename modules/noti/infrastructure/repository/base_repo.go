package repository

import (
	"math"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"gorm.io/gorm"
)

// BaseRepo type
type BaseRepo struct {
	handler          *handler.Handler
	dbConnectionName string
}

// SetHandler set Handler
func (s *BaseRepo) SetHandler(h *handler.Handler) {
	s.handler = h
}

// SetDBConnectionName set DBConnectionName
func (s *BaseRepo) SetDBConnectionName(v string) {
	s.dbConnectionName = v
}

type Pagination struct {
	Limit      int
	Page       int
	Sort       string
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	// if p.Limit == 0 {
	// 	p.Limit = 10
	// }
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

//Paging 및 Sorting을 위한 코드
func paginate(value interface{}, pagination *Pagination, dbCon *gorm.DB, queryName string) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	var tmpLimit int

	dbCon.Model(value).Where("name like ?", "%"+queryName+"%").Count(&totalRows)
	pagination.TotalRows = totalRows

	if pagination.Limit <= 0 {
		tmpLimit = 1
	} else {
		tmpLimit = pagination.Limit
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(tmpLimit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Where("name like ?", "%"+queryName+"%")
	}
}
