package database

import (
	"testing"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/initialize"
)

func newHandler(t *testing.T) (*handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	c, err := newConfig(t)
	if err != nil {
		return nil, err
	}
	// set config for test
	c.IAM.Casbin.ModelPath = "../conf/" + c.IAM.Casbin.ModelPath

	h.SetConfig(c)
	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, err
	}

	return h, nil
}

func TestM01FirstInstall(t *testing.T) {
	h, err := newHandler(t)
	if err != nil {
		t.Errorf("newHandler: %s", err.Error())
	}

	if err := M01FirstInstall(h); err != nil {
		t.Errorf("M01FirstInstall: %s", err.Error())
	}
}
