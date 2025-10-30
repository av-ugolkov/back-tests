package main

import (
	"fmt"
	"time"

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

func (p *KafkaProducer) Send(topic string, key []byte, value []byte) (*kafka.Message, error) {
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: value,
	}, deliveryChan)
	if err != nil {
		return nil, err
	}

	event := <-deliveryChan
	switch m := event.(type) {
	case *kafka.Message:
		fmt.Printf("Sender: key=%s value=%s topic=%s partition=%d offset=%d time=%v\n",
			string(m.Key),
			string(m.Value),
			*m.TopicPartition.Topic,
			m.TopicPartition.Partition,
			m.TopicPartition.Offset,
			time.Now())
		return m, nil
	default:
		return nil, fmt.Errorf("unknown format: %v", m)
	}
}
