package feature

import (
	"net/http"
	"strconv"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/interface/restapi/response"
	appProject "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application"
	appProjectDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/dto"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFProject new FProject
func NewProject(h *handler.Handler) (*FProject, error) {
	f := new(FProject)
	f.handler = h

	ProjectApp, err := h.GetApp("project")
	if err != nil {
		return nil, err
	}

	f.appProject = ProjectApp.(*appProject.ProjectApp)

	return f, nil
}

// FProject represent Email Feature
type FProject struct {
	BaseFeature
	appProject *appProject.ProjectApp
}

// @Summary Create Project
// @Description  프로젝트 생성
// @Tags Project
// @Accept json
// @Produce json
// @Param body body appProjectDTO.CreateProjectRequestDTO true "Create Project"
// @Security BearerAuth
// @Router     /projects [post]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appProjectDTO.CreateProjectResponseDTO}}
func (f *FProject) Create(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appProjectDTO.CreateProjectRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}

	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appProject.ProjectSvc.Create(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Delete Project
// @Description 프로젝트 삭제
// @Tags Project
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Security BearerAuth
// @Router      /projects/{projectID} [delete]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appProjectDTO.DeleteProjectResponseDTO}}
func (f *FProject) Delete(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appProjectDTO.DeleteProjectRequestDTO)
	projectID := c.Param("projectID")
	req.ProjectID = projectID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appProject.ProjectSvc.Delete(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Edit Project
// @Description 프로젝트 정보수정
// @Tags Project
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Param body body appProjectDTO.UpdateProjectRequestDTO true "Update Project Info"
// @Security BearerAuth
// @Router     /projects/{projectID} [patch]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appProjectDTO.UpdateProjectResponseDTO}}
func (f *FProject) Update(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appProjectDTO.UpdateProjectRequestDTO)
	if err := c.Bind(req); err != nil {
		return f.translateErrorMessage(err, c)
	}
	projectID := c.Param("projectID")
	req.ProjectID = projectID
	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appProject.ProjectSvc.UpdateProject(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get Project
// @Description 프로젝트 상세조회
// @Tags Project
// @Accept json
// @Produce json
// @Param projectID path string true "projectID"
// @Security BearerAuth
// @Router      /projects/{projectID} [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appProjectDTO.GetProjectResponseDTO}}
func (f *FProject) GetByID(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appProjectDTO.GetProjectRequestDTO)
	projectID := c.Param("projectID")
	req.ProjectID = projectID

	resp, err := f.appProject.ProjectSvc.GetByID(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}

// @Summary Get Project List
// @Description 프로젝트 리스트
// @Tags Project
// @Accept json
// @Produce json
// @Param name query string false "queryName"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort query string false "sort"
// @Security BearerAuth
// @Router      /projects [get]
// @Success 200 {object} response.RootResponse{response=response.Response{result=appProjectDTO.GetProjectListResponseDTO}}
func (f *FProject) GetList(c echo.Context) error {
	//identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appProjectDTO.GetProjectListRequestDTO)

	req.Name = c.QueryParam("name")
	req.Page, _ = strconv.Atoi((c.QueryParam("page")))
	req.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	req.Sort = c.QueryParam("sort")

	// projectID := c.Param("projectID")
	// req.ProjectID = projectID

	resp, err := f.appProject.ProjectSvc.GetList(req, i)
	if err != nil {
		return f.translateErrorMessage(err, c)
	}

	return response.OkWithData(resp, c)

}
