package repository

import domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/entity"

// IEmailRepo interface
type IWebHookEventRepo interface {
	Save(req *domEntity.WebHookEvent) error
	GetByID(webHookEventID string) (*domEntity.WebHookEvent, error)
	GetForUpdate(webHookEventID string) (*domEntity.WebHookEvent, error)
	GetList(name string, pagination interface{}, filterEntity interface{}) ([]*domEntity.WebHookEvent, interface{}, error)
	Delete(webHookEventID string) error
}
