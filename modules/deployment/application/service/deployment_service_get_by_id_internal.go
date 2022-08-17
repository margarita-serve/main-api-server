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
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/application/dto"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"

	//"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"

	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/deployment/infrastructure/repository"
)

// DeploymentService type
type DeploymentGetByIDInternalService struct {
	BaseService
	repo domRepo.IDeploymentRepo
}

// NewDeploymentService new DeploymentService
func NewDeploymentGetByIDInternalService(h *handler.Handler) (*DeploymentGetByIDInternalService, error) {
	var err error

	svc := new(DeploymentGetByIDInternalService)

	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repo, err = infRepo.NewDeploymentRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

func (s *DeploymentGetByIDInternalService) GetByIDInternal(req *appDTO.GetDeploymentRequestDTO, i identity.Identity) (*appDTO.GetDeploymentResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByIDInternal(req.DeploymentID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.GetDeploymentResponseDTO)
	resDTO.ID = res.ID
	resDTO.ProjectID = res.ProjectID
	resDTO.ModelPackageID = res.ModelPackageID
	resDTO.PredictionEnvID = res.PredictionEnvID
	resDTO.Name = res.Name
	resDTO.Description = res.Description
	resDTO.Importance = res.Importance
	resDTO.RequestCPU = res.RequestCPU
	resDTO.RequestMEM = res.RequestMEM
	resDTO.ActiveStatus = res.ActiveStatus
	resDTO.ServiceStatus = res.ServiceStatus
	resDTO.ChangeRequested = res.ChangeRequested

	return resDTO, nil
}
