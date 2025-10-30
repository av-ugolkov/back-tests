package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const topicName = "kafka-test-request"

func main() {
	cfgProducer := kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092",
		"client.id":         "backend-examples",
		"acks":              "all",
	}
	producer := NewKafkaProducer(&cfgProducer)

	cfgConsumer := kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092",
		"group.id":          "backend-examples",
		"auto.offset.reset": "earliest",
	}
	consumer := NewKafkaConsumer(&cfgConsumer)

	var wg sync.WaitGroup
	r := NewReceiver(&wg)
	consumer.SubscribeTopics(topicName, r)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		err := consumer.Listen(ctx, []string{topicName})
		if err != nil {
			fmt.Printf("error listen: %v\n", err)
		}
	}()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		_, err := producer.Send(topicName, []byte(fmt.Sprintf("key-%d", i)), []byte(fmt.Sprintf("value-%d", i)))
		if err != nil {
			fmt.Printf("error producer send: %v\n", err)
		}
	}

	wg.Wait()
}
