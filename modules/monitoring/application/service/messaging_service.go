package service

import (
	"encoding/json"
	"fmt"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	infMsgSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/message_broker/kafka"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

type MessagingService struct {
	BaseService
}

func NewMessagingService(h *handler.Handler) {

	svc := new(MessagingService)

	svc.handler = h
	// base service init
	if err := svc.initBaseService(); err != nil {
		fmt.Errorf(err.Error())
	}
	consumer := infMsgSvc.NewConsumerKafka()

	consumer.SetTopic("datadrift-monitoring-data")
	err := consumer.RegisterConsumer()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	ch := make(chan infMsgSvc.OrgMsg, 1000)

	go func() {
		err := consumer.ConsumeMessage(ch)
		// 실행
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}()

	go func() {
		svc.MessageListener(ch)
	}()
}

func (m *MessagingService) MessageListener(ch chan infMsgSvc.OrgMsg) {
	MonitorSvc, err := NewMonitorService(m.handler)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	for n := range ch {
		mapMsg := infMsgSvc.KafkaMsg{}
		msgType := n.MsgType
		err := json.Unmarshal(n.Msg, &mapMsg)

		deploymentID := mapMsg.InferenceName
		status := mapMsg.DriftResult
		getByIdDTO := new(appDTO.MonitorGetByIDRequestDTO)
		getByIdDTO.ID = deploymentID

		domAggregateMonitor, err := MonitorSvc.GetByID(getByIdDTO)
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
