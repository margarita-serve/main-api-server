package service

import (
	"encoding/json"
	"fmt"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	infMsgSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/message_broker/kafka"
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

	ch := make(chan infMsgSvc.OrgMsg, 1000)

	cfg, err := h.GetConfig()
	if err != nil {
		return err
	}
	RegisterReq := new(appDTO.RegisterServer)
	RegisterReq.Endpoint = cfg.Connectors.Kafka.Endpoint
	RegisterReq.GroupID = cfg.Connectors.Kafka.GroupID
	RegisterReq.AutoOffsetReset = cfg.Connectors.Kafka.AutoOffsetReset

	// datadrift go routine
	driftConsumer := infMsgSvc.NewConsumerKafka()
	driftConsumer.SetTopic("datadrift-monitoring-data")
	err = driftConsumer.RegisterConsumer(RegisterReq)
	if err != nil {
		return err
	}
	go func() {
		err := func() error {
			err := driftConsumer.ConsumeMessage(ch, "datadrift")
			// 실행
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
		}
	}()

	// accuracy go routine
	accuracyConsumer := infMsgSvc.NewConsumerKafka()
	accuracyConsumer.SetTopic("accuracy-monitoring-data")
	err = accuracyConsumer.RegisterConsumer(RegisterReq)
	if err != nil {
		return err
	}
	go func() {
		err := func() error {
			err := accuracyConsumer.ConsumeMessage(ch, "accuracy")
			// 실행
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
		}
	}()

	// Message Listener go routine
	go func() {
		svc.MessageListener(ch)
	}()
	return nil
}

func (m *MessagingService) MessageListener(ch chan infMsgSvc.OrgMsg) {

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
			} else {
				result := domAggregateMonitor.Monitor.CheckDriftStatus(status)
				if result == true {
					// 변경 완료된 경우
				}
			}
			//else if status == "pass" {
			//	domAggregateMonitor.Monitor.SetDriftStatusPass()
			//} else if status == "atrisk" {
			//	domAggregateMonitor.Monitor.SetDriftStatusAtRisk()
			//} else if status == "failing" {
			//	domAggregateMonitor.Monitor.SetDriftStatusFailing()
			//} else {
			//	domAggregateMonitor.Monitor.SetDriftStatusUnknown()
			//}
		} else if msgType == "accuracy" {
			if err != nil {
				fmt.Printf(err.Error())
			} else if status == "pass" {
				domAggregateMonitor.Monitor.SetAccuracyStatusPass()
			} else if status == "atrisk" {
				domAggregateMonitor.Monitor.SetAccuracyStatusAtRisk()
			} else if status == "failing" {
				domAggregateMonitor.Monitor.SetAccuracyStatusFailing()
			} else {
				domAggregateMonitor.Monitor.SetAccuracyStatusUnknown()
			}
		}
	}
}
