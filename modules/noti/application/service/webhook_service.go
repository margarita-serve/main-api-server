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
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/application/dto"
	domEntity "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/entity"
	domRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/domain/repository"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/identity"
	"github.com/rs/xid"

	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/infrastructure/repository"
)

// WebHookService type
type WebHookService struct {
	BaseService
	repo            domRepo.IWebHookRepo
	webHookEventSvc IWebHookEventService
}

type IWebHookEventService interface {
	ProcEvent(req *appDTO.WebHook) error
	SendEvent(req *appDTO.WebHook) error
}

// NewWebHookService new WebHookService
func NewWebHookService(h *handler.Handler, webHookEventService IWebHookEventService) (*WebHookService, error) {
	var err error

	svc := new(WebHookService)
	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repo, err = infRepo.NewWebHookRepo(h); err != nil {
		return nil, err
	}
	svc.webHookEventSvc = webHookEventService

	return svc, nil
}

// Create
func (s *WebHookService) Create(req *appDTO.CreateWebHookRequestDTO, i identity.Identity) (*appDTO.CreateWebHookResponseDTO, error) {
	//authorization
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
	domAggregateWebHook, err := domEntity.NewWebHook(
		guid,
		req.Name,
		req.DeploymentID,
		req.TriggerSource,
		req.URL,
		req.Method,
		req.CustomHeader,
		req.MessageBody,
		i.Claims.Username,
		req.TriggerStatus,
	)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(domAggregateWebHook)
	if err != nil {
		return nil, err
	}

	//response dto
	resDTO := new(appDTO.CreateWebHookResponseDTO)
	resDTO.WebHookID = domAggregateWebHook.ID

	return resDTO, nil
}

func (s *WebHookService) Delete(req *appDTO.DeleteWebHookRequestDTO, i identity.Identity) error {
	//authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	domAggregateWebHook, err := s.repo.GetByID(req.WebHookID)
	if err != nil {
		return err
	}

	err = s.repo.Delete(domAggregateWebHook.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *WebHookService) UpdateWebHook(req *appDTO.UpdateWebHookRequestDTO, i identity.Identity) error {
	//authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	//toBe...
	//userID := "testID"

	//Find Domain Entity
	domAggregateWebHook, err := s.repo.GetForUpdate(req.WebHookID)
	if err != nil {
		return err
	}

	if req.Name != nil {
		domAggregateWebHook.SetName(*req.Name)
	}
	if req.TriggerSource != nil {
		domAggregateWebHook.SetTriggerSource(*req.TriggerSource)
	}
	if req.CustomHeader != nil {
		domAggregateWebHook.SetCustomHeader(*req.CustomHeader)
	}
	if req.MessageBody != nil {
		domAggregateWebHook.SetMessageBody(*req.MessageBody)
	}
	if req.Method != nil {
		domAggregateWebHook.SetMethod(*req.Method)
	}
	if req.URL != nil {
		domAggregateWebHook.SetURL(*req.URL)
	}

	err = domEntity.Validate(domAggregateWebHook)
	if err != nil {
		return err
	}

	err = s.repo.Save(domAggregateWebHook)
	if err != nil {
		return err
	}

	return nil
}

func (s *WebHookService) GetByID(req *appDTO.GetWebHookRequestDTO, i identity.Identity) (*appDTO.GetWebHookResponseDTO, error) {
	//authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	res, err := s.repo.GetByID(req.WebHookID)
	if err != nil {
		return nil, err
	}

	// response dto
	resDTO := new(appDTO.GetWebHookResponseDTO)
	resDTO.ID = res.ID
	resDTO.DeploymentID = res.DeploymentID
	resDTO.Name = res.Name
	resDTO.CustomHeader = res.CustomHeader
	resDTO.MessageBody = res.MessageBody
	resDTO.Method = res.Method
	resDTO.URL = res.URL
	resDTO.TriggerStatus = res.TriggerStatus
	resDTO.TriggerSource = res.TriggerSource

	return resDTO, nil
}

func (s *WebHookService) GetList(req *appDTO.GetWebHookListRequestDTO, i identity.Identity) (*appDTO.GetWebHookListResponseDTO, error) {
	//authorization
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

	filterEntity := domEntity.WebHook{
		DeploymentID: req.DeploymentID,
	}

	resultList, pagination, err := s.repo.GetList(req.Name, reqP, filterEntity)
	if err != nil {
		return nil, err
	}

	//interface type을 concrete type으로 변환
	//domain layer에서 pagination type을 모르기 때문에 interface type으로 정의 후 변환한다
	p := pagination.(infRepo.Pagination)

	// response dto
	resDTO := new(appDTO.GetWebHookListResponseDTO)
	resDTO.Limit = p.Limit
	resDTO.Page = p.Page
	resDTO.TotalRows = p.TotalRows
	resDTO.TotalPages = p.TotalPages

	var listWebHook []*appDTO.GetWebHookResponseDTO
	for _, rec := range resultList {
		tmp := new(appDTO.GetWebHookResponseDTO)

		tmp.ID = rec.ID
		tmp.DeploymentID = rec.DeploymentID
		tmp.Name = rec.Name
		tmp.TriggerSource = rec.TriggerSource
		tmp.CustomHeader = rec.CustomHeader
		tmp.MessageBody = rec.MessageBody
		tmp.Method = rec.Method
		tmp.URL = rec.URL
		tmp.TriggerStatus = rec.TriggerStatus

		listWebHook = append(listWebHook, tmp)
	}

	resDTO.Rows = listWebHook

	return resDTO, nil
}

func (s *WebHookService) GetListByInternal(req *appDTO.InternalGetWebHookRequestDTO, i identity.Identity) (*appDTO.InternalGetWebHookResponseDTO, error) {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }
	filterEntity := domEntity.WebHook{DeploymentID: req.DeploymentID,
		TriggerSource: req.TriggerSource,
		TriggerStatus: req.TriggerStatus,
	}

	resultList, err := s.repo.GetListByInternal(filterEntity)
	if err != nil {
		return nil, err
	}
	// response dto
	resDTO := new(appDTO.InternalGetWebHookResponseDTO)
	var listWebHook []*appDTO.WebHook
	for _, rec := range resultList {
		tmp := new(appDTO.WebHook)

		tmp.ID = rec.ID
		tmp.DeploymentID = rec.DeploymentID
		tmp.Name = rec.Name
		tmp.TriggerSource = rec.TriggerSource
		tmp.TriggerStatus = rec.TriggerStatus
		tmp.CustomHeader = rec.CustomHeader
		tmp.MessageBody = rec.MessageBody
		tmp.Method = rec.Method
		tmp.URL = rec.URL

		listWebHook = append(listWebHook, tmp)
	}

	resDTO.WebHookList = listWebHook

	return resDTO, nil
}

