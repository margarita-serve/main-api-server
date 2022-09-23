package application

import (
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// NewProjectApp new ProjectApp
func NewProjectApp(h *handler.Handler) (*ProjectApp, error) {
	var err error

	app := new(ProjectApp)
	app.handler = h

	if app.ProjectSvc, err = appSvc.NewProjectService(h); err != nil {
		return nil, err
	}

	return app, nil
}

// ProjectApp represent DDD Module: Email (Application Layer)
type ProjectApp struct {
	handler    *handler.Handler
	ProjectSvc *appSvc.ProjectService
}
