package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var _ KafkaReciever = (*Receiver)(nil)

type Receiver struct {
	wg *sync.WaitGroup
}

func NewReceiver(wg *sync.WaitGroup) *Receiver {
	return &Receiver{
		wg: wg,
	}
}

func (r *Receiver) Receive(msg *kafka.Message) {
	fmt.Printf("Received: key=%s value=%s topic=%s partition=%d offset=%d time=%v\n",
		string(msg.Key),
		string(msg.Value),
		*msg.TopicPartition.Topic,
		msg.TopicPartition.Partition,
		msg.TopicPartition.Offset,
		time.Now())
	r.wg.Done()
}
