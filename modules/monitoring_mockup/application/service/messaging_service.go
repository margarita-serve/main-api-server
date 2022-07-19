package service

import (
	"encoding/json"
	"fmt"

	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/application/dto"
	infMsgSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring_mockup/infrastructure/message_broker/kafka"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

type IMonitorService interface {
	GetByID(req *appDTO.MonitorGetByIDRequestDTO) (*appDTO.MonitorGetByIDResponseDTO, error)
}

type MessagingService struct {
	BaseService
	MonitorService IMonitorService
}

func NewMessagingService(h *handler.Handler, monitorSvc IMonitorService) error {

	svc := new(MessagingService)

	svc.handler = h
	svc.MonitorService = monitorSvc
	// base service init
	if err := svc.initBaseService(); err != nil {
		return err
	}
	consumer := infMsgSvc.NewConsumerKafka()

	consumer.SetTopic("datadrift-monitoring-data")
	err := consumer.RegisterConsumer()
	if err != nil {
		return err
	}
	ch := make(chan infMsgSvc.OrgMsg, 1000)

	go func() error {
		err := consumer.ConsumeMessage(ch)
		// 실행
		if err != nil {
			return err
		}
		return nil
	}()

	go func() {
		svc.messageListener(ch)
	}()

	return nil
}

func (m *MessagingService) messageListener(ch chan infMsgSvc.OrgMsg) {
	//MonitorSvc, err := NewMonitorService(m.handler)
	// if err != nil {
	// 	fmt.Errorf(err.Error())
	// }
	for n := range ch {
		mapMsg := infMsgSvc.KafkaMsg{}
		msgType := n.MsgType
		err := json.Unmarshal(n.Msg, &mapMsg)

		deploymentID := mapMsg.InferenceName
		status := mapMsg.DriftResult
		getByIdDTO := new(appDTO.MonitorGetByIDRequestDTO)
		getByIdDTO.ID = deploymentID

		domAggregateMonitor, err := m.MonitorService.GetByID(getByIdDTO)
		if msgType == "datadrift" {
			if err != nil {
				fmt.Printf(err.Error())
			} else if status == "pass" {
				domAggregateMonitor.Monitor.SetDriftStatusPass()
			} else if status == "warning" {
				domAggregateMonitor.Monitor.SetDriftStatusWarning()
			} else if status == "unknown" {
				domAggregateMonitor.Monitor.SetDriftStatusUnknown()
			} else {
				domAggregateMonitor.Monitor.SetDriftStatusUnknown()
			}
		} else if msgType == "accuracy" {

		}
	}
}
