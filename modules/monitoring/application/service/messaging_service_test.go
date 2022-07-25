package service_test

import (
	"fmt"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/service"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"testing"
)

func newMessagingService(t *testing.T) {
	h, err := handler.NewHandler()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	service.NewMessagingService(h, "accuracy-monitoring-data", "accuracy")
}

func TestMessageBrokerSvc_Create(t *testing.T) {
	newMessagingService(t)
}
