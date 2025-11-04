package main

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaReciever interface {
	Receive(msg *kafka.Message)
}

type KafkaConsumer struct {
	consumer *kafka.Consumer
	handlers map[string]KafkaReciever
}

func NewKafkaConsumer(conf *kafka.ConfigMap) *KafkaConsumer {
	c, err := kafka.NewConsumer(conf)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaConsumer{
		consumer: c,
		handlers: map[string]KafkaReciever{},
	}
}

func (c *KafkaConsumer) SubscribeTopics(topicName string, reciever KafkaReciever) {
	c.handlers[topicName] = reciever
}

func (c *KafkaConsumer) Listen(ctx context.Context, topics []string) error {
	c.consumer.Unsubscribe()
	err := c.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	for {
		event := c.consumer.Poll(100)
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *kafka.Message:
			handler, ok := c.handlers[*e.TopicPartition.Topic]
			if !ok {
				fmt.Printf("error unknown handlers: %s\n", *e.TopicPartition.Topic)
				continue
			}
			handler.Receive(e)

			_, err = c.consumer.CommitMessage(e)
			if err != nil {
				fmt.Printf("error commit message: %v\n", err)
			}
		case kafka.Error:
			fmt.Printf("error: %v\n", e)
		default:
			fmt.Printf("unknown event: %v\n", e)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
}
