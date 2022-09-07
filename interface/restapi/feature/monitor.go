package feature

import (
	"mime/multipart"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	appMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application"
	appMonitorDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	appNotiSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

func NewMonitor(h *handler.Handler, modelPackageSvc *appModelPackageSvc.ModelPackageService, NotiSvc *appNotiSvc.NotiService) (*FMonitor, error) {
	var err error

	f := new(FMonitor)
	f.handler = h

	if f.appMonitor, err = appMonitor.NewMonitorApp(h, modelPackageSvc, NotiSvc); err != nil {
		return nil, err
	}

	return f, nil
}

type FMonitor struct {
	BaseFeature
	appMonitor *appMonitor.MonitorApp
}

// Create
// @Summary Test Create
// @Description Test
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appMonitorDTO.MonitorCreateRequestDTO true "Create Monitor"
// @Success 200 {object} appMonitorDTO.MonitorCreateResponseDTO
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor [post]
//func (f *FMonitor) Create(c echo.Context) error {
//	req := new(appMonitorDTO.MonitorCreateRequestDTO)
//	if err := c.Bind(req); err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	deploymentID := c.Param("deploymentID")
//	req.DeploymentID = deploymentID
//
//	resp, err := f.appMonitor.MonitorSvc.Create(req)
//	if err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	return response.OkWithData(resp, c)
//}

// PatchDriftSetting
// @Summary Patch Drift Monitor Setting
// @Description  드리프트 모니터링 설정 변경
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appMonitorDTO.MonitorDriftPatchRequestDTO true "Patch Drift Setting"
// @Security BearerAuth
// @Router        /deployments/{deploymentID}/monitor/drift [patch]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appMonitorDTO.MonitorDriftPatchResponseDTO}}
//func (f *FMonitor) PatchDriftSetting(c echo.Context) error {
//	req := new(appMonitorDTO.MonitorDriftPatchRequestDTO)
//	if err := c.Bind(req); err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//	deploymentID := c.Param("deploymentID")
//	req.DeploymentID = deploymentID
//
//	resp, err := f.appMonitor.MonitorSvc.PatchDriftMonitorSetting(req)
//	if err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	return response.OkWithData(resp, c)
//}

// PatchMonitorSetting
// @Summary Patch Monitor Setting
// @Description  모니터링 설정 변경
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appMonitorDTO.MonitorPatchRequestDTO true "Patch Monitor"
// @Security BearerAuth
// @Router        /deployments/{deploymentID}/monitor [patch]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appMonitorDTO.MonitorPatchResponseDTO}}
func (f *FMonitor) PatchMonitorSetting(c echo.Context) error {
	req := new(appMonitorDTO.MonitorPatchRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID

	resp, err := f.appMonitor.MonitorSvc.PatchMonitorSetting(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// UpdateAssociationID
// @Summary Patch AssociationID
// @Description  AssociationID 패치, 변경이전의 Association ID로 예측한 데이터들은 사용 불가능 합니다. 테스트용 삭제예정.
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param body body appMonitorDTO.UpdateAssociationIDRequestDTO true "Patch AssociationID"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Success 200 {object} appMonitorDTO.UpdateAssociationIDResponseDTO
// @Router        /deployments/{deploymentID}/monitor/association-id [patch]
func (f *FMonitor) UpdateAssociationID(c echo.Context) error {
	req := new(appMonitorDTO.UpdateAssociationIDRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID

	resp, err := f.appMonitor.MonitorSvc.UpdateAssociationID(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// GetDetail
// @Summary Get Feature Detail
// @Description 피쳐 드리프트 디테일
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01"
// @Param endTime query string true "example=2022-05-05:01"
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/detail [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appMonitorDTO.FeatureDriftGetResponseDTO}}
func (f *FMonitor) GetDetail(c echo.Context) error {
	req := new(appMonitorDTO.FeatureDriftGetRequestDTO)
	deploymentID := c.Param("deploymentID")
	modelHistoryID := c.QueryParam("modelHistoryID")
	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")
	req.DeploymentID = deploymentID
	req.ModelHistoryID = modelHistoryID
	req.StartTime = startTime
	req.EndTime = endTime

	resp, err := f.appMonitor.MonitorSvc.GetFeatureDetail(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// GetDrift
// @Summary Get Data Drift
// @Description 데이터 드리프트
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01"
// @Param endTime query string true "example=2022-08-01:01"
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/drift [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appMonitorDTO.FeatureDriftGetResponseDTO}}
func (f *FMonitor) GetDrift(c echo.Context) error {
	req := new(appMonitorDTO.FeatureDriftGetRequestDTO)
	deploymentID := c.Param("deploymentID")
	modelHistoryID := c.QueryParam("modelHistoryID")
	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")
	req.DeploymentID = deploymentID
	req.ModelHistoryID = modelHistoryID
	req.StartTime = startTime
	req.EndTime = endTime

	resp, err := f.appMonitor.MonitorSvc.GetFeatureDrift(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// GetAccuracy
// @Summary Get Accuracy
// @Description 정확도 조회
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param type query string true "timeline or aggregation"
// @Param startTime query string true "example=2022-05-05:01"
// @Param endTime query string true "example=2022-08-01:01"
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/accuracy [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appMonitorDTO.AccuracyGetResponseDTO}}
func (f *FMonitor) GetAccuracy(c echo.Context) error {
	req := new(appMonitorDTO.AccuracyGetRequestDTO)
	deploymentID := c.Param("deploymentID")
	modelHistoryID := c.QueryParam("modelHistoryID")
	mType := c.QueryParam("type")
	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")
	req.DeploymentID = deploymentID
	req.ModelHistoryID = modelHistoryID
	req.Type = mType
	req.StartTime = startTime
	req.EndTime = endTime

	resp, err := f.appMonitor.MonitorSvc.GetAccuracy(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

func (f *FMonitor) GetByID(c echo.Context) error {
	req := new(appMonitorDTO.MonitorGetByIDRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appMonitor.MonitorSvc.GetByID(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// GetMonitorSetting
// @Summary Get Monitor Setting
// @Description 모니터 설정 조회
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appMonitorDTO.MonitorGetSettingResponseDTO}}
func (f *FMonitor) GetMonitorSetting(c echo.Context) error {
	req := new(appMonitorDTO.MonitorGetSettingRequestDTO)

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID

	resp, err := f.appMonitor.MonitorSvc.GetMonitorSetting(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// UploadActual
// @Summary Upload Actual file
// @Description upload Actual file
// @Tags Monitor
// @Accept json
// @Produce json
// @Param file formData file true "actual file upload"
// @Param deploymentID path string true "deploymentID"
// @Param targetLabel query string true "target column name"
// @Param associationColumn query string true "association column id"
// @Security BearerAuth
// @Router      /deployments/{deploymentID}/monitor/actual [post]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appMonitorDTO.UploadActualResponseDTO}}
func (f *FMonitor) UploadActual(c echo.Context) error {
	deploymentID := c.Param("deploymentID")
	actualResponse := c.QueryParam("targetLabel")
	associationColumn := c.QueryParam("associationColumn")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			panic(err)
		}
	}(src)
	req := new(appMonitorDTO.UploadActualRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	req.DeploymentID = deploymentID
	req.File = src
	req.FileName = file.Filename
	req.ActualResponse = actualResponse
	req.AssociationColumn = associationColumn

	resp, err := f.appMonitor.MonitorSvc.UploadActual(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// GetFeatureDetailGraph
// @Summary Get Feature Detail Graph
// @Description 피쳐 디테일 그래프
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01"
// @Param endTime query string true "example=2022-08-01:01"
// @Success 200 string html
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/graph/detail [get]
func (f *FMonitor) GetFeatureDetailGraph(c echo.Context) error {

	return response.Ok(c)
}

// GetDriftGraph
// @Summary Get Drift Graph
// @Description 드리프트 그래프
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01"
// @Param endTime query string true "example=2022-08-01:01"
// @Success 200 string html
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/graph/drift [get]
func (f *FMonitor) GetDriftGraph(c echo.Context) error {

	return response.Ok(c)
}

// GetPredictionOverTimeGraph
// @Summary Get PredictionOverTime Graph
// @Description 예측값 변화 그래프
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01"
// @Param endTime query string true "example=2022-08-01:01"
// @Success 200 string html
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/graph/prediction-over-time [get]
func (f *FMonitor) GetPredictionOverTimeGraph(c echo.Context) error {

	return response.Ok(c)
}

// GetAccuracyGraph
// @Summary Get Accuracy Graph
// @Description 정확도 그래프
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01 (UTC+0)"
// @Param endTime query string true "example=2022-08-01:01 (UTC+0)"
// @Success 200 string html
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/graph/accuracy [get]
func (f *FMonitor) GetAccuracyGraph(c echo.Context) error {

	return response.Ok(c)
}

// GetPredictedActual
// @Summary Get PredictedActual Graph
// @Description 예측값 변화 그래프
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01 (UTC+0)"
// @Param endTime query string true "example=2022-08-01:01 (UTC+0)"
// @Success 200 string html
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/graph/predicted-actual [get]
func (f *FMonitor) GetPredictedActual(c echo.Context) error {

	return response.Ok(c)
}

// GetServiceGraph
// @Summary Get Service Graph
// @Description 서비스 상태 그래프
// @Tags Monitor
// @Accept json
// @Produce json
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "example=2022-05-05:01 (UTC+0)"
// @Param endTime query string true "example=2022-08-01:01 (UTC+0)"
// @Success 200 string html
// @Security BearerAuth
// @Router       /deployments/{deploymentID}/monitor/graph/service [get]
func (f *FMonitor) GetServiceGraph(c echo.Context) error {

	return response.Ok(c)
}
