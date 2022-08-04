package service

/*응용서비스 영역
- 리포지터리에서 애그리거트 조회/저장/생성/결과리턴/도메인기능실행
- 트랜잭션 처리
- 이벤트처리
- 접근제어
- 복잡하다면 도메인 로직이 포함되지않았나 의심
- 표현영역과 도메인 영역을 연결하는 매개체 열할
- 한 응용 서비스 클래스에서 한개내지 2~3개의 기능 구현 중복되는 로직은 별도 클래스로 작성해서 포함
*/

import (
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	"github.com/rs/xid"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	appModelPackageDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/dto"
	appModelPackageSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/model_package/application/service"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/project/infrastructure/repository"
)

// ProjectService type
type IModelPackageService interface {
	GetListByProject(req *appModelPackageDTO.GetModelPackageListByProjectRequestDTO) (*appModelPackageDTO.GetModelPackageListByProjectResponseDTO, error)
}

type IDeploymentService interface {
	GetListByProject(req *appModelPackageDTO.InternalGetModelPackageRequestDTO) (*appModelPackageDTO.InternalGetModelPackageResponseDTO, error)
}

type ProjectService struct {
	BaseService
	repo            domRepo.IProjectRepo
	modelPackageSvc IModelPackageService
}

// NewProjectService new ProjectService
func NewProjectService(h *handler.Handler) (*ProjectService, error) {
	var err error

	svc := new(ProjectService)
	svc.handler = h
	// if err := svc.initBaseService(); err != nil {
	// 	return nil, err
	// }

	if svc.repo, err = infRepo.NewProjectRepo(h); err != nil {
		return nil, err
	}

	if svc.modelPackageSvc, err = appModelPackageSvc.NewModelPackageService(h, svc); err != nil {
		return nil, err
	}

	return svc, nil
}

// Create
func (s *ProjectService) Create(req *appDTO.CreateProjectRequestDTO, i identity.Identity) (*appDTO.CreateProjectResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }

	guid := xid.New().String()

	// New deployment domain Instance
	domAggregateProject, err := domEntity.NewProject(
		guid,
		req.Name,
		req.Description,
		i.Claims.Username,
	)
	if err != nil {
		return nil, err
	}

	// domEntity.Validate(domAggregateProject)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregateProject)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.CreateProjectResponseDTO)
	resDTO.ProjectID = domAggregateProject.ID

	return resDTO, nil
}

func (s *ProjectService) Delete(req *appDTO.DeleteProjectRequestDTO, i identity.Identity) (*appDTO.DeleteProjectResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	domAggregateProject, err := s.repo.GetByID(req.ProjectID, i)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.DeleteProjectResponseDTO)

	err = s.repo.Delete(domAggregateProject.ID)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "Project Delete Success"

	return resDTO, nil
}

func (s *ProjectService) UpdateProject(req *appDTO.UpdateProjectRequestDTO, i identity.Identity) (*appDTO.UpdateProjectResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	//userID := "testID"

	//Find Domain Entity
	domAggregateProject, err := s.repo.GetForUpdate(req.ProjectID, i)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UpdateProjectResponseDTO)

	if req.Name != "" {
		domAggregateProject.SetName(req.Name)
	}
	if req.Description != "" {
		domAggregateProject.SetDescription(req.Description)
	}

	// domEntity.Validate(domAggregateProject)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregateProject)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "Project Update Success"

	return resDTO, nil
}

func (s *ProjectService) GetByID(req *appDTO.GetProjectRequestDTO, i identity.Identity) (*appDTO.GetProjectResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.ProjectID, i)
	if err != nil {
		return nil, err
	}
	reqByProject := &appModelPackageDTO.GetModelPackageListByProjectRequestDTO{
		ProjectID: res.ID,
	}

	resModelPackage, err := s.modelPackageSvc.GetListByProject(reqByProject)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.GetProjectResponseDTO)
	resDTO.ProjectID = reqByProject.ProjectID
	resDTO.Name = res.Name
	resDTO.Description = res.Description
	resDTO.ModelPackages = resModelPackage.Rows

	return resDTO, nil
}

func (s *ProjectService) GetList(req *appDTO.GetProjectListRequestDTO, i identity.Identity) (*appDTO.GetProjectListResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	reqP := infRepo.Pagination{
		Limit: req.Limit,
		Page:  req.Page,
		Sort:  req.Sort,
	}

	resultList, pagination, err := s.repo.GetList(req.Name, reqP, i)
	if err != nil {
		return nil, err
	}

	//interface type을 concrete type으로 변환
	//domain layer에서 pagination type을 모르기 때문에 interface type으로 정의 후 변환한다
	p := pagination.(infRepo.Pagination)

	// response dto
	resDTO := new(appDTO.GetProjectListResponseDTO)
	resDTO.Limit = p.Limit
	resDTO.Page = p.Page
	resDTO.TotalRows = p.TotalRows
	resDTO.TotalPages = p.TotalPages

	var listProject []appDTO.GetProjectResponseDTO
	for _, rec := range resultList {
		tmp := new(appDTO.GetProjectResponseDTO)

		tmp.ProjectID = rec.ID
		tmp.Name = rec.Name
		tmp.Description = rec.Description

		listProject = append(listProject, *tmp)
	}

	resDTO.Rows = listProject

	return resDTO, nil
}