func (s *WebHookService) TestWebHookSend(req *appDTO.TestWebHookRequestDTO, i identity.Identity) error {
	//authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	reqWebHookTest := &appDTO.WebHook{
		TriggerSource: req.TriggerSource,
		URL:           req.URL,
		Method:        req.Method,
		CustomHeader:  req.CustomHeader,
		MessageBody:   req.MessageBody,
	}
	err := s.webHookEventSvc.SendEvent(reqWebHookTest)
	if err != nil {
		return err
	}

	// response dto

	return nil
}

func (s *WebHookService) SendWebHook(req *appDTO.SendWebHookRequestDTO, i identity.Identity) error {
	// //authorization
	// if i.CanAccessCurrentRequest() == false {
	// 	errMsg := fmt.Sprintf("You are not authorized to access [`%s.%s`]",
	// 		i.RequestInfo.RequestObject, i.RequestInfo.RequestAction)
	// 	return nil, sysError.CustomForbiddenAccess(errMsg)
	// }

	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }

	reqWebHookList := &appDTO.InternalGetWebHookRequestDTO{
		DeploymentID:  req.DeploymentID,
		TriggerSource: req.TriggerSource,
		TriggerStatus: req.TriggerStatus,
	}
	resWebHookList, err := s.GetListByInternal(reqWebHookList, i)
	if err != nil {
		return err
	}

	for _, rec := range resWebHookList.WebHookList {
		switch rec.TriggerStatus {
		case "AtRisk":
			if req.TriggerStatus == "AtRisk" || req.TriggerStatus == "Failing" {
				go s.webHookEventSvc.ProcEvent(rec)
			}
		case "Failing":
			if req.TriggerStatus == "Failing" {
				go s.webHookEventSvc.ProcEvent(rec)
			}
		default:
			go s.webHookEventSvc.ProcEvent(rec)
		}

	}

	return nil
}
