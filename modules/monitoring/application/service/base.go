package service

import "git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

type BaseService struct {
	handler *handler.Handler
}

func (b *BaseService) initBaseService() error {
	return nil
}

func (b *BaseService) initSystemIdentity() error {
	return nil
}
