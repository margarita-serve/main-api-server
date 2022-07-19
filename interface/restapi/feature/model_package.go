package feature

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

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

// @Summary Create ModelPackage
// @Description  모델 패키지 생성
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param body body appModelPackageDTO.CreateModelPackageRequestDTO true "Create ModelPackage"
// @Success 200 {object} appModelPackageDTO.CreateModelPackageResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages [post]
func (f *FModelPackage) Create(c echo.Context) error {
	// identity
	// i, err := f.SetIdentity(c)
	// if err != nil {
	// 	return f.translateErrorMessage(err, c)
	// }
	// if !i.IsLogin || i.IsAnonymous {
	// 	return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	// }
	req := new(appModelPackageDTO.CreateModelPackageRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appModelPackage.ModelPackageSvc.Create(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Delete ModelPackage
// @Description 모델 패키지 삭제
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Success 200 {object} appModelPackageDTO.DeleteModelPackageResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/model-packages/{modelPackageID} [delete]
func (f *FModelPackage) Delete(c echo.Context) error {
	//
	req := new(appModelPackageDTO.DeleteModelPackageRequestDTO)
	modelPackageID := c.Param("modelPackageID")
	req.ModelPackageID = modelPackageID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appModelPackage.ModelPackageSvc.Delete(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Archive ModelPackage
// @Description 모델 패키지 보관
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Success 200 {object} appModelPackageDTO.ArchiveModelPackageResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/model-packages/{modelPackageID}/archive [put]
func (f *FModelPackage) Archive(c echo.Context) error {
	//
	req := new(appModelPackageDTO.ArchiveModelPackageRequestDTO)
	modelPackageID := c.Param("modelPackageID")
	req.ModelPackageID = modelPackageID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appModelPackage.ModelPackageSvc.Archive(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Edit ModelPackage
// @Description 모델 패키지 정보수정
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Param body body appModelPackageDTO.UpdateModelPackageRequestDTO true "Update ModelPackage Info"
// @Success 200 {object} appModelPackageDTO.UpdateModelPackageResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages/{modelPackageID} [patch]
func (f *FModelPackage) Update(c echo.Context) error {
	//
	req := new(appModelPackageDTO.UpdateModelPackageRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	modelPackageID := c.Param("modelPackageID")
	req.ModelPackageID = modelPackageID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appModelPackage.ModelPackageSvc.UpdateModelPackage(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get ModelPackage
// @Description 모델 패키지 상세조회
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Success 200 {object} appModelPackageDTO.GetModelPackageResponseDTO
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router       /projects/{projectID}/model-packages/{modelPackageID} [get]
func (f *FModelPackage) GetByID(c echo.Context) error {
	//
	req := new(appModelPackageDTO.GetModelPackageRequestDTO)
	modelPackageID := c.Param("modelPackageID")
	req.ModelPackageID = modelPackageID
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appModelPackage.ModelPackageSvc.GetByID(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get ModelPackage List
// @Description 모델 패키지 리스트
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param name query string false "queryName"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort query string false "sort"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Success 200 {object} appModelPackageDTO.GetModelPackageListResponseDTO
// @Router       /projects/{projectID}/model-packages [get]
func (f *FModelPackage) GetList(c echo.Context) error {
	//
	req := new(appModelPackageDTO.GetModelPackageListRequestDTO)

	req.Name = c.QueryParam("name")
	req.Page, _ = strconv.Atoi((c.QueryParam("page")))
	req.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	req.Sort = c.QueryParam("sort")

	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appModelPackage.ModelPackageSvc.GetList(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Upload Model File
// @Description 모델 파일 업로드
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Param upfile formData file true "file upload"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages/{modelPackageID}/upload-model [post]
func (f *FModelPackage) UploadModel(c echo.Context) error {
	modelPackageID := c.Param("modelPackageID")
	projectID := c.Param("projectID")
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("upfile")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	req := new(appModelPackageDTO.UploadModelRequestDTO)
	req.ModelPackageID = modelPackageID
	req.File = src
	req.ProjectID = projectID
	req.FileName = file.Filename

	resp, err := f.appModelPackage.ModelPackageSvc.UploadModel(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// @Summary Upload Training Dataset File
// @Description 훈련 데이터셋 파일 업로드
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Param upfile formData file true "file upload"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages/{modelPackageID}/upload-training-dataset [post]
func (f *FModelPackage) UploadTrainingDataset(c echo.Context) error {
	modelPackageID := c.Param("modelPackageID")
	projectID := c.Param("projectID")
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("upfile")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	req := new(appModelPackageDTO.UploadTrainingDatasetRequestDTO)
	req.ModelPackageID = modelPackageID
	req.File = src
	req.ProjectID = projectID
	req.FileName = file.Filename

	resp, err := f.appModelPackage.ModelPackageSvc.UploadTrainingDataset(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// @Summary Upload Holdout Dataset File
// @Description 훈련 검증 데이터셋 파일 업로드
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Param upfile formData file true "file upload"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages/{modelPackageID}/upload-holdout-dataset [post]
func (f *FModelPackage) UploadHoldoutDataset(c echo.Context) error {
	modelPackageID := c.Param("modelPackageID")
	projectID := c.Param("projectID")
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("upfile")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	req := new(appModelPackageDTO.UploadHoldoutDatasetRequestDTO)
	req.ModelPackageID = modelPackageID
	req.File = src
	req.ProjectID = projectID
	req.FileName = file.Filename

	resp, err := f.appModelPackage.ModelPackageSvc.UploadHoldoutDataset(req)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)
}

// @Summary Download Model File
// @Description 모델 파일 다운로드
// @Tags ModelPackage
// @Produce octet-stream
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages/{modelPackageID}/download-model [get]
func (f *FModelPackage) DownloadModelFile(c echo.Context) error {
	modelPackageID := c.Param("modelPackageID")

	fileReader, fileName, err := f.appModelPackage.ModelPackageSvc.GetModelFile(modelPackageID)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	// Add headers to the response
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%q", fileName))
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)

	_, err = io.Copy(c.Response().Writer, fileReader)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	c.Response().WriteHeader(http.StatusOK)

	return nil
}

// @Summary Download Training Dataset File
// @Description 훈련 데이터셋 파일 다운로드
// @Tags ModelPackage
// @Produce octet-stream
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages/{modelPackageID}/download-training-dataset [get]
func (f *FModelPackage) DownloadTrainingDataset(c echo.Context) error {
	modelPackageID := c.Param("modelPackageID")

	fileReader, fileName, err := f.appModelPackage.ModelPackageSvc.GetTrainingDatasetFile(modelPackageID)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	// Add headers to the response
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%q", fileName))
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)

	_, err = io.Copy(c.Response().Writer, fileReader)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	c.Response().WriteHeader(http.StatusOK)

	return nil
}

// @Summary Download Holdout Dataset File
// @Description 홀드아웃 데이터셋 파일 다운로드
// @Tags ModelPackage
// @Produce octet-stream
// @Param projectID path string false "projectID"
// @Param modelPackageID path string true "modelPackageID"
// @Param Authorization header string true "Insert your access token" default(bearer <Add access token here>)
// @Router      /projects/{projectID}/model-packages/{modelPackageID}/download-holdout-dataset [get]
func (f *FModelPackage) DownloadHoldoutDataset(c echo.Context) error {
	modelPackageID := c.Param("modelPackageID")

	fileReader, fileName, err := f.appModelPackage.ModelPackageSvc.GetHoldoutDatasetFile(modelPackageID)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	// Add headers to the response
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%q", fileName))
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)

	_, err = io.Copy(c.Response().Writer, fileReader)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	c.Response().WriteHeader(http.StatusOK)

	return nil
}
