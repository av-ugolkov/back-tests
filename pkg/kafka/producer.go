package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(conf *kafka.ConfigMap) *KafkaProducer {
	p, err := kafka.NewProducer(conf)
	if err != nil {
		return nil
	}

	return &KafkaProducer{
		producer: p,
	}
}

func (p *KafkaProducer) Send(topic string, key []byte, value any) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
		},
		Key:   key,
		Value: b,
	}, deliveryChan)
	if err != nil {
		return err
	}

	event := <-deliveryChan
	switch m := event.(type) {
	case *kafka.Message:
		fmt.Println(m)
	default:
		fmt.Printf("unknown format: %v", m)
	}

	return nil
}
