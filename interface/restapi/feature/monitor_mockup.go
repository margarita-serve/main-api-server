package feature

import (
	"mime/multipart"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	appMonitor "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/application"
	appMonitorDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

func NewMonitorMockup(h *handler.Handler, modelPackageSvc *appModelPackageSvc.ModelPackageService) (*FMonitor, error) {
	var err error

	f := new(FMonitor)
	f.handler = h

	if f.appMonitor, err = appMonitor.NewMonitorApp(h, modelPackageSvc); err != nil {
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
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param body body appMonitorDTO.MonitorCreateRequestDTO true "Create Monitor"
// @Success 200 {object} appMonitorDTO.MonitorCreateResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor [post]
func (f *FMonitor) Create(c echo.Context) error {
	req := new(appMonitorDTO.MonitorCreateRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID

	resp, err := f.appMonitor.MonitorSvc.Create(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

//func (f *FMonitor) Delete(c echo.Context) error {
//	req := new(appMonitorDTO.MonitorDeleteRequestDTO)
//	if err := c.Bind(req); err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//	resp, err := f.appMonitor.MonitorSvc.Delete(req)
//	if err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	return response.OkWithData(resp, c)
//}

//func (f *FMonitor) DriftMonitorActive(c echo.Context) error {
//	req := new(appMonitorDTO.MonitorDriftActiveRequestDTO)
//	if err := c.Bind(req); err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	resp, err := f.appMonitor.MonitorSvc.SetDriftMonitorActive(req)
//	if err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	return response.OkWithData(resp, c)
//}
//
//func (f *FMonitor) DriftMonitorInActive(c echo.Context) error {
//	req := new(appMonitorDTO.MonitorDriftInActiveRequestDTO)
//	if err := c.Bind(req); err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	resp, err := f.appMonitor.MonitorSvc.SetDriftMonitorInActive(req)
//	if err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	return response.OkWithData(resp, c)
//}

// PatchMonitor
// @Summary Patch Monitor Setting
// @Description  모니터링 설정 변경
// @Tags Monitor
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param body body appMonitorDTO.MonitorPatchRequestDTO true "Patch Monitor"
// @Router        /projects/{projectID}/deployments/{deploymentID}/monitor [patch]
func (f *FMonitor) PatchMonitor(c echo.Context) error {
	req := new(appMonitorDTO.MonitorPatchRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	deploymentID := c.Param("deploymentID")
	req.DeploymentID = deploymentID
	reqDrift := &appMonitorDTO.MonitorDriftPatchRequestDTO{
		DeploymentID:     req.DeploymentID,
		DataDriftSetting: req.DataDriftSetting,
	}

	_, err := f.appMonitor.MonitorSvc.PatchDriftMonitorSetting(reqDrift)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	reqAccuracy := &appMonitorDTO.MonitorAccuracyPatchRequestDTO{
		DeploymentID:    req.DeploymentID,
		AccuracySetting: req.AccuracySetting,
	}

	_, err = f.appMonitor.MonitorSvc.PatchAccuracyMonitorSetting(reqAccuracy)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp := "Monitor setting update success"
	return response.OkWithData(resp, c)
}

// GetDetail
// @Summary Get Feature Detail
// @Description 피쳐 드리프트 디테일
// @Tags Monitor
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "startTime"
// @Param endTime query string true "endTime"
// @Success 200 {object} appMonitorDTO.FeatureDriftGetResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor/detail [get]
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
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "startTime"
// @Param endTime query string true "endTime"
// @Success 200 {object} appMonitorDTO.FeatureDriftGetResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor/drift [get]
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

//func (f *FMonitor) AccuracyMonitorActive(c echo.Context) error {
//	req := new(appMonitorDTO.MonitorAccuracyActiveRequestDTO)
//	if err := c.Bind(req); err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	resp, err := f.appMonitor.MonitorSvc.SetAccuracyMonitorActive(req)
//	if err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	return response.OkWithData(resp, c)
//}

//func (f *FMonitor) AccuracyMonitorInActive(c echo.Context) error {
//	req := new(appMonitorDTO.MonitorAccuracyInActiveRequestDTO)
//	if err := c.Bind(req); err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	resp, err := f.appMonitor.MonitorSvc.SetAccuracyMonitorInActive(req)
//	if err != nil {
//		return f.translateErrorMessage(err, c)
//	}
//
//	return response.OkWithData(resp, c)
//}

// GetAccuracy
// @Summary Get Accuracy
// @Description 정확도 조회
// @Tags Monitor
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param type query string true "type"
// @Param startTime query string true "startTime"
// @Param endTime query string true "endTime"
// @Success 200 {object} appMonitorDTO.AccuracyGetResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor/accuracy [get]
func (f *FMonitor) GetAccuracy(c echo.Context) error {
	req := new(appMonitorDTO.AccuracyGetRequestDTO)
	deploymentID := c.Param("deploymentID")
	modelHistoryID := c.QueryParam("modelHistoryID")
	Mtype := c.QueryParam("type")
	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")
	req.DeploymentID = deploymentID
	req.ModelHistoryID = modelHistoryID
	req.Type = Mtype
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
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Success 200 {object} appMonitorDTO.MonitorGetSettingResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor [get]
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
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param actualResponse query string true "actualResponse"
// @Router      /projects/{projectID}/deployments/{deploymentID}/monitor/actual [post]
func (f *FMonitor) UploadActual(c echo.Context) error {
	deploymentID := c.Param("deploymentID")
	actualResponse := c.QueryParam("actualResponse")

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
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "startTime"
// @Param endTime query string true "endTime"
// @Success 200 {object} appMonitorDTO.FeatureDriftGetResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor/detail/graph [get]
func (f *FMonitor) GetFeatureDetailGraph(c echo.Context) error {
	req := new(appMonitorDTO.FeatureDriftGetRequestDTO)
	deploymentID := c.Param("deploymentID")
	modelHistoryID := c.QueryParam("modelHistoryID")
	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")
	req.DeploymentID = deploymentID
	req.ModelHistoryID = modelHistoryID
	req.StartTime = startTime
	req.EndTime = endTime

	resp, err := f.appMonitor.MonitorSvc.GetFeatureDetailGraph(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// GetDriftGraph
// @Summary Get Drift Graph
// @Description 드리프트 그래프
// @Tags Monitor
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param deploymentID path string true "deploymentID"
// @Param modelHistoryID query string true "modelHistoryID"
// @Param startTime query string true "startTime"
// @Param endTime query string true "endTime"
// @Success 200 {object} appMonitorDTO.FeatureDriftGetResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor/drift/graph [get]
func (f *FMonitor) GetDriftGraph(c echo.Context) error {
	req := new(appMonitorDTO.FeatureDriftGetRequestDTO)
	deploymentID := c.Param("deploymentID")
	modelHistoryID := c.QueryParam("modelHistoryID")
	startTime := c.QueryParam("startTime")
	endTime := c.QueryParam("endTime")
	req.DeploymentID = deploymentID
	req.ModelHistoryID = modelHistoryID
	req.StartTime = startTime
	req.EndTime = endTime

	resp, err := f.appMonitor.MonitorSvc.GetFeatureDriftGraph(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// GetMonitorStatusList
// @Summary Get Monitor Status List
// @Description 모니터 상태 리스트 조회
// @Tags Monitor
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param deploymentsID path string true "deploymentsID"
// @Success 200 {object} appMonitorDTO.MonitorGetStatusListResponseDTO
// @Router       /projects/{projectID}/deployments/{deploymentID}/monitor/statuses [get]
func (f *FMonitor) GetMonitorStatusList(c echo.Context) error {
	req := new(appMonitorDTO.MonitorGetStatusListRequestDTO)
	deploymentsID := c.QueryParam("deploymentsID")
	req.DeploymentsID = deploymentsID

	resp, err := f.appMonitor.MonitorSvc.GetMonitorStatusList(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	return response.OkWithData(resp, c)
}
