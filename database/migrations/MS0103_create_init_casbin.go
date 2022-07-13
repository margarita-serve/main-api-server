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

// NewMS0103CreateAndInitCasbin new MS0103CreateAndInitCasbin
func NewMS0103CreateAndInitCasbin(h *handler.Handler) base.IMigration {
	m := new(MS0103CreateAndInitCasbin)
	m.handler = h

	return m
}

// MS0103CreateAndInitCasbin type
type MS0103CreateAndInitCasbin struct {
	BaseMigration
}

// CasbinRule Optimation

// IamCasbinRule type
// Original: https://github.com/casbin/gorm-adapter/blob/master/adapter.go#L31
type IamCasbinRule struct {
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
func (c *IamCasbinRule) TableName() string {
	return "iam_casbin_rule"
}

// Run migration
func (ms01 *MS0103CreateAndInitCasbin) Run(db *gorm.DB) error {
	// DB Identity

	if err := db.AutoMigrate(&IamCasbinRule{}); err != nil {
		return err
	}

	cPs := []IamCasbinRule{
		// role:system - email
		{PType: "p", V0: "role:system", V1: "system.module.email.send", V2: "EXECUTE"},
		{PType: "p", V0: "role:system", V1: "system.module.email.template.findbycode", V2: "READ"},

		// role:admin - email
		{PType: "p", V0: "role:admin", V1: "/api/v1/email/send", V2: "POST"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/email/templates/list-all", V2: "GET"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/email/template", V2: "POST"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/email/template/*", V2: "GET"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/email/template/update/*", V2: "PUT"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/email/template/set-active/*", V2: "PUT"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/email/template/*", V2: "DELETE"},

		// role:admin - geolocation.country
		{PType: "p", V0: "role:admin", V1: "/api/v1/geolocation/countries/list-all", V2: "GET"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/geolocation/countries/indexer/refresh", V2: "POST"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/geolocation/countries/indexer/search", V2: "POST"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/geolocation/country", V2: "POST"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/geolocation/country/*", V2: "GET"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/geolocation/country/*", V2: "PUT"},
		{PType: "p", V0: "role:admin", V1: "/api/v1/geolocation/country/*", V2: "DELETE"},
		// role:admin - covid19
		{PType: "p", V0: "role:admin", V1: "/api/v1/covid19/current/by-country", V2: "POST"},

		// role:default - geolocation.country
		{PType: "p", V0: "role:default", V1: "/api/v1/geolocation/countries/list-all", V2: "GET"},
		{PType: "p", V0: "role:default", V1: "/api/v1/geolocation/countries/indexer/search", V2: "POST"},
		{PType: "p", V0: "role:default", V1: "/api/v1/geolocation/country/*", V2: "GET"},
		// role:default - covid19
		{PType: "p", V0: "role:default", V1: "/api/v1/covid19/current/by-country", V2: "POST"},

		// group -> role (for flexibility)
		{PType: "g", V0: "group:system", V1: "role:system"},
		{PType: "g", V0: "group:admin", V1: "role:admin"},
		{PType: "g", V0: "group:default", V1: "role:default"},
	}

	if err := db.Create(&cPs).Error; err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0103CreateAndInitCasbin) Rollback(db *gorm.DB) error {
	return nil
}
