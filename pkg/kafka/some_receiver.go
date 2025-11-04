package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var _ KafkaReciever = (*Receiver)(nil)

type Receiver struct {
}

func NewReceiver() *Receiver {
	return &Receiver{}
}

func (r *Receiver) Receive(msg *kafka.Message) {
	fmt.Printf("Received: key=%s value=%s partition=%d offset=%d time=%v\n",
		string(msg.Key),
		string(msg.Value),
		msg.TopicPartition.Partition,
		msg.TopicPartition.Offset,
		time.Now().Format(time.TimeOnly))
}
