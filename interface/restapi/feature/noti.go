package feature

import (

	//captcha "git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/feature/captcha"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"
	appSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFNoti new  FNoti
func NewFNoti(h *handler.Handler, emailSvc appSvc.IEmailService, deploymentSvc appSvc.IDeploymentService, projectSvc appSvc.IProjectService, authSvc appSvc.IAuthService, governanceHistorySvc appSvc.IGovernanceHistoryService) (*FNoti, error) {
	var err error

	f := new(FNoti)
	f.handler = h

	if f.appNoti, err = application.NewNotiApp(h, emailSvc, deploymentSvc, projectSvc, authSvc, governanceHistorySvc); err != nil {
		return nil, err
	}

	return f, nil
}

// FNoti feature Noti
type FNoti struct {
	BaseFeature
	appNoti *application.NotiApp
}

func (f *FNoti) Create(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	// if !i.IsLogin || i.IsAnonymous {
	// 	return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	// }

	//req := new(appDTO.NotiRequestDTO)
	// if err := c.Bind(req); err != nil {
	// 	return f.translateErrorMessage(err, c)
	// }

	// req := &appDTO.NotiRequestDTO{
	// 	DeploymentID: "aaaaaa",
	// 	NotiCategory: "Datadrift",
	// 	Data:         {"adf": "asdf"},
	// }

	req := &appDTO.NotiRequestDTO{}
	req.DeploymentID = "cbpjvgfr2g4ng38tb86g"
	req.NotiCategory = "Datadrift"
	req.AdditionalData = "status is Failling"

	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	err = f.appNoti.NotiSvc.SendNoti(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.Ok(c)

}
