package entity

import (
	"testing"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/config"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/initialize"
)

func newConfig(t *testing.T) (*config.Config, error) {
	c, _, err := config.NewConfig("../../../../../conf")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func newHandler(t *testing.T) (*handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}

	h.SetConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		t.Errorf("LoadAllDatabaseConnection: %s", err.Error())
	}

	return h, nil
}

func TestMigratoin(t *testing.T) {
	h, err := newHandler(t)
	if err != nil {
		t.Errorf("newHandler: %s", err.Error())
		return
	}

	cfg, err := h.GetConfig()
	if err != nil {
		t.Errorf("GetConfig: %s", err.Error())
		return
	}

	db, err := h.GetGormDB(cfg.Databases.EmailDB.ConnectionName)
	if err != nil {
		t.Errorf("GetGormDB: %s", err.Error())
		return
	}

	if err := db.AutoMigrate(&Email{}, &EmailTemplate{}, &EmailTemplateVersion{}); err != nil {
		t.Errorf("AutoMigrate: %s", err.Error())
	}
}
