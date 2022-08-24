package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	NotiModel "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewCreateStructureNoti new CreateStructureNoti
func NewCreateStructureNoti(h *handler.Handler) base.IMigration {
	m := new(CreateStructureNoti)
	m.handler = h

	return m
}

// CreateStructureNoti type
type CreateStructureNoti struct {
	BaseMigration
}

// Run migration
func (ms01 *CreateStructureNoti) Run(db *gorm.DB) error {
	// DB

	if err := db.AutoMigrate(&NotiModel.WebHook{}, &NotiModel.WebHookEvent{}); err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *CreateStructureNoti) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&NotiModel.WebHook{}) {
		db.Find(&NotiModel.WebHook{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&NotiModel.WebHook{}); err != nil {
				return err
			}
		}
	}

	if db.Migrator().HasTable(&NotiModel.WebHookEvent{}) {
		db.Find(&NotiModel.WebHookEvent{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&NotiModel.WebHookEvent{}); err != nil {
				return err
			}
		}
	}

	return nil
}
