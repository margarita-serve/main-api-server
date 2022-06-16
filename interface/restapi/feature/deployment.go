package feature

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	appDeployment "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application"
	appDeploymentDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFDeployment new FDeployment
func NewDeployment(h *handler.Handler) (*FDeployment, error) {
	var err error

	f := new(FDeployment)
	f.handler = h

	if f.appDeployment, err = appDeployment.NewDeploymentApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FDeployment represent Email Feature
type FDeployment struct {
	BaseFeature
	appDeployment *appDeployment.DeploymentApp
}

// @Summary Deploy Package
// @Description  배포 생성
// @Tags Deploy
// @Accept json
// @Produce json
// @Param body body PostDeploysRequest true "Create Deployment"
// @Success 200 {object} appDeploymentDTO.DeploymentCreateRequestDTO
func (f *FDeployment) Create(c echo.Context) error {
	// identity
	// i, err := f.SetIdentity(c)
	// if err != nil {
	// 	return f.translateErrorMessage(err, c)
	// }
	// if !i.IsLogin || i.IsAnonymous {
	// 	return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	// }
	req := new(appDeploymentDTO.DeploymentCreateRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appDeployment.DeploymentSvc.Create(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary 배포 조회
// @Description
// @Tags Deploy
// @Accept json
// @Produce json
// @Success 200 {object} appDeploymentDTO.DeploymentDeleteResponseDTO
func (f *FDeployment) Delete(c echo.Context) error {
	//
	req := new(appDeploymentDTO.DeploymentDeleteRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appDeployment.DeploymentSvc.Delete(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary 배포 조회
// @Description
// @Tags Deploy
// @Accept json
// @Produce json
// @Success 200 {object} appDeploymentDTO.DeploymentGetResponseDTO
func (f *FDeployment) Get(c echo.Context) error {
	//
	req := new(appDeploymentDTO.DeploymentGetRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appDeployment.DeploymentSvc.GetByID(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary 배포 조회
// @Description
// @Tags Deploy
// @Accept json
// @Produce json
// @Success 200 {object} appDeploymentDTO.DeploymentGetByNameResponseDTO
func (f *FDeployment) GetByName(c echo.Context) error {
	//

	name := c.QueryParam("name")
	print("NAME:" + name)
	req := new(appDeploymentDTO.DeploymentGetByNametRequestDTO)
	req.Name = name

	resp, err := f.appDeployment.DeploymentSvc.GetByName(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary 배포 삭제
// @Description
// @Tags Deploy
// @Accept json
// @Produce json
// @Param modelPackageID path string true "Required. "
// @Success      200  {string}  string    "ok"
// @Router /deployments/{deploymentsId} [delete]
// func Delete(c echo.Context) error {
// 	//...

// 	post_response := new(PostDeploysResponse)
// 	return c.JSONPretty(http.StatusOK, *post_response, "  ")
// }

// @Summary 배포 수정
// @Description
// @Tags Deploy
// @Accept json
// @Produce json
// @Param modelPackageID query string true "Required. "
// @Param body body PatchDeploysRequest true "Archive Package"
// @Success      200  {string}  string    "ok"
// @Router /deployments/{deploymentsId} [patch]
// func Patch(c echo.Context) error {

// 	//...

// 	post_response := new(PostDeploysResponse)
// 	return c.JSONPretty(http.StatusOK, *post_response, "  ")
// }

// @Summary 배포 모델변경
// @Description
// @Tags Deploy
// @Accept json
// @Produce json
// @Param modelPackageID query string true "Required. "
// @Param body body PatchDeploysRequest true "Archive Package"
// @Success      200  {string}  string    "ok"
// @Router /deployments/{deploymentsId} [patch]
// func PatchModel(c echo.Context) error {

// 	//...

// 	post_response := new(PostDeploysResponse)
// 	return c.JSONPretty(http.StatusOK, *post_response, "  ")
// }
