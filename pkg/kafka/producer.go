package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(conf *kafka.ConfigMap) *KafkaProducer {
	p, err := kafka.NewProducer(conf)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered key=%s value=%s partition=%d offset=%d time=%v\n",
						string(ev.Key),
						string(ev.Value),
						ev.TopicPartition.Partition,
						ev.TopicPartition.Offset,
						time.Now().Format(time.TimeOnly))
				}
			}
		}
	}()

	return &KafkaProducer{
		producer: p,
	}
}

func (p *KafkaProducer) Send(topic string, key []byte, value []byte) error {
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: value,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
func (p *KafkaProducer) Close() {
	p.producer.Flush(10000)
	p.producer.Close()
}
