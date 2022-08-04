package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMS0104CreateAndInitCasbin new MS0104CreateAndInitCasbin
func NewMS0104CreateAndInitCasbin(h *handler.Handler) base.IMigration {
	m := new(MS0104CreateAndInitCasbin)
	m.handler = h

	return m
}

// MS0104CreateAndInitCasbin type
type MS0104CreateAndInitCasbin struct {
	BaseMigration
}

// CasbinRule Optimation

// IamCasbinRuleShare type
// Original: https://github.com/casbin/gorm-adapter/blob/master/adapter.go#L31
type IamCasbinRuleShare struct {
	TablePrefix string `gorm:"-"`
	PType       string `gorm:"size:100;index;index:idx_unique,unique"`
	V0          string `gorm:"size:100;index;index:idx_unique,unique"`
	V1          string `gorm:"size:100;index;index:idx_unique,unique"`
	V2          string `gorm:"size:100;index;index:idx_unique,unique"`
	V3          string `gorm:"size:100;index;index:idx_unique,unique"`
	V4          string `gorm:"size:100;index;index:idx_unique,unique"`
	V5          string `gorm:"size:100;index;index:idx_unique,unique"`
}

// TableName func
func (c *IamCasbinRuleShare) TableName() string {
	return "iam_casbin_rule"
}

// Run migration
func (ms01 *MS0104CreateAndInitCasbin) Run(db *gorm.DB) error {
	// DB Identity

	if err := db.AutoMigrate(&IamCasbinRuleShare{}); err != nil {
		return err
	}

	cPs := []IamCasbinRuleShare{
		// role:system - email
		{PType: "p", V0: "JH HWANG1", V1: "cbbl3mnr2g4ji45s9av0", V2: "OWNER"},
	}

	if err := db.Create(&cPs).Error; err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0104CreateAndInitCasbin) Rollback(db *gorm.DB) error {
	return nil
}
