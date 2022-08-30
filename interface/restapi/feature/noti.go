package feature

import (
	"net/http"
	"strconv"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application"
	appNoti "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application"
	appNotiDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewNoti new  FNoti
func NewNoti(h *handler.Handler) (*FNoti, error) {

	f := new(FNoti)
	f.handler = h

	NotiApp, err := h.GetApp("noti")
	if err != nil {
		return nil, err
	}

	f.appNoti = NotiApp.(*appNoti.NotiApp)

	return f, nil
}

// FNoti feature Noti
type FNoti struct {
	BaseFeature
	appNoti *application.NotiApp
}

// @Summary Create WebHook
// @Description  WebHook 생성
// @Tags Noti
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appNotiDTO.CreateWebHookRequestDTO true "Create WebHook"
// @Security BearerAuth
// @Router     /deployments/{deploymentID}/noti/web-hooks [post]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appNotiDTO.CreateWebHookResponseDTO}}
func (f *FNoti) CreateWebHook(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appNotiDTO.CreateWebHookRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID

	resp, err := f.appNoti.WebHookSvc.Create(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Delete WebHook
// @Description WebHook 삭제
// @Tags Noti
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param webHookID path string true "webHookID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/noti/web-hooks/{webHookID} [delete]
// @Success 200 {object} response.RootResponse{response=response.Response{}}
func (f *FNoti) DeleteWebHook(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appNotiDTO.DeleteWebHookRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	webHookID := c.Param("webHookID")
	req.WebHookID = webHookID

	err = f.appNoti.WebHookSvc.Delete(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.Ok(c)

}

// @Summary Edit WebHook
// @Description WebHook 정보수정
// @Tags Noti
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param webHookID path string true "webHookID"
// @Param body body appNotiDTO.UpdateWebHookRequestDTO true "Update WebHook Info"
// @Security BearerAuth
// @Router     /deployments/{deploymentID}/noti/web-hooks/{webHookID} [patch]
// @Success 200 {object} response.RootResponse{response=response.Response{}}
func (f *FNoti) UpdateWebHook(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appNotiDTO.UpdateWebHookRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	webHookID := c.Param("webHookID")
	req.WebHookID = webHookID

	err = f.appNoti.WebHookSvc.UpdateWebHook(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.Ok(c)

}

// @Summary Get WebHook
// @Description WebHook 상세조회
// @Tags Noti
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param webHookID path string true "webHookID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/noti/web-hooks/{webHookID} [get]
// @Success 200 {object} response.RootResponse{response=response.Response{}}
func (f *FNoti) GetByIDWebHook(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appNotiDTO.GetWebHookRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	webHookID := c.Param("webHookID")
	req.WebHookID = webHookID

	resp, err := f.appNoti.WebHookSvc.GetByID(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get WebHook List
// @Description WebHook 리스트
// @Tags Noti
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param name query string false "queryName"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort query string false "sort"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/noti/web-hooks [get]
// @Success 200 {object} response.RootResponse{response=response.Response{}}
func (f *FNoti) GetListWebHook(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appNotiDTO.GetWebHookListRequestDTO)

	req.Name = c.QueryParam("name")
	req.Page, _ = strconv.Atoi((c.QueryParam("page")))
	req.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	req.Sort = c.QueryParam("sort")

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID

	resp, err := f.appNoti.WebHookSvc.GetList(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Test WebHook Send
// @Description Test WebHook
// @Tags Noti
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appNotiDTO.CreateWebHookRequestDTO true "Test WebHook"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/noti/web-hooks/test [POST]
// @Success 200 {object} response.RootResponse{response=response.Response{}}
func (f *FNoti) TestWebHookSend(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appNotiDTO.TestWebHookRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID

	err = f.appNoti.WebHookSvc.TestWebHookSend(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.Ok(c)

}
