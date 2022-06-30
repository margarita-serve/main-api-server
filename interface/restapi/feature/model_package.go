package feature

import (
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	appModelPackage "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application"
	appModelPackageDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFModelPackage new FModelPackage
func NewModelPackage(h *handler.Handler) (*FModelPackage, error) {
	var err error

	f := new(FModelPackage)
	f.handler = h

	if f.appModelPackage, err = appModelPackage.NewModelPackageApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FModelPackage represent Email Feature
type FModelPackage struct {
	BaseFeature
	appModelPackage *appModelPackage.ModelPackageApp
}

func (f *FModelPackage) Create(c echo.Context) error {
	// identity
	// i, err := f.SetIdentity(c)
	// if err != nil {
	// 	return f.translateErrorMessage(err, c)
	// }
	// if !i.IsLogin || i.IsAnonymous {
	// 	return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	// }
	req := new(appModelPackageDTO.ModelPackageCreateRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appModelPackage.ModelPackageSvc.Create(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

func (f *FModelPackage) Delete(c echo.Context) error {
	//
	req := new(appModelPackageDTO.ModelPackageDeleteRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appModelPackage.ModelPackageSvc.Delete(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

func (f *FModelPackage) Get(c echo.Context) error {
	//
	req := new(appModelPackageDTO.ModelPackageGetRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	resp, err := f.appModelPackage.ModelPackageSvc.GetByID(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

func (f *FModelPackage) GetByName(c echo.Context) error {
	//

	name := c.QueryParam("name")
	print("NAME:" + name)
	req := new(appModelPackageDTO.ModelPackageGetByNametRequestDTO)
	req.Name = name

	resp, err := f.appModelPackage.ModelPackageSvc.GetByName(req)
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
