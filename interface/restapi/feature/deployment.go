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
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	appMonitorSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFDeployment new FDeployment
func NewDeployment(h *handler.Handler, modelPackageSvc *appModelPackageSvc.ModelPackageService, monitorSvc *appMonitorSvc.MonitorService) (*FDeployment, error) {
	var err error

	f := new(FDeployment)
	f.handler = h

	if f.appDeployment, err = appDeployment.NewDeploymentApp(h, modelPackageSvc, monitorSvc); err != nil {
		return nil, err
	}

	return f, nil
}

// FDeployment represent Email Feature
type FDeployment struct {
	BaseFeature
	appDeployment *appDeployment.DeploymentApp
}

//bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MiwiVVVJRCI6IjM1MzlkNTQ1LWU2YmUtNDI0Yi1hYTZhLTcyMmQ4OTE0NGZmYiIsIlVzZXJuYW1lIjoiSkggSFdBTkcxIiwiTmlja05hbWUiOiJqaHl1biIsIkF1dGhvcml0eUlEIjoiZ3JvdXA6ZGVmYXVsdCIsImV4cCI6MTY1ODQ3MDM2OCwiaXNzIjoiS09SRVNFUlZFIiwibmJmIjoxNjU3ODY0NTY4fQ.uA7Pg_yMAXC1hGTGk_p1vcNxulk0GZyVEcyeSBE1NsA

// @Summary Create Deployment
// @Description  배포 생성
// @Tags Deployment
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param body body appDeploymentDTO.CreateDeploymentRequestDTO true "Create Deployment"
// @Success 200 {object} appDeploymentDTO.CreateDeploymentResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/deployments [post]
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
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.Create(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Success 200 {object} appDeploymentDTO.DeleteDeploymentResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/deployments/{deploymentID} [delete]
func (f *FDeployment) Delete(c echo.Context) error {
	//
	req := new(appDeploymentDTO.DeleteDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.Delete(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param body body appDeploymentDTO.UpdateDeploymentRequestDTO true "Update Deployment Info"
// @Success 200 {object} appDeploymentDTO.UpdateDeploymentResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/deployments/{deploymentID} [patch]
func (f *FDeployment) Update(c echo.Context) error {
	//
	req := new(appDeploymentDTO.UpdateDeploymentRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.UpdateDeployment(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Success 200 {object} appDeploymentDTO.GetDeploymentResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/deployments/{deploymentID} [get]
func (f *FDeployment) GetByID(c echo.Context) error {
	//
	req := new(appDeploymentDTO.GetDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetByID(req)
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
// @Param projectID path string false "projectID"
// @Param name query string false "queryName"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort query string false "sort"
// @Success 200 {object} appDeploymentDTO.GetDeploymentListResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/deployments [get]
func (f *FDeployment) GetList(c echo.Context) error {
	//
	req := new(appDeploymentDTO.GetDeploymentListRequestDTO)

	req.Name = c.QueryParam("name")
	req.Page, _ = strconv.Atoi((c.QueryParam("page")))
	req.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	req.Sort = c.QueryParam("sort")

	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetList(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Success 200 {object} appDeploymentDTO.GetGovernanceHistoryResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/deployments/{deploymentID}/governance-log [get]
func (f *FDeployment) GetGovernanceHistory(c echo.Context) error {
	//
	req := new(appDeploymentDTO.GetGovernanceHistoryRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetGovernanceHistory(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Success 200 {object} appDeploymentDTO.GetModelHistoryResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/deployments/{deploymentID}/model-history [get]
func (f *FDeployment) GetModelHistory(c echo.Context) error {
	//
	req := new(appDeploymentDTO.GetModelHistoryRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.GetModelHistory(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param body body appDeploymentDTO.ReplaceModelRequestDTO true "Create Deployment"
// @Success 200 {object} appDeploymentDTO.ReplaceModelResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router        /projects/{projectID}/deployments/{deploymentID}/replace-model [patch]
func (f *FDeployment) ReplaceModel(c echo.Context) error {

	//...
	req := new(appDeploymentDTO.ReplaceModelRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.ReplaceModel(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param json body string true "application/json" SchemaExample({\n"association_id": ["abcd1234", "abcd1235"], \n"instances": [[1.483887, 1.865988, 2.234620, 1.018782, -2.530891, -1.604642, 0.774676, -0.465148, -0.495225], [1.483887, 1.865988, 2.234620, 1.018782, -2.530891, -1.604642, 0.774676, -0.465148, -0.495225]]\n}) "Json data for prediction"
// @Success 200 {object} appDeploymentDTO.ReplaceModelResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router        /projects/{projectID}/deployments/{deploymentID}/predict [post]
func (f *FDeployment) SendPrediction(c echo.Context) error {

	//...
	req := new(appDeploymentDTO.SendPredictionRequestDTO)
	a, _ := ioutil.ReadAll(c.Request().Body)

	req.JsonData = bytes.NewBuffer(a).String()

	// if err := c.Bind(req); err != nil {
	// 	return f.translateErrorMessage(err, c)
	// }

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.SendPrediction(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Success 200 {object} appDeploymentDTO.ActiveDeploymentResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/deployments/{deploymentID}/active [put]
func (f *FDeployment) Active(c echo.Context) error {
	//
	req := new(appDeploymentDTO.ActiveDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.SetActive(req)
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
// @Param projectID path string false "projectID"
// @Param deploymentID path string true "deploymentID"
// @Success 200 {object} appDeploymentDTO.InActiveDeploymentResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/deployments/{deploymentID}/inactive [put]
func (f *FDeployment) InActive(c echo.Context) error {
	//
	req := new(appDeploymentDTO.InActiveDeploymentRequestDTO)
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appDeployment.DeploymentSvc.SetInActive(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}
