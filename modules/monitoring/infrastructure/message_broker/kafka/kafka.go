package kafka

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
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

func (c *ConsumerKafka) RegisterConsumer() error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.88.151:32038",
		"group.id":          "datadrift-monitor-group",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		return err
	}
	c.Consumer = consumer

	return nil
}

func (c *ConsumerKafka) ConsumeMessage(ch chan OrgMsg) error {
	err := c.Consumer.SubscribeTopics([]string{c.Topic}, nil)
	if err != nil {
		return err
	}
	for {
		msg, err := c.Consumer.ReadMessage(-1)
		if err == nil {
			orgMsg := OrgMsg{
				Msg:     msg.Value,
				MsgType: "datadrift",
			}
			ch <- orgMsg
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Consumer.Close()
	return err
}
