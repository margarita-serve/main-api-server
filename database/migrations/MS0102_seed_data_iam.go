package migrations

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/database/base"
	iamEntt "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/auths/domain/entity"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/utils"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewMS0102SeedDataIAM new MS0102SeedDataIAM
func NewMS0102SeedDataIAM(h *handler.Handler) base.IMigration {
	m := new(MS0102SeedDataIAM)
	m.handler = h

	return m
}

// MS0102SeedDataIAM type
type MS0102SeedDataIAM struct {
	BaseMigration
}

// Run migration
func (ms01 *MS0102SeedDataIAM) Run(db *gorm.DB) error {
	// DB Identity
	cfg, err := ms01.handler.GetConfig()
	if err != nil {
		return err
	}

	// 과제 연동용 유저생성
	// add default user: user
	defaultUser := iamEntt.SysUser{
		UUID:        utils.GenerateUUID(),
		Username:    cfg.IAM.DefaultUser.Username,
		Password:    utils.MD5([]byte(cfg.IAM.DefaultUser.Password)),
		NickName:    cfg.IAM.DefaultUser.NickName,
		Email:       cfg.IAM.DefaultUser.Email,
		IsActive:    true,
		AuthorityID: cfg.IAM.DefaultUser.AuthorityID,
	}
	defaultUser.CreatedBy = "system.koreserve@installation"

	if err := db.Create(&defaultUser).Error; err != nil {
		return err
	}

	// add default user: super admin
	superAdmin := iamEntt.SysUser{
		UUID:        utils.GenerateUUID(),
		Username:    cfg.IAM.DefaultAdmin.Username,
		Password:    utils.MD5([]byte(cfg.IAM.DefaultAdmin.Password)),
		NickName:    cfg.IAM.DefaultAdmin.NickName,
		Email:       cfg.IAM.DefaultAdmin.Email,
		IsActive:    true,
		AuthorityID: cfg.IAM.DefaultAdmin.AuthorityID,
	}
	superAdmin.CreatedBy = "system.koreserve@installation"

	if err := db.Create(&superAdmin).Error; err != nil {
		return err
	}

	// add default user client app: super admin
	superAdminApp := iamEntt.SysUserClientApps{
		UUID:          utils.GenerateUUID(),
		ClientAppCode: "super-admin-app",
		ClientAppName: "Super Admin Client Application",
		ClientAppDesc: "Default Installation",
		ClientKey:     utils.GenerateClientKey(),
		SecretKey:     utils.GenerateSecretKey(),
		IsActive:      true,
		UserID:        superAdmin.ID,
	}
	superAdminApp.CreatedBy = "system.koreserve@installation"

	if err := db.Create(&superAdminApp).Error; err != nil {
		return err
	}

	return nil
}

// Rollback migration
func (ms01 *MS0102SeedDataIAM) Rollback(db *gorm.DB) error {
	return nil
}
