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
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	"github.com/rs/xid"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/resource/infrastructure/repository"
)

type IClusterInfoService interface {
	GetByIDInternal(req *appDTO.InternalGetClusterInfoRequestDTO, i identity.Identity) (*appDTO.InternalGetClusterInfoResponseDTO, error)
}

type PredictionEnvService struct {
	BaseService
	repo           domRepo.IPredictionEnvRepo
	clusterInfoSvc IClusterInfoService
}

// NewPredictionEnvService new PredictionEnvService
func NewPredictionEnvService(h *handler.Handler, clusterInfoSvc IClusterInfoService) (*PredictionEnvService, error) {
	var err error

	svc := new(PredictionEnvService)
	svc.handler = h
	// if err := svc.initBaseService(); err != nil {
	// 	return nil, err
	// }

	if svc.repo, err = infRepo.NewPredictionEnvRepo(h); err != nil {
		return nil, err
	}

	svc.clusterInfoSvc = clusterInfoSvc

	return svc, nil
}

// Create
func (s *PredictionEnvService) Create(req *appDTO.CreatePredictionEnvRequestDTO, i identity.Identity) (*appDTO.CreatePredictionEnvResponseDTO, error) {
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
	domAggregatePredictionEnv, err := domEntity.NewPredictionEnv(
		guid,
		req.Name,
		req.Description,
		i.Claims.Username,
	)
	if err != nil {
		return nil, err
	}

	// domEntity.Validate(domAggregatePredictionEnv)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregatePredictionEnv)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.CreatePredictionEnvResponseDTO)
	resDTO.PredictionEnvID = domAggregatePredictionEnv.ID

	return resDTO, nil
}

func (s *PredictionEnvService) Delete(req *appDTO.DeletePredictionEnvRequestDTO, i identity.Identity) (*appDTO.DeletePredictionEnvResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	domAggregatePredictionEnv, err := s.repo.GetByID(req.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.DeletePredictionEnvResponseDTO)

	err = s.repo.Delete(domAggregatePredictionEnv.ID)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "PredictionEnv Delete Success"

	return resDTO, nil
}

func (s *PredictionEnvService) UpdatePredictionEnv(req *appDTO.UpdatePredictionEnvRequestDTO, i identity.Identity) (*appDTO.UpdatePredictionEnvResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	//userID := "testID"

	//Find Domain Entity
	domAggregatePredictionEnv, err := s.repo.GetForUpdate(req.PredictionEnvID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UpdatePredictionEnvResponseDTO)

	if req.Name != "" {
		domAggregatePredictionEnv.SetName(req.Name)
	}
	if req.Description != "" {
		domAggregatePredictionEnv.SetDescription(req.Description)
	}
	if req.Description != "" {
		domAggregatePredictionEnv.SetDescription(req.Description)
	}

	// domEntity.Validate(domAggregatePredictionEnv)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregatePredictionEnv)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "PredictionEnv Update Success"

	return resDTO, nil
}

func (s *PredictionEnvService) GetByID(req *appDTO.GetPredictionEnvRequestDTO, i identity.Identity) (*appDTO.GetPredictionEnvResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.PredictionEnvID)
	if err != nil {
		return nil, err
	}
	// response dto
	resDTO := new(appDTO.GetPredictionEnvResponseDTO)

	resDTO.Name = res.Name
	resDTO.Description = res.Description

	return resDTO, nil
}

func (s *PredictionEnvService) GetList(req *appDTO.GetPredictionEnvListRequestDTO, i identity.Identity) (*appDTO.GetPredictionEnvListResponseDTO, error) {
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

	filter := map[string]interface{}{}

	resultList, pagination, err := s.repo.GetList(req.Name, reqP, filter)
	if err != nil {
		return nil, err
	}

	//interface type을 concrete type으로 변환
	//domain layer에서 pagination type을 모르기 때문에 interface type으로 정의 후 변환한다
	p := pagination.(infRepo.Pagination)

	// response dto
	resDTO := new(appDTO.GetPredictionEnvListResponseDTO)
	resDTO.Limit = p.Limit
	resDTO.Page = p.Page
	resDTO.TotalRows = p.TotalRows
	resDTO.TotalPages = p.TotalPages

	var listPredictionEnv []appDTO.GetPredictionEnvResponseDTO
	for _, rec := range resultList {
		tmp := new(appDTO.GetPredictionEnvResponseDTO)

		tmp.PredictionEnvID = rec.ID
		tmp.Name = rec.Name
		tmp.Description = rec.Description

		listPredictionEnv = append(listPredictionEnv, *tmp)
	}

	resDTO.Rows = listPredictionEnv

	return resDTO, nil
}

func (s *PredictionEnvService) GetListByProject(req *appDTO.GetProjectsPredictionEnvsRequestDTO, i identity.Identity) (*appDTO.GetProjectsPredictionEnvsResponseDTO, error) {
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

	filter := map[string]interface{}{
		"project_id": req.ProjectID,
	}

	resultList, pagination, err := s.repo.GetList(req.Name, reqP, filter)
	if err != nil {
		return nil, err
	}

	//interface type을 concrete type으로 변환
	//domain layer에서 pagination type을 모르기 때문에 interface type으로 정의 후 변환한다
	p := pagination.(infRepo.Pagination)

	// response dto
	resDTO := new(appDTO.GetProjectsPredictionEnvsResponseDTO)
	resDTO.Limit = p.Limit
	resDTO.Page = p.Page
	resDTO.TotalRows = p.TotalRows
	resDTO.TotalPages = p.TotalPages

	var listPredictionEnv []appDTO.GetPredictionEnvResponseDTO
	for _, rec := range resultList {
		tmp := new(appDTO.GetPredictionEnvResponseDTO)

		tmp.PredictionEnvID = rec.ID
		tmp.Name = rec.Name
		tmp.Description = rec.Description

		//Find ModelPackage
		resClusterInfo, err := s.getClusterInfoByID(rec.ClusterInfoID, i)
		if err != nil {
			return nil, err
		}
		tmp.ClusterInfo.ClusterInfoID = resClusterInfo.ClusterInfo.ClusterInfoID
		tmp.ClusterInfo.Name = resClusterInfo.ClusterInfo.Name
		tmp.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint = resClusterInfo.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint

		listPredictionEnv = append(listPredictionEnv, *tmp)
	}

	resDTO.Rows = listPredictionEnv

	return resDTO, nil
}

func (s *PredictionEnvService) getClusterInfoByID(ClusterInfoID string, i identity.Identity) (*appDTO.InternalGetClusterInfoResponseDTO, error) {
	req := &appDTO.InternalGetClusterInfoRequestDTO{
		ClusterInfoID: ClusterInfoID,
	}

	res, err := s.clusterInfoSvc.GetByIDInternal(req, i)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *PredictionEnvService) GetByIDInternal(req *appDTO.InternalGetPredictionEnvRequestDTO, i identity.Identity) (*appDTO.InternalGetPredictionEnvResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.PredictionEnvID)
	if err != nil {
		return nil, err
	}
	// response dto
	resDTO := new(appDTO.InternalGetPredictionEnvResponseDTO)

	resDTO.Name = res.Name
	resDTO.Description = res.Description

	return resDTO, nil
}
