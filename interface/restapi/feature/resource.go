package feature

import (
	"net/http"
	"strconv"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	appResource "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application"
	appResourceDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFResource new FResource
func NewResource(h *handler.Handler) (*FResource, error) {
	var err error

	f := new(FResource)
	f.handler = h

	if f.appResource, err = appResource.NewResourceApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FResource represent Email Feature
type FResource struct {
	BaseFeature
	appResource *appResource.ResourceApp
}

// @Summary Create PredictionEnv
// @Description  프로젝트 생성
// @Tags PredictionEnv
// @Accept json
// @Produce json
// @Param body body appResourceDTO.CreatePredictionEnvRequestDTO true "Create PredictionEnv"
// @Success 200 {object} appResourceDTO.CreatePredictionEnvResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router     /prediction-envs [post]
func (f *FResource) Create(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.CreatePredictionEnvRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	// predictionEnvID := c.Param("predictionEnvID")
	// req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.PredictionEnvSvc.Create(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Delete PredictionEnv
// @Description 프로젝트 삭제
// @Tags PredictionEnv
// @Accept json
// @Produce json
// @Param predictionEnvID path string true "predictionEnvID"
// @Success 200 {object} appResourceDTO.DeletePredictionEnvResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /prediction-envs/{predictionEnvID} [delete]
func (f *FResource) Delete(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.DeletePredictionEnvRequestDTO)
	predictionEnvID := c.Param("predictionEnvID")
	req.PredictionEnvID = predictionEnvID
	// predictionEnvID := c.Param("predictionEnvID")
	// req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.PredictionEnvSvc.Delete(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Edit PredictionEnv
// @Description 프로젝트 정보수정
// @Tags PredictionEnv
// @Accept json
// @Produce json
// @Param predictionEnvID path string true "predictionEnvID"
// @Param body body appResourceDTO.UpdatePredictionEnvRequestDTO true "Update PredictionEnv Info"
// @Success 200 {object} appResourceDTO.UpdatePredictionEnvResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router     /prediction-envs/{predictionEnvID} [patch]
func (f *FResource) Update(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.UpdatePredictionEnvRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	predictionEnvID := c.Param("predictionEnvID")
	req.PredictionEnvID = predictionEnvID
	// predictionEnvID := c.Param("predictionEnvID")
	// req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.PredictionEnvSvc.UpdatePredictionEnv(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get PredictionEnv
// @Description 프로젝트 상세조회
// @Tags PredictionEnv
// @Accept json
// @Produce json
// @Param predictionEnvID path string true "predictionEnvID"
// @Success 200 {object} appResourceDTO.GetPredictionEnvResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /prediction-envs/{predictionEnvID} [get]
func (f *FResource) GetByID(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.GetPredictionEnvRequestDTO)
	predictionEnvID := c.Param("predictionEnvID")
	req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.PredictionEnvSvc.GetByID(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get PredictionEnv List
// @Description 프로젝트 리스트
// @Tags PredictionEnv
// @Accept json
// @Produce json
// @Param name query string false "queryName"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort query string false "sort"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Success 200 {object} appResourceDTO.GetPredictionEnvListResponseDTO
// @Router      /prediction-envs [get]
func (f *FResource) GetList(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.GetPredictionEnvListRequestDTO)

	req.Name = c.QueryParam("name")
	req.Page, _ = strconv.Atoi((c.QueryParam("page")))
	req.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	req.Sort = c.QueryParam("sort")

	// predictionEnvID := c.Param("predictionEnvID")
	// req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.PredictionEnvSvc.GetList(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Create ClusterInfo
// @Description  클러스터정보 생성
// @Tags ClusterInfo
// @Accept json
// @Produce json
// @Param body body appResourceDTO.CreateClusterInfoRequestDTO true "Create ClusterInfo"
// @Success 200 {object} appResourceDTO.CreateClusterInfoResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router     /cluster-infos [post]
func (f *FResource) CreateClusterInfo(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.CreateClusterInfoRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	// predictionEnvID := c.Param("predictionEnvID")
	// req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.ClusterInfoSvc.Create(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Delete ClusterInfo
// @Description 프로젝트 삭제
// @Tags ClusterInfo
// @Accept json
// @Produce json
// @Param predictionEnvID path string true "predictionEnvID"
// @Success 200 {object} appResourceDTO.DeleteClusterInfoResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /cluster-infos/{clusterInfoID} [delete]
func (f *FResource) DeleteClusterInfo(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.DeleteClusterInfoRequestDTO)
	predictionEnvID := c.Param("predictionEnvID")
	req.ClusterInfoID = predictionEnvID
	// predictionEnvID := c.Param("predictionEnvID")
	// req.ClusterInfoID = predictionEnvID

	resp, err := f.appResource.ClusterInfoSvc.Delete(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Edit ClusterInfo
// @Description  정보수정
// @Tags ClusterInfo
// @Accept json
// @Produce json
// @Param clusterInfoID path string true "clusterInfoID"
// @Param body body appResourceDTO.UpdateClusterInfoRequestDTO true "Update ClusterInfo Info"
// @Success 200 {object} appResourceDTO.UpdateClusterInfoResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router     /cluster-infos/{clusterInfoID} [patch]
func (f *FResource) ClusterInfoUpdate(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.UpdateClusterInfoRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	predictionEnvID := c.Param("predictionEnvID")
	req.ClusterInfoID = predictionEnvID
	// predictionEnvID := c.Param("predictionEnvID")
	// req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.ClusterInfoSvc.UpdateClusterInfo(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get ClusterInfo
// @Description  상세조회
// @Tags ClusterInfo
// @Accept json
// @Produce json
// @Param clusterInfoID path string true "clusterInfoID"
// @Success 200 {object} appResourceDTO.GetClusterInfoResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /cluster-infos/{clusterInfoID} [get]
func (f *FResource) ClusterInfoGetByID(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.GetClusterInfoRequestDTO)
	predictionEnvID := c.Param("predictionEnvID")
	req.ClusterInfoID = predictionEnvID

	resp, err := f.appResource.ClusterInfoSvc.GetByID(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get ClusterInfo List
// @Description  리스트
// @Tags ClusterInfo
// @Accept json
// @Produce json
// @Param name query string false "queryName"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort query string false "sort"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Success 200 {object} appResourceDTO.GetClusterInfoListResponseDTO
// @Router      /cluster-infos [get]
func (f *FResource) ClusterInfoGetList(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appResourceDTO.GetClusterInfoListRequestDTO)

	req.Name = c.QueryParam("name")
	req.Page, _ = strconv.Atoi((c.QueryParam("page")))
	req.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	req.Sort = c.QueryParam("sort")

	// predictionEnvID := c.Param("predictionEnvID")
	// req.PredictionEnvID = predictionEnvID

	resp, err := f.appResource.ClusterInfoSvc.GetList(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}
