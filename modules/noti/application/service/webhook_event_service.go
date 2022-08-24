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
	"github.com/rs/xid"

	infRepo "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/infrastructure/repository"
	webHookEventSendSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/noti/infrastructure/webhook_event_sender"
)

// webHookEventService type
type WebHookEventService struct {
	BaseService
	repo                domRepo.IWebHookEventRepo
	webHookEventSendSvc *webHookEventSendSvc.WebHookEventSender
}

// NewwebHookEventService new webHookEventService
func NewWebHookEventService(h *handler.Handler) (*WebHookEventService, error) {
	var err error

	svc := new(WebHookEventService)
	svc.handler = h
	if err := svc.initBaseService(); err != nil {
		return nil, err
	}

	if svc.repo, err = infRepo.NewWebHookEventRepo(h); err != nil {
		return nil, err
	}

	if svc.webHookEventSendSvc, err = webHookEventSendSvc.NewWebHookEventSendService(h); err != nil {
		return nil, err
	}

	return svc, nil
}

// Create
func (s *WebHookEventService) ProcEvent(rec *appDTO.WebHook) error {
	guid := xid.New().String()

	// New deployment domain Instance
	domAggregateWebHookEvent, err := domEntity.NewWebHookEvent(
		guid,
		rec.DeploymentID,
		rec.URL,
		rec.Method,
		rec.CustomHeader,
		rec.MessageBody,
		rec.TriggerSource,
	)
	if err != nil {
		return err
	}

	var sendStatus = ""
	err = s.SendEvent(rec)
	if err != nil {
		sendStatus = "Fail"
	} else {
		sendStatus = "Success"
	}

	domAggregateWebHookEvent.SetSendStatus(sendStatus)

	err = s.repo.Save(domAggregateWebHookEvent)
	if err != nil {
		return err
	}

	return nil
}

func (s *WebHookEventService) SendEvent(rec *appDTO.WebHook) error {
	_, err := s.webHookEventSendSvc.SendWebHookEvent(rec.URL, rec.Method, rec.CustomHeader, rec.MessageBody)
	if err != nil {
		return err
	}

	return nil
}
