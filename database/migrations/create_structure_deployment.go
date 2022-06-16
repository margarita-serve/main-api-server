package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	deploymentModel "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewCreateStructureDeployment new CreateStructureDeployment
func NewCreateStructureDeployment(h *handler.Handler) base.IMigration {
	m := new(CreateStructureDeployment)
	m.handler = h

	return m
}

// CreateStructureDeployment type
type CreateStructureDeployment struct {
	BaseMigration
}

// Run migration
func (ms01 *CreateStructureDeployment) Run(db *gorm.DB) error {
	// DB

	if err := db.AutoMigrate(&deploymentModel.Deployment{}, &deploymentModel.ModelHistory{}); err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *CreateStructureDeployment) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&deploymentModel.Deployment{}) {
		db.Find(&deploymentModel.Deployment{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&deploymentModel.Deployment{}); err != nil {
				return err
			}
		}
	}

	if db.Migrator().HasTable(&deploymentModel.ModelHistory{}) {
		db.Find(&deploymentModel.ModelHistory{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&deploymentModel.ModelHistory{}); err != nil {
				return err
			}
		}
	}

	return nil
}
