package service

import (
	"encoding/json"
	"fmt"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	infMsgSvc "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/infrastructure/message_broker/kafka"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

type IMonitorService interface {
	monitorStatusCheck(req *appDTO.MonitorStatusCheckRequestDTO) error
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

	//go func() {
	//	err := func() error {
	//		err := driftConsumer.ConsumeMessage(ch, "datadrift")
	//		// 실행
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	}()
	//	if err != nil {
	//	}
	//}()
	go consumeLoop(driftConsumer, ch, "datadrift")

	// accuracy go routine
	accuracyConsumer := infMsgSvc.NewConsumerKafka()
	accuracyConsumer.SetTopic("accuracy-monitoring-data")
	err = accuracyConsumer.RegisterConsumer(RegisterReq)
	if err != nil {
		return err
	}
	//go func() {
	//	err := func() error {
	//		err := accuracyConsumer.ConsumeMessage(ch, "accuracy")
	//		// 실행
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	}()
	//	if err != nil {
	//	}
	//}()
	go consumeLoop(accuracyConsumer, ch, "accuracy")

	// service health go routine
	serviceHealthConsumer := infMsgSvc.NewConsumerKafka()
	serviceHealthConsumer.SetTopic("servicehealth-monitoring-data")
	err = serviceHealthConsumer.RegisterConsumer(RegisterReq)
	if err != nil {
		return err
	}
	//go func() {
	//	err := func() error {
	//		err := serviceHealthConsumer.ConsumeMessage(ch, "servicehealth")
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	}()
	//	if err != nil {
	//	}
	//}()
	go consumeLoop(serviceHealthConsumer, ch, "servicehealth")

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
		if err != nil {
			fmt.Printf(err.Error())
		}

		deploymentID := mapMsg.InferenceName
		status := mapMsg.Result

		req := appDTO.MonitorStatusCheckRequestDTO{
			DeploymentID: deploymentID,
			Status:       status,
			Kind:         msgType,
		}

		err = m.MonitorService.monitorStatusCheck(&req)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}

func consumeLoop(consumer *infMsgSvc.ConsumerKafka, ch chan infMsgSvc.OrgMsg, target string) {
	defer func() {
		recover()
		go consumeLoop(consumer, ch, target)
	}()
	err := consumer.ConsumeMessage(ch, target)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
