package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	modelPackageModel "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewCreateStructureModelPackage new CreateStructureModelPackage
func NewCreateStructureModelPackage(h *handler.Handler) base.IMigration {
	m := new(CreateStructureModelPackage)
	m.handler = h

	return m
}

// CreateStructureModelPackage type
type CreateStructureModelPackage struct {
	BaseMigration
}

// Run migration
func (ms01 *CreateStructureModelPackage) Run(db *gorm.DB) error {
	// DB

	if err := db.AutoMigrate(&modelPackageModel.ModelPackage{}); err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *CreateStructureModelPackage) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&modelPackageModel.ModelPackage{}) {
		db.Find(&modelPackageModel.ModelPackage{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&modelPackageModel.ModelPackage{}); err != nil {
				return err
			}
		}
	}

	return nil
}
