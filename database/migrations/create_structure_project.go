package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	ProjectModel "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewCreateStructureProject new CreateStructureProject
func NewCreateStructureProject(h *handler.Handler) base.IMigration {
	m := new(CreateStructureProject)
	m.handler = h

	return m
}

// CreateStructureProject type
type CreateStructureProject struct {
	BaseMigration
}

// Run migration
func (ms01 *CreateStructureProject) Run(db *gorm.DB) error {
	// DB

	if err := db.AutoMigrate(&ProjectModel.Project{}); err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *CreateStructureProject) Rollback(db *gorm.DB) error {
	var count int64

	if db.Migrator().HasTable(&ProjectModel.Project{}) {
		db.Find(&ProjectModel.Project{}).Count(&count)
		if count == 0 {
			if err := db.Migrator().DropTable(&ProjectModel.Project{}); err != nil {
				return err
			}
		}
	}

	return nil
}
