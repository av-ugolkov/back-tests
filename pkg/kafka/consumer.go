package main

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(conf *kafka.ConfigMap) *KafkaConsumer {
	c, err := kafka.NewConsumer(conf)
	if err != nil {
		return nil
	}
	return &KafkaConsumer{
		consumer: c,
	}
}

func (c *KafkaConsumer) Listen(ctx context.Context, topics []string) error {
	err := c.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	for {
		event := c.consumer.Poll(1000)
		switch e := event.(type) {
		case *kafka.Message:
			fmt.Println(e)
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
