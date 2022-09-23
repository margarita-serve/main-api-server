package repository

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
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
