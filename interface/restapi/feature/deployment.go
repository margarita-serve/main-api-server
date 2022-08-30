package feature

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	appDeployment "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application"
	appDeploymentDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFDeployment new FDeployment
func NewDeployment(h *handler.Handler) (*FDeployment, error) {

	f := new(FDeployment)
	f.handler = h

	DeploymentApp, err := h.GetApp("deployment")
	if err != nil {
		return nil, err
	}

	f.appDeployment = DeploymentApp.(*appDeployment.DeploymentApp)

	return f, nil
}

// FDeployment represent Email Feature
type FDeployment struct {
	BaseFeature
	appDeployment *appDeployment.DeploymentApp
}

// @Summary Create Deployment
// @Description  배포 생성
// @Tags Deployment
// @Accept json
// @Produce json
// @Param body body appDeploymentDTO.CreateDeploymentRequestDTO true "Create Deployment"
// @Security BearerAuth
// @Router     /deployments [post]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.CreateDeploymentResponseDTO}}
func (f *FDeployment) Create(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.CreateDeploymentRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.Create(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Delete Deployment
// @Description 배포 삭제
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID} [delete]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.DeleteDeploymentResponseDTO}}
func (f *FDeployment) Delete(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.DeleteDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.Delete(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Edit Deployment
// @Description 배포 정보수정
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appDeploymentDTO.UpdateDeploymentRequestDTO true "Update Deployment Info"
// @Security BearerAuth
// @Router     /deployments/{deploymentID} [patch]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.UpdateDeploymentResponseDTO}}
func (f *FDeployment) Update(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.UpdateDeploymentRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.UpdateDeployment(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get Deployment
// @Description 배포 상세조회
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID} [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.GetDeploymentResponseDTO}}
func (f *FDeployment) GetByID(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.GetDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetByID(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp.URI = f.getPredictionURL(c)

	return response.OkWithData(resp, c)

}

func (f *FDeployment) getPredictionURL(c echo.Context) string {
	//...
	//cfg, _ := f.handler.GetConfig()
	hostname := c.Request().Host
	//port := cfg.Applications.Servers.RestAPI.Options.Listener.Port

	predictURI := fmt.Sprintf("%s%s/%s", hostname, c.Request().RequestURI, "predict")

	return predictURI
}

// @Summary Get Deployment List
// @Description 배포 리스트
// @Tags Deployment
// @Accept json
// @Produce json
// @Param name query string false "queryName"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort query string false "sort"
// @Security BearerAuth
// @Router      /deployments [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.GetDeploymentListResponseDTO}}
func (f *FDeployment) GetList(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.GetDeploymentListRequestDTO)

	req.Name = c.QueryParam("name")
	req.Page, _ = strconv.Atoi((c.QueryParam("page")))
	req.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	req.Sort = c.QueryParam("sort")

	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetList(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get Governance Log
// @Description 배포 거버넌스 로그 조회
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/governance-log [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.GetGovernanceHistoryResponseDTO}}
func (f *FDeployment) GetGovernanceHistory(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.GetGovernanceHistoryRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetGovernanceHistory(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get Model History
// @Description 배포 모델 변경이력 조회
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/model-history [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.GetModelHistoryResponseDTO}}
func (f *FDeployment) GetModelHistory(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.GetModelHistoryRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetModelHistory(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Replace Model
// @Description  배포 모델변경
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appDeploymentDTO.ReplaceModelRequestDTO true "Create Deployment"
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/replace-model [patch]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.ReplaceModelResponseDTO}}
func (f *FDeployment) ReplaceModel(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}
	//...
	req := new(appDeploymentDTO.ReplaceModelRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.ReplaceModel(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// @Summary Send Prediction
// @Description  예측 요청
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body string true "application/json" SchemaExample({\n"association_id": ["abcd1234"], \n"instances": [[-122.12,	37.68,	45.0,	2179.0,	401.0,	1159.0,	399.0,	3.4839]]\n}) "Json data for prediction"
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/predict [post]
// @Success 200 {object} response.RootResponse{}
func (f *FDeployment) SendPrediction(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	//...
	req := new(appDeploymentDTO.SendPredictionRequestDTO)
	a, _ := ioutil.ReadAll(c.Request().Body)

	req.JsonData = bytes.NewBuffer(a).String()
	//cbhj65fr2g4h69ej9k3g
	// if err := c.Bind(req); err != nil {
	// 	return f.translateErrorMessage(err, c)
	// }

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.SendPrediction(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// @Summary Active Deployment
// @Description 배포 활성화
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/active [put]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.ActiveDeploymentResponseDTO}}
func (f *FDeployment) Active(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.ActiveDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.SetActive(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Inactive Deployment
// @Description 배포 비활성화
// @Tags Deployment
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/inactive [put]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appDeploymentDTO.InActiveDeploymentResponseDTO}}
func (f *FDeployment) InActive(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appDeploymentDTO.InActiveDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.SetInActive(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}
