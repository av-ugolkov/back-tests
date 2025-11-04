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
	close    bool
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

	for !c.close {
		msg, err := c.consumer.ReadMessage(100)
		if err != nil {
			continue
		}
		handler, ok := c.handlers[*msg.TopicPartition.Topic]
		if !ok {
			fmt.Printf("error unknown handlers: %s\n", *msg.TopicPartition.Topic)
			continue
		}
		handler.Receive(msg)

		_, err = c.consumer.CommitMessage(msg)
		if err != nil {
			fmt.Printf("error commit message: %v\n", err)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}

	return nil
}

func (c *KafkaConsumer) Close() {
	c.close = true
}
