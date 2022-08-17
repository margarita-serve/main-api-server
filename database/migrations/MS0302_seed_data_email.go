package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/migrations/data"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMS0301SeedDataEmail new MS0301SeedDataEmail
func NewMS0301SeedDataEmail(h *handler.Handler) base.IMigration {
	m := new(MS0301SeedDataEmail)
	m.handler = h

	return m
}

// MS0301SeedDataEmail type
type MS0301SeedDataEmail struct {
	BaseMigration
}

// Run migration
func (ms01 *MS0301SeedDataEmail) Run(db *gorm.DB) error {

	// data01
	eTpl := data.EmailTemplate01()
	eTpl.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTpl).Error; err != nil {
		return err
	}

	eTplVer := data.EmailTemplate01Version()
	eTplVer.EmailTemplateID = eTpl.ID
	eTplVer.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTplVer).Error; err != nil {
		return err
	}

	eTpl.DefaultVersionID = eTplVer.ID
	if err := db.Save(&eTpl).Error; err != nil {
		return err
	}

	// data02
	eTpl = data.EmailTemplate02()
	eTpl.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTpl).Error; err != nil {
		return err
	}

	eTplVer = data.EmailTemplate02Version()
	eTplVer.EmailTemplateID = eTpl.ID
	eTplVer.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTplVer).Error; err != nil {
		return err
	}

	eTpl.DefaultVersionID = eTplVer.ID
	if err := db.Save(&eTpl).Error; err != nil {
		return err
	}

	// data03
	eTpl = data.EmailTemplate03()
	eTpl.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTpl).Error; err != nil {
		return err
	}

	eTplVer = data.EmailTemplate03Version()
	eTplVer.EmailTemplateID = eTpl.ID
	eTplVer.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTplVer).Error; err != nil {
		return err
	}

	eTpl.DefaultVersionID = eTplVer.ID
	if err := db.Save(&eTpl).Error; err != nil {
		return err
	}

	// data04
	eTpl = data.EmailTemplate04()
	eTpl.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTpl).Error; err != nil {
		return err
	}

	eTplVer = data.EmailTemplate04Version()
	eTplVer.EmailTemplateID = eTpl.ID
	eTplVer.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTplVer).Error; err != nil {
		return err
	}

	eTpl.DefaultVersionID = eTplVer.ID
	if err := db.Save(&eTpl).Error; err != nil {
		return err
	}

	// data05
	eTpl = data.EmailTemplate05()
	eTpl.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTpl).Error; err != nil {
		return err
	}

	eTplVer = data.EmailTemplate05Version()
	eTplVer.EmailTemplateID = eTpl.ID
	eTplVer.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTplVer).Error; err != nil {
		return err
	}

	eTpl.DefaultVersionID = eTplVer.ID
	if err := db.Save(&eTpl).Error; err != nil {
		return err
	}

	// data06
	eTpl = data.EmailTemplate06()
	eTpl.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTpl).Error; err != nil {
		return err
	}

	eTplVer = data.EmailTemplate06Version()
	eTplVer.EmailTemplateID = eTpl.ID
	eTplVer.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTplVer).Error; err != nil {
		return err
	}

	eTpl.DefaultVersionID = eTplVer.ID
	if err := db.Save(&eTpl).Error; err != nil {
		return err
	}

	// data07
	eTpl = data.EmailTemplate07()
	eTpl.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTpl).Error; err != nil {
		return err
	}

	eTplVer = data.EmailTemplate07Version()
	eTplVer.EmailTemplateID = eTpl.ID
	eTplVer.CreatedBy = "system.koreserve@installation"
	if err := db.Create(&eTplVer).Error; err != nil {
		return err
	}

	eTpl.DefaultVersionID = eTplVer.ID
	if err := db.Save(&eTpl).Error; err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0301SeedDataEmail) Rollback(db *gorm.DB) error {
	return nil
}
