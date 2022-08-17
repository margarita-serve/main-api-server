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
type DeploymentGovernanceHistoryService struct {
	BaseService
	repo domRepo.IDeploymentRepo
}

// NewDeploymentService new DeploymentService
func NewDeploymentGovernanceHistoryService(h *handler.Handler) (*DeploymentGovernanceHistoryService, error) {
	var err error

	svc := new(DeploymentGovernanceHistoryService)

	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repo, err = infRepo.NewDeploymentRepo(h); err != nil {
		return nil, err
	}

	return svc, nil
}

func (s *DeploymentGovernanceHistoryService) AddGovernanceHistory(req *appDTO.AddGovernanceHistoryRequestDTO, i identity.Identity) error {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByIDInternal(req.DeploymentID)
	if err != nil {
		return err
	}

	err = res.AddEventHistory(req.EventType, req.LogMessage, i.Claims.Username)
	if err != nil {
		return err
	}

	err = s.repo.Save(res)
	if err != nil {
		return err
	}

	return nil
}
