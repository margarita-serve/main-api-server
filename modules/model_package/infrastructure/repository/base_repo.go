package repository

import "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

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
