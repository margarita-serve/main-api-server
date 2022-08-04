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

type ClusterInfoService struct {
	BaseService
	repo domRepo.IClusterInfoRepo
}

// NewClusterInfoService new ClusterInfoService
func NewClusterInfoService(h *handler.Handler) (*ClusterInfoService, error) {
	var err error

	svc := new(ClusterInfoService)
	svc.handler = h
	// if err := svc.initBaseService(); err != nil {
	// 	return nil, err
	// }

	if svc.repo, err = infRepo.NewClusterInfoRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// Create
func (s *ClusterInfoService) Create(req *appDTO.CreateClusterInfoRequestDTO, i identity.Identity) (*appDTO.CreateClusterInfoResponseDTO, error) {
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
	domAggregateClusterInfo, err := domEntity.NewClusterInfo(
		guid,
		req.Name,
		req.InfereceSvcAPISvrEndPoint,
		req.InfereceSvcHostName,
		req.InferenceSvcIngressEndPoint,
		i.Claims.Username,
	)
	if err != nil {
		return nil, err
	}

	// domEntity.Validate(domAggregateClusterInfo)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregateClusterInfo)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.CreateClusterInfoResponseDTO)
	resDTO.ClusterInfoID = domAggregateClusterInfo.ID

	return resDTO, nil
}

func (s *ClusterInfoService) Delete(req *appDTO.DeleteClusterInfoRequestDTO, i identity.Identity) (*appDTO.DeleteClusterInfoResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	domAggregateClusterInfo, err := s.repo.GetByID(req.ClusterInfoID)
	if err != nil {
		return nil, err
	}

	resDTO := new(appDTO.DeleteClusterInfoResponseDTO)

	err = s.repo.Delete(domAggregateClusterInfo.ID)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "ClusterInfo Delete Success"

	return resDTO, nil
}

func (s *ClusterInfoService) UpdateClusterInfo(req *appDTO.UpdateClusterInfoRequestDTO, i identity.Identity) (*appDTO.UpdateClusterInfoResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	//userID := "testID"

	//Find Domain Entity
	domAggregateClusterInfo, err := s.repo.GetForUpdate(req.ClusterInfoID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.UpdateClusterInfoResponseDTO)

	if req.Name != "" {
		domAggregateClusterInfo.SetName(req.Name)
	}
	if req.InfereceSvcAPISvrEndPoint != "" {
		domAggregateClusterInfo.SetInfereceSvcAPISvrEndPoint(req.InfereceSvcAPISvrEndPoint)
	}
	if req.InfereceSvcHostName != "" {
		domAggregateClusterInfo.SetInfereceSvcHostName(req.InfereceSvcHostName)
	}
	if req.InferenceSvcIngressEndPoint != "" {
		domAggregateClusterInfo.SetInferenceSvcIngressEndPoint(req.InferenceSvcIngressEndPoint)
	}

	// domEntity.Validate(domAggregateClusterInfo)
	// if err != nil {
	// 	return nil, err
	// }

	err = s.repo.Save(domAggregateClusterInfo)
	if err != nil {
		return nil, err
	}

	resDTO.Message = "ClusterInfo Update Success"

	return resDTO, nil
}

func (s *ClusterInfoService) GetByID(req *appDTO.GetClusterInfoRequestDTO, i identity.Identity) (*appDTO.GetClusterInfoResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.ClusterInfoID)
	if err != nil {
		return nil, err
	}
	// response dto
	resDTO := new(appDTO.GetClusterInfoResponseDTO)

	resDTO.ClusterInfo.Name = res.Name
	resDTO.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint = res.InferenceSvcInfo.InfereceSvcAPISvrEndPoint
	resDTO.ClusterInfo.InferenceSvcInfo.InfereceSvcHostName = res.InferenceSvcInfo.InfereceSvcHostName
	resDTO.ClusterInfo.InferenceSvcInfo.InferenceSvcIngressEndPoint = res.InferenceSvcInfo.InferenceSvcIngressEndPoint

	return resDTO, nil
}

func (s *ClusterInfoService) GetList(req *appDTO.GetClusterInfoListRequestDTO, i identity.Identity) (*appDTO.GetClusterInfoListResponseDTO, error) {
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
	resDTO := new(appDTO.GetClusterInfoListResponseDTO)
	resDTO.Limit = p.Limit
	resDTO.Page = p.Page
	resDTO.TotalRows = p.TotalRows
	resDTO.TotalPages = p.TotalPages

	var listClusterInfo []appDTO.GetClusterInfoResponseDTO
	for _, rec := range resultList {
		tmp := new(appDTO.GetClusterInfoResponseDTO)

		tmp.ClusterInfo.ClusterInfoID = rec.ID
		tmp.ClusterInfo.Name = rec.Name
		tmp.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint = rec.InferenceSvcInfo.InfereceSvcAPISvrEndPoint
		tmp.ClusterInfo.InferenceSvcInfo.InfereceSvcHostName = rec.InferenceSvcInfo.InfereceSvcHostName
		tmp.ClusterInfo.InferenceSvcInfo.InferenceSvcIngressEndPoint = rec.InferenceSvcInfo.InferenceSvcIngressEndPoint

		listClusterInfo = append(listClusterInfo, *tmp)
	}

	resDTO.Rows = listClusterInfo

	return resDTO, nil
}

func (s *ClusterInfoService) GetByIDInternal(req *appDTO.InternalGetClusterInfoRequestDTO, i identity.Identity) (*appDTO.InternalGetClusterInfoResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.ClusterInfoID)
	if err != nil {
		return nil, err
	}
	// response dto
	resDTO := new(appDTO.InternalGetClusterInfoResponseDTO)

	resDTO.ClusterInfo.Name = res.Name
	resDTO.ClusterInfo.InferenceSvcInfo.InfereceSvcAPISvrEndPoint = res.InferenceSvcInfo.InfereceSvcAPISvrEndPoint
	resDTO.ClusterInfo.InferenceSvcInfo.InfereceSvcHostName = res.InferenceSvcInfo.InfereceSvcHostName
	resDTO.ClusterInfo.InferenceSvcInfo.InferenceSvcIngressEndPoint = res.InferenceSvcInfo.InferenceSvcIngressEndPoint

	return resDTO, nil
}
