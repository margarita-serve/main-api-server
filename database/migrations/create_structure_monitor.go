package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	monitorModel "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"gorm.io/gorm"
)

func NewCreateStructureMonitor(h *handler.Handler) base.IMigration {
	m := new(CreateStructureMonitor)
	m.handler = h

	return m
}

type CreateStructureMonitor struct {
	BaseMigration
}

func (ms01 *CreateStructureMonitor) Run(db *gorm.DB) error {
	if err := db.AutoMigrate(&monitorModel.Monitor{}); err != nil {
		return err
	}

	return nil
}

func (ms01 *CreateStructureMonitor) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&monitorModel.Monitor{}) {
		db.Find(&monitorModel.Monitor{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&monitorModel.Monitor{}); err != nil {
				return err
			}
		}
	}

	return nil
}
