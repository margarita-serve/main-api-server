package kafka

import (
	"fmt"
	appDTO "git.k3.acornsoft.io/msit-auto-ml/koreserv/modules/monitoring/application/dto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"time"
)

type ConsumerKafka struct {
	Topic    string
	Consumer *kafka.Consumer
}

func NewConsumerKafka() *ConsumerKafka {
	consumerKafka := &ConsumerKafka{
		Topic:    "",
		Consumer: nil,
	}
	return consumerKafka
}

func (c *ConsumerKafka) SetTopic(topic string) {
	c.Topic = topic
}

func (c *ConsumerKafka) RegisterConsumer(server *appDTO.RegisterServer) error {
	// config 수정
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server.Endpoint,
		"group.id":          server.GroupID,
		//"group.id":          "test",
		"auto.offset.reset": server.AutoOffsetReset,
	})
	if err != nil {
		return err
	}
	c.Consumer = consumer

	return nil
}

func (c *ConsumerKafka) ConsumeMessage(ch chan OrgMsg, msgType string) error {
	err := c.Consumer.SubscribeTopics([]string{c.Topic}, nil)
	if err != nil {
		return err
	}
	for {
		msg, err := c.Consumer.ReadMessage(-1)
		if err == nil {
			orgMsg := OrgMsg{
				Msg:     msg.Value,
				MsgType: msgType,
			}
			ch <- orgMsg
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v", err)
			time.Sleep(30000 * time.Millisecond)
		}
	}

	c.Consumer.Close()
	return err
}
