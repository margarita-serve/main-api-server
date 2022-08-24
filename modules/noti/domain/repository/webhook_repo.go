package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/entity"

// IEmailRepo interface
type IWebHookRepo interface {
	Save(req *domEntity.WebHook) error
	GetByID(webHookID string) (*domEntity.WebHook, error)
	GetForUpdate(webHookID string) (*domEntity.WebHook, error)
	GetList(name string, pagination interface{}, filterEntity interface{}) ([]*domEntity.WebHook, interface{}, error)
	GetListByInternal(filterEntity interface{}) ([]*domEntity.WebHook, error)
	Delete(webHookID string) error
}
